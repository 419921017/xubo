package main

import "fmt"

func main() {
	angel := "Heros never die"
	fmt.Println(modifyString(angel))
}

func modifyString(str string) string {
	strBytes := []byte(str)

	for i := 5; i < len(strBytes)-3; i++ {
		strBytes[i] = ' '
	}

	return string(strBytes)

}
