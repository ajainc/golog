package golog

import (
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

// RotatableFileAppender RotatableFileAppender struct
type RotatableFileAppender struct {
	*FileAppender
	mu *sync.Mutex
}

// NewRotatableFileAppender returns new FileAppender
func NewRotatableFileAppender(fileName string) (asyncFileAppender *RotatableFileAppender, err error) {
	return NewRotatableFileAppenderWithBufferSize(fileName, defaultBufferSize)
}

// NewRotatableFileAppenderWithBufferSize returns new FileAppender
func NewRotatableFileAppenderWithBufferSize(fileName string, bufferSize int) (asyncFileAppender *RotatableFileAppender, err error) {
	return NewRotatableFileAppenderWithBufferSizeAndFlushInterval(fileName, bufferSize, defaultFlushInterval)
}

// NewRotatableFileAppenderWithBufferSizeAndFlushInterval returns new FileAppender
func NewRotatableFileAppenderWithBufferSizeAndFlushInterval(fileName string, bufferSize int, flushInterval time.Duration) (asyncFileAppender *RotatableFileAppender, err error) {

	fileAppender, err := NewFileAppenderWithBufferSizeAndFlushInterval(fileName, bufferSize, flushInterval)
	if err != nil {
		return nil, err
	}

	appender := &RotatableFileAppender{
		FileAppender: fileAppender,
		mu:           new(sync.Mutex),
	}

	hup := make(chan os.Signal, 1)
	signal.Notify(hup, syscall.SIGHUP)

	go func() {
		for {
			<-hup
			appender.mu.Lock()

			appender.FileAppender.Close()

			newAppender, err := NewFileAppenderWithBufferSizeAndFlushInterval(fileName, bufferSize, flushInterval)
			if err != nil {
				panic(err)
			}
			appender.FileAppender = newAppender

			appender.mu.Unlock()
		}
	}()

	return appender, nil
}

// Write implements io.Write
func (appender *RotatableFileAppender) Write(data []byte) (n int, err error) {
	appender.mu.Lock()
	defer appender.mu.Unlock()
	return appender.FileAppender.Write(data)
}

// Close implements io.Closer
func (appender *RotatableFileAppender) Close() error {
	appender.mu.Lock()
	defer appender.mu.Unlock()
	return appender.FileAppender.Close()
}
