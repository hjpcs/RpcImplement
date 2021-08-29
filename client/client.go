package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	client := http.Client{}
	response, err := client.Get("http://localhost:8080/goland")
	if err != nil {
		fmt.Printf("%v\n", err)
		fmt.Printf("%+v\n", err)
		fmt.Printf("%#v\n", err)
		return
	}
	println(response.Status)
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("%+v", err)
		return
	}
	fmt.Printf("%v\n", string(data))
	//s := "你好呀" + "12"
	//fmt.Printf("%d\n", len(s))
}
