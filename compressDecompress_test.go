package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math/rand"
	"testing"
	"time"
)
import "os"
func TestCompressDecompress(t *testing.T) {
	f := RandStringWrite(150)//150000000)
	file, err:= os.Open(f)
	if err != nil {
		log.Fatal(err)
	}
	compressed:=compress(file)
	defer file.Close()
	_, err = decompress(compressed)
	if err != nil {
		log.Fatal(fmt.Sprintf("decompress %s\n", err))
	}
	resultFile, err1 := os.Open("decompressedFile.txt")
	testFile, err2 := os.Open("testFile.txt")
	if err1 != nil || err2 != nil {
		if err1 == io.EOF && err2 == io.EOF {
			return
		} else if err1 == io.EOF || err2 == io.EOF {
			log.Fatal("Files are not the same size")
		} else {
			log.Fatal(fmt.Sprintf("Bad symbol"))
		}
	}

	defer resultFile.Close()
	defer testFile.Close()




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
	file, err := os.Create(fileName)
	if err != nil{
		fmt.Println("Unable to open file:", err)
		os.Exit(1)
	}
	writer := bufio.NewWriter(file)
	defer func() {
		writer.Flush()
		file.Close()
	}()
	var letters [128]byte

	for i, _ := range letters{
		letters[i] = byte(i)
	}

	rand.Seed(time.Now().UnixNano())
	for i := 0; i<n; i++{
		writer.Write([]byte(string(letters[rand.Intn(len(letters))])))
	}

	return fileName
}


