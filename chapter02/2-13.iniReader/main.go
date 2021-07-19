package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(getValue("example.ini", `remote "origin"`, "url"))
}

func getValue(filename, expectSection, expectKey string) string {
	file, err := os.Open(filename)

	if err != nil {
		return ""
	}

	defer file.Close()

	reader := bufio.NewReader(file)

	var sectionName string

	for {
		// 使用reader.ReadString()从文件中读取字符串，直到碰到“\n”，也就是行结束
		lineStr, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		// 每一行的文本可能会在标识符两边混杂有空格、回车符、换行符等不可见的空白字符，使用strings.Trim Space()可以去掉这些空白字符
		lineStr = strings.TrimSpace(lineStr)

		if lineStr == "" {
			continue
		}
		// 当行首的字符为“；”分号时，表示这一行是注释行，忽略一整行的读取
		if lineStr[0] == ';' {
			continue
		}

		if lineStr[0] == '[' && lineStr[len(lineStr)-1] == ']' {
			sectionName = lineStr[1 : len(lineStr)-1]
		} else if sectionName == expectSection {
			pair := strings.Split(lineStr, "=")
			if len(pair) == 2 {
				key := strings.TrimSpace(pair[0])
				if key == expectKey {
					return strings.TrimSpace(pair[1])
				}
			}
		}
	}
	return ""
}
