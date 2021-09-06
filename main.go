package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	filepath := "/Users/jinpenghe/Downloads/test.txt"
	file, err := os.OpenFile(filepath, os.O_RDWR, 0666)
	if err != nil {
		fmt.Println("Open file error!", err)
		return
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		panic(err)
	}
	var size = stat.Size()
	fmt.Println("file size =", size)

	var resList []string
	buf := bufio.NewReader(file)
	for {
		line, err := buf.ReadString('\n')
		line = strings.TrimSpace(line)
		//fmt.Println(line)
		resList = append(resList, line)
		if err != nil {
			if err == io.EOF {
				fmt.Println("File read ok!")
				break
			} else {
				fmt.Println("Read file error!", err)
				return
			}
		}
	}

	resMap := make(map[string][]int)
	for i := 0; i < len(resList); i++ {
		resMap[resList[i]] = append(resMap[resList[i]], i)
	}
	//fmt.Println(resMap)

	for k, v := range resMap {
		if len(v) < 2 {
			continue
		}
		for i := 1; i < len(v); i++ {
			gap := v[i] - v[i-1]
			if gap > 1 {
				fmt.Println(k)
				fmt.Println(v)
			}
		}
	}
}
