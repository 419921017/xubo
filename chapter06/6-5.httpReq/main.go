package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func main() {
	client := &http.Client{}

	req, err := http.NewRequest("POST", "http://www.163.com/", strings.NewReader("key=value"))

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
		return
	}

	req.Header.Add("User-Agent", "myClient")

	res, err := client.Do(req)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
		return
	}

	data, _ := ioutil.ReadAll(res.Body)

	fmt.Println(string(data))

	defer res.Body.Close()

}
