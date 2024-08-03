package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	filePath := "tmp/data.txt"
	content, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(content))
}
