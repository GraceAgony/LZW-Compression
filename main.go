package main

import (
	"log"
	"os"
)


func main() {
	file, err := os.Open("text.txt")
	compressed := compress(file)
	defer file.Close()
	_, err = decompress(compressed)
	if err != nil {
		log.Fatal(err)
	}
}

