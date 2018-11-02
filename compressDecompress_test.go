package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"testing"
	"time"
)

func TestCompressDecompress(t *testing.T) {
	f := RandStringWrite(172000000)
	file, err:= os.Open(f)
	Error(err)
	compressed:=compress(file)
	_, err = decompress(compressed)
	Error(err)
	resultFile, err1 := os.Open("decompressedFile.txt")
	testFile, err2 := os.Open("testFile.txt")
	Error(err1)
	Error(err2)
	CompareFiles(resultFile, testFile)
	defer resultFile.Close()
	defer testFile.Close()
	defer file.Close()
}

func CompareFiles(resultFile *os.File, testFile *os.File){
	readerResultFile := bufio.NewReader(resultFile)
	readerTestFile := bufio.NewReader(testFile)
	for{
		charResultFile, _, err1 := readerResultFile.ReadRune()
		charTestFile, _, err2 := readerTestFile.ReadRune()
		if err1 != nil || err2 != nil {
			if err1 == io.EOF && err2 == io.EOF {
				return
			} else if err1 == io.EOF || err2 == io.EOF {
				log.Fatal("Files are not the same size")
			} else {
				log.Fatal(fmt.Sprintf("Bad symbol"))
			}
		}
		if(charResultFile != charTestFile){
			log.Fatal(fmt.Sprintf("Expected symbol, %s\n got %s",
				string(charTestFile), string(charResultFile)))
		}
	}
}

func RandStringWrite(n int) string {
	fileName:="testFile.txt"
	var letters [128]byte
	file, err := os.Create(fileName)
	Error(err)
	writer := bufio.NewWriter(file)
	defer func() {
		writer.Flush()
		file.Close()
	}()
	for i, _ := range letters{
		letters[i] = byte(i)
	}
	rand.Seed(time.Now().UnixNano())
	for i := 0; i<n; i++{
		writer.Write([]byte(string(letters[rand.Intn(len(letters))])))
	}
	return fileName
}


