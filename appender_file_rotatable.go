package golog

import (
	"os"
	"os/signal"
	"syscall"
	"sync"
)

// RotatableFileAppender RotatableFileAppender struct
type RotatableFileAppender struct {
	*FileAppender
}

// NewRotatableFileAppender returns new FileAppender
func NewRotatableFileAppender(fileName string) (asyncFileAppender *RotatableFileAppender, err error) {
	return NewRotatableFileAppenderWithBufferSize(fileName, defaultBufferSize)
}

// NewRotatableFileAppenderWithBufferSize returns new FileAppender
func NewRotatableFileAppenderWithBufferSize(fileName string, bufferSize int) (asyncFileAppender *RotatableFileAppender, err error) {

	fileAppender, err := NewFileAppenderWithBufferSize(fileName, bufferSize)
	if err != nil {
		return nil, err
	}

	appender := &RotatableFileAppender{
		FileAppender: fileAppender,
	}

	hup := make(chan os.Signal, 1)
	signal.Notify(hup, syscall.SIGHUP)
	mutex := new(sync.Mutex)

	go func() {
		for {
			<-hup
			mutex.Lock()

			appender.FileAppender.Close()

			newAppender, err := NewFileAppenderWithBufferSize(fileName, bufferSize)
			if err != nil {
				panic(err)
			}
			appender.FileAppender = newAppender

			mutex.Unlock()
		}
	}()

	return appender, nil
}

// Write implements io.Write
func (appender *RotatableFileAppender) Write(data []byte) (n int, err error) {
	return appender.FileAppender.Write(data)
}

// Close implements io.Closer
func (appender *RotatableFileAppender) Close() error {
	return appender.FileAppender.Close()
}
