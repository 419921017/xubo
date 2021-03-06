package main

import "fmt"

func main() {
	list := []*Profile{
		{Name: "张三", Age: 30, Married: true},
		{Name: "李四", Age: 21},
		{Name: "王五", Age: 21},
	}
	buildIndex(list)
	queryData("张三", 30)
}

type Profile struct {
	Name    string
	Age     int
	Married bool
}

func simpleHash(str string) (ret int) {
	for i := 0; i < len(str); i++ {
		c := str[i]
		ret += int(c)
	}
	return
}

type classicQueryKey struct {
	Name string
	Age  int
}

func (c *classicQueryKey) hash() int {
	return simpleHash(c.Name) + c.Age*1000000
}

var mapper = make(map[int][]*Profile)

func buildIndex(list []*Profile) {
	for _, profile := range list {
		key := classicQueryKey{profile.Name, profile.Age}
		existValue := mapper[key.hash()]
		existValue = append(existValue, profile)
		mapper[key.hash()] = existValue
	}
}

func queryData(name string, age int) {
	keyToQuery := classicQueryKey{name, age}
	resultList := mapper[keyToQuery.hash()]
	for _, result := range resultList {
		if result.Name == name && result.Age == age {
			fmt.Println(result)
			return
		}
	}
	fmt.Println("no found")
}
