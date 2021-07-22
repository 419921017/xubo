package main

import (
	"fmt"
	"xubo/chapter07/7-2.logger/console"
	"xubo/chapter07/7-2.logger/file"
	"xubo/chapter07/7-2.logger/logger"
)

func main() {
	l := createLogger()
	l.Log("hello")
}

func createLogger() *logger.Logger {
	// 创建日志器
	l := logger.NewLogger()
	// 创建命令行写入器
	cw := console.NewConsoleWriter()

	l.RegisterWriter(cw)

	fw := file.NewFileWriter()

	if err := fw.SetFile("log.log"); err != nil {
		fmt.Println(err)
	}

	l.RegisterWriter(fw)

	return l

}
