package main

import "fmt"

func main() {

}

type Profile struct {
	Name    string
	Age     int
	Married bool
}

type queryKey struct {
	Name string
	Age  int
}

var mapper = make(map[queryKey]*Profile)

func buildIndex(list []*Profile) {
	for _, profile := range list {
		key := queryKey{
			Name: profile.Name,
			Age:  profile.Age,
		}
		mapper[key] = profile
	}
}

func queryData(name string, age int) {
	key := queryKey{name, age}
	result, ok := mapper[key]
	if ok {
		fmt.Println(result)
	} else {
		fmt.Println("no found")
	}
}
