package golog

import (
	"context"
	"fmt"
	"os"
	"sync"
	"time"
)

// defaultBufferSize
const defaultBufferSize = 4096

const defaultFlushInterval = time.Minute * 5

// FileAppender
type FileAppender struct {
	file           *os.File
	bufferedWriter *bufferedWriter
	mu             *sync.Mutex
	activated      bool
	ticker         *time.Ticker
	stopTicker     context.CancelFunc
}

// NewFileAppender returns new FileAppender
func NewFileAppender(fileName string) (asyncFileAppender *FileAppender, err error) {
	return NewFileAppenderWithBufferSizeAndFlushInterval(fileName, defaultBufferSize, defaultFlushInterval)
}

// NewFileAppenderWithBufferSizeAndFlushInterval returns new FileAppender
func NewFileAppenderWithBufferSizeAndFlushInterval(fileName string, bufferSize int, flushInterval time.Duration) (asyncFileAppender *FileAppender, err error) {
	file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return nil, err
	}

	size := defaultBufferSize
	if bufferSize > 0 {
		size = bufferSize
	}

	ctx, cancel := context.WithCancel(context.Background())
	appender := &FileAppender{
		file:           file,
		bufferedWriter: newBufferedWriter(file, withBufferSize(size)),
		mu:             new(sync.Mutex),
		activated:      true,
		ticker:         time.NewTicker(flushInterval),
		stopTicker:     cancel,
	}

	go func() {
		for {
			select {
			case <-appender.ticker.C:
			case <-ctx.Done():
			}
			appender.mu.Lock()
			if !appender.activated {
				appender.ticker.Stop()
				appender.mu.Unlock()
				return
			}
			appender.bufferedWriter.Flush()
			appender.mu.Unlock()
		}
	}()

	return appender, nil
}

// Write implements io.Write
func (appender *FileAppender) Write(data []byte) (n int, err error) {
	appender.mu.Lock()
	defer appender.mu.Unlock()
	if appender.activated {
		data = append(data, '\n')
		return appender.bufferedWriter.Write(data)
	}
	return 0, fmt.Errorf("appender is closed")
}

// Close implements io.Closer
func (appender *FileAppender) Close() error {
	appender.mu.Lock()
	defer func() {
		appender.activated = false
		appender.stopTicker()
		appender.mu.Unlock()
	}()

	if appender.activated {

		if err := appender.bufferedWriter.Flush(); err != nil {
			return err
		}

		return appender.file.Close()
	}

	return nil
}
