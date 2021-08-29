package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
)

func main() {
	h := &hello{
		endpoint: "http://localhost:8080/",
	}
	SetFuncField(h)
	msg, _ := h.FuncField("golang")
	fmt.Println(msg)
}

func SetFuncField(val interface{}) {
	v := reflect.ValueOf(val) // 获得指针
	ele := v.Elem()           // 获得指针对应的结构体
	t := ele.Type()           // 获得结构体的类型
	num := t.NumField()
	for i := 0; i < num; i++ {
		f := ele.Field(i)
		if f.CanSet() {
			fn := func(args []reflect.Value) (results []reflect.Value) {
				name := args[0].Interface().(string)
				client := http.Client{}
				response, err := client.Get("http://localhost:8080/" + name)
				if err != nil {
					fmt.Printf("%+v\n", err)
					return []reflect.Value{reflect.ValueOf(""), reflect.ValueOf(err)}
				}
				println(response.Status)
				data, err := ioutil.ReadAll(response.Body)
				if err != nil {
					fmt.Printf("%+v\n", err)
					return []reflect.Value{reflect.ValueOf(""), reflect.ValueOf(err)}
				}
				return []reflect.Value{reflect.ValueOf(string(data)), reflect.Zero(reflect.TypeOf(new(error)).Elem())}
			}

			f.Set(reflect.MakeFunc(f.Type(), fn))

		}
	}
}

type HelloService interface {
	SayHello(name string) (string, error)
}

type hello struct {
	endpoint  string
	FuncField func(name string) (string, error)
}

func (h hello) SayHello(name string) (string, error) {
	return "", nil
}
