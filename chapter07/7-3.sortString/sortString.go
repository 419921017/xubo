package main

import (
	"fmt"
	"sort"
)

func main() {
	names1 := MyStringList{
		"3. Triple Kill",
		"5. Penta Kill",
		"2. Double Kill",
		"4. Quadra Kill",
		"1. First Blood",
	}
	names2 := sort.StringSlice{
		"3. Triple Kill",
		"5. Penta Kill",
		"2. Double Kill",
		"4. Quadra Kill",
		"1. First Blood",
	}
	sort.Sort(names1)
	sort.Sort(names2)
	for _, v := range names1 {
		fmt.Printf("%s\n", v)
	}
	fmt.Println()
	for _, v := range names2 {
		fmt.Printf("%s\n", v)
	}
}

type MyStringList []string

func (m MyStringList) Len() int {
	return len(m)
}

func (m MyStringList) Less(i, j int) bool {
	return m[i] < m[j]
}

func (m MyStringList) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}
