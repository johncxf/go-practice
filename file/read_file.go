package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

// ReadJson 读取 json 文件
func ReadJson(filePath string) map[string]interface{} {
	content, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}
	var jsonData map[string]interface{}
	err = json.Unmarshal(content, &jsonData)
	if err != nil {
		log.Fatal(err)
	}
	return jsonData
}

// ReadLineReader 逐行读取文件 bufio.NewReader
func ReadLineReader(filePath string) {
	start := time.Now()
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer file.Close()
	lineReader := bufio.NewReader(file)
	for {
		line, _, err := lineReader.ReadLine()
		if err == io.EOF {
			break
		}
		// 打印
		fmt.Println(string(line))
	}
	fmt.Println("ReadLineReader done, cost : ", time.Now().Sub(start))
}

// ReadLineScanner 逐行读取文件 bufio.NewScanner
func ReadLineScanner(filePath string) {
	start := time.Now()
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer file.Close()
	lineScanner := bufio.NewScanner(file)
	for lineScanner.Scan() {
		line := lineScanner.Text()
		// 打印
		fmt.Println(line)
	}
	fmt.Println("ReadLineScanner done, cost : ", time.Now().Sub(start))
}

// ReadBlock 分块读取
func ReadBlock(filePath string) {
	start := time.Now()
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	buffer := make([]byte, 1024)
	for {
		n, err := file.Read(buffer)
		// 控制条件,根据实际调整
		if err != nil && err != io.EOF {
			log.Println(err)
		}
		if n == 0 {
			break
		}
		// 打印出每次读取的文件块(字节数)
		fmt.Println(string(buffer[:n]))
	}
	fmt.Println("ReadBlock done, cost : ", time.Now().Sub(start))
}

func main() {
	filePath1 := "tmp/data.json"
	fmt.Println(ReadJson(filePath1))

	filePath2 := "tmp/data.txt"
	ReadLineReader(filePath2)
	ReadLineScanner(filePath2)
	ReadBlock(filePath2)
}
