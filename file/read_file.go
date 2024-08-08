package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	filePath := "tmp/data.json"
	content, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(content))
}
