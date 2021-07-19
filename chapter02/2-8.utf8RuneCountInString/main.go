package main

import (
	"fmt"
	"strings"
)

func main() {
	// tip1 := "genji is a ninja"
	// tip2 := "忍者"
	// fmt.Println(len(tip1))
	// fmt.Println(len(tip2))
	// fmt.Println(utf8.RuneCountInString(tip1))
	// fmt.Println(utf8.RuneCountInString(tip2))

	theme := "狙击 start"

	for i := 0; i < len(theme); i++ {
		fmt.Printf("ascii: %c %d\n", theme[i], theme[i])
	}

	for _, v := range theme {
		fmt.Printf("ascii: %c %d\n", v, v)
	}

	tracer := "死神来了,死神bye bye"
	index := strings.Index(tracer, ",")
	pos := strings.Index(tracer[index:], "死神")
	fmt.Println(index, pos, "\n", tracer[index+pos:])

	fmt.Println(tracer[strings.LastIndex(tracer, "死神"):])

}
