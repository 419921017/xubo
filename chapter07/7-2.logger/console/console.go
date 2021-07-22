package console

import (
	"fmt"
	"os"
)

type consoleWriter struct{}

func NewConsoleWriter() *consoleWriter {
	return &consoleWriter{}
}

func (c *consoleWriter) Write(data interface{}) error {
	str := fmt.Sprintf("%v", data)
	_, err := os.Stdout.Write([]byte(str))
	return err
}
