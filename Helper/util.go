package main

import (
	"io/ioutil"
	"log"
	"os"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// read 负责读取文件
// 这是一个通用的方法
func read(path string) []byte {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	return data
}

func write(path, content string) {
	err := ioutil.WriteFile(path, []byte(content), 0755)
	if err != nil {
		log.Fatal(err)
	}
}
