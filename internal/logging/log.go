package logging

import (
	"bufio"
	"fmt"
	goLog "log"
	"os"
	"time"
)

type Log struct {
	buffer *bufio.Writer
	logger *goLog.Logger
}

// Creates new logger with a buffer to Stdout
func New() Log {
	buf := bufio.NewWriter(os.Stdout)

	return NewUsingBuffer(buf)
}

// Create a new logger using a buffer, such as to a file
func NewUsingBuffer(buffer *bufio.Writer) Log {
	return Log{
		buffer: buffer,
		logger: goLog.New(buffer, "", 0),
	}
}

// Used internally to print in the desired format
func (l Log) internalPrint(level string, message string) {
	// Format now as dd-mm-yyyy hh:mm:ss ±hhmm
	timestamp := time.Now().UTC().Format("02-01-2006 15:04:05 -0700")

	l.logger.Printf("[%s - %s] %s", timestamp, level, message)
	l.buffer.Flush()
}

func (l Log) EPrintf(format string, v ...interface{}) {
	l.internalPrint("ERROR", fmt.Sprintf(format, v...))
}

func (l Log) IPrintf(format string, v ...interface{}) {
	l.internalPrint("INFO", fmt.Sprintf(format, v...))
}
