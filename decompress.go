package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)


func getDecompressDictionary() map[int][]byte{
	dictSize := 128
	dictionary := make(map[int][]byte, dictSize)
	for i := 0; i < dictSize; i++ {
		dictionary[i] = []byte{byte(i)}
	}
	return dictionary
}


func decompress(compressedFileName string) (string, error) {
	dictSize := 128
	dictionary:=getDecompressDictionary()
	var result strings.Builder
	var bytesString []byte
	var currentBytesString []byte

	decompressedFile, err := os.Create("decompressedFile.txt")
	if(err != nil){
		fmt.Println("err")
	}
	writer:= bufio.NewWriter(decompressedFile)
	defer func() {
		writer.Flush()
		decompressedFile.Close()
	}()
	compressedFile, err := os.Open(compressedFileName)
	Error(err)
	reader := csv.NewReader(compressedFile)
	var str string
	for {
		inputChar,  err := reader.Read()
		if err == io.EOF {
			break
		}
		Error(err)
		str = strings.Join(inputChar, "")
		int64, err := strconv.ParseInt(str, 10, 32)
		currentSymbolCode := int(int64)
		if currentSymbol, ok := dictionary[currentSymbolCode];
			ok { currentBytesString = currentSymbol[:len(currentSymbol):len(currentSymbol)]
		} else if currentSymbolCode == dictSize && len(bytesString) > 0 {
			currentBytesString = append(bytesString, bytesString[0])
		} else {
			return result.String(), BadSymbolError(currentSymbolCode)
		}
		writer.Write(currentBytesString)
		addToDecompressDictionary(&bytesString, &currentBytesString, dictionary, &dictSize)
		str = ""
	}
	return result.String(), nil
}


func addToDecompressDictionary(bytesString *[]byte,currentBytesString *[]byte, dictionary map[int][]byte, dictSize *int,){
	if len(*bytesString) > 0 {
		*bytesString = append(*bytesString, (*currentBytesString)[0])
		dictionary[*dictSize] = *bytesString
		*dictSize++
	}
	*bytesString = *currentBytesString
}
