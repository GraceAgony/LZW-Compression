package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strings"
)


func getCompressDictionary() map[string]int{
	dictSize := 128
	dictionary := make(map[string]int, dictSize)
	for i := 0; i < dictSize; i++ {
		dictionary[string([]byte{byte(i)})] = i;
	}

	return dictionary
}


func compress(file *os.File) string{
	dictSize := 128
	dictionary:= getCompressDictionary()
	var bytesString []byte
	var fileName = "compressedFile.csv"
	compressedFile, err := os.Create(fileName)
	Error(err)
	writer:= csv.NewWriter(compressedFile)
	defer func() {
		writer.Flush()
		compressedFile.Close()
	}()
	reader := bufio.NewReader(file)
	for{
		inputChar, _, err := reader.ReadRune()
		if err == io.EOF {
			break
		}
		Error(err)
		currentSymbol := byte(inputChar);
		currentBytesString := append(bytesString, currentSymbol);
		if _, ok := dictionary[string(currentBytesString)];
			ok { bytesString = currentBytesString;
		} else {
			addToCompressDictionary(writer, dictionary, &bytesString, &dictSize, &currentBytesString, currentSymbol);
		}
	}
	if len(bytesString) > 0 {
		err := writer.Write(strings.Split(fmt.Sprintf("%d", dictionary[string(bytesString)]),","))
		Error(err)
	}
	return fileName
}

func getNextChar()  error{

}

func addToCompressDictionary(writer *csv.Writer, dictionary map[string]int,
	bytesString *[]byte, dictSize *int,
	currentBytesString *[]byte,
	currentSymbol byte) {
		err := writer.Write(strings.Split(fmt.Sprintf("%d", dictionary[string(*bytesString)]), ","))
		Error(err)
		dictionary[string(*currentBytesString)] = *dictSize
		*dictSize++
		*currentBytesString = ([]byte{currentSymbol})
		*bytesString = *currentBytesString
}
