package main

import (
	"fmt"
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
}
