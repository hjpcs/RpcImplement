package main

import (
	"errors"
	"fmt"
)

func main() {
	values := []interface{}{3, 5,
		4, 2, 7}
	newValues, err := Add(values, 6, 4)
	if err != nil {
		fmt.Print(err.Error())
		return
	}
	// output is 3, 5, 4, 2, 6, 7
	fmt.Println("after add element, the data is")
	for i, value := range newValues {
		fmt.Printf("index is %d, value is %d\n", i, value)
	}
	newValues = Delete(newValues, 2)
	// output is 3, 5, 2, 6, 7
	fmt.Println("after delete element, the data is")
	for i, value := range newValues {
		fmt.Printf("index is %d, value is %d\n", i, value)
	}
}

func Add(values []interface{}, val interface{}, index int) ([]interface{}, error) {
	// judge whether the index is valid
	if index < 0 || index > len(values) {
		return nil, errors.New("index is invalid")
	}
	// put the previous value into res
	var res []interface{}
	for i := 0; i < index; i++ {
		v := values[i]
		res = append(res, v)
	}
	// put the insert value into res
	res = append(res, val)
	// put the remaining value into res
	for i := index; i < len(values); i++ {
		v := values[i]
		res = append(res, v)
	}
	return res, nil
}

func Delete(values []interface{}, index int) []interface{} {
	if index < 0 || index > len(values) {
		return values
	}
	var res []interface{}
	for i := 0; i < len(values); i++ {
		if index == i {
			continue
		}
		v := values[i]
		res = append(res, v)
	}
	return res
}
