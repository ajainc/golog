package golog

import (
	"io/ioutil"
	"os"
	"syscall"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestRotatableFileAppender_Write(t *testing.T) {

	var setup = func() {
		file1, _ := os.OpenFile("appender_file_rotatable", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
		file1.Close()
		file2, _ := os.OpenFile("appender_file_rotatable_bk", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
		file2.Close()
	}

	var cleanup = func() {
		os.Remove("appender_file_rotatable")
		os.Remove("appender_file_rotatable_bk")
	}

	t.Run("", func(t *testing.T) {
		setup()

		appender, err := NewRotatableFileAppender("appender_file_rotatable")
		assert.Nil(t, err)
		appender.Write([]byte("test1"))

		os.Rename("appender_file_rotatable", "appender_file_rotatable_bk")
		syscall.Kill(syscall.Getpid(), syscall.SIGHUP)
		time.Sleep(1 * time.Second)

		appender.Write([]byte("test2"))
		appender.Close()

		file1, err := os.OpenFile("appender_file_rotatable", os.O_RDONLY, 0666)
		assert.Nil(t, err)
		actual, err := ioutil.ReadAll(file1)
		expected := "test2\n"
		assert.Equal(t, expected, string(actual))

		file2, err := os.OpenFile("appender_file_rotatable_bk", os.O_RDONLY, 0666)
		assert.Nil(t, err)
		actual, err = ioutil.ReadAll(file2)
		assert.Nil(t, err)
		expected = "test1\n"
		assert.Equal(t, expected, string(actual))

		cleanup()
	})
}
