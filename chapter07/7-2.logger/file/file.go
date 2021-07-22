package file

import (
	"errors"
	"fmt"
	"os"
)

type FileWriter struct {
	file *os.File
}

func NewFileWriter() *FileWriter {
	return &FileWriter{}
}

func (f *FileWriter) SetFile(filename string) (err error) {
	if f.file != nil {
		f.file.Close()
	}

	f.file, err = os.Create(filename)
	return err
}

func (f *FileWriter) Write(data interface{}) (err error) {
	if f.file == nil {
		return errors.New("file not created")
	}
	str := fmt.Sprintf("%v\n", data)
	_, err = f.file.Write([]byte(str))
	return err
}
