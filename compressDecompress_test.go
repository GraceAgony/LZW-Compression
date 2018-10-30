package main

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)
import "os"

func TestCompressDecompress(t *testing.T) {
	fileName:="text.txt"
	file, err := os.Open(fileName)
	if err != nil{
		fmt.Println("Unable to open file:", err)
		os.Exit(1)
	}
	data := RandString(1000000)
	file.Write([]byte(data))

	defer file.Close()
	compressed:=compress(data)
	result, err:= decompress(compressed)
	if result != data {
		t.Error(fmt.Sprintf("Expected data, %s\n got %s", data, result))
	}
}


func RandString(n int) string {
	var letters [128]byte

	for i, _ := range letters{
		letters[i] = byte(i)
	}

	rand.Seed(time.Now().UnixNano())
	arr := make([]byte, n)
	for i := range arr {
		arr[i] = letters[rand.Intn(len(letters))]
	}
	return string(arr)
}





