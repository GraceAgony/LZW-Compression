package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func compress(inputText string) []int {
	dictSize := 128
	dictionary:= getCompressDictionary()
	var result []int
	var bytesString []byte
	for i := 0; i < len(inputText); i++ {
		currentSymbol := inputText[i];
		currentBytesString := append(bytesString, currentSymbol);
		if _, ok := dictionary[string(currentBytesString)];
		ok { bytesString = currentBytesString;
		} else {
			addToCompressDictionary(&result, dictionary, &bytesString, &dictSize, &currentBytesString, currentSymbol);
		}
	}
	if len(bytesString) > 0 {
		result = append(result, dictionary[string(bytesString)])
	}
	return result
}

func addToCompressDictionary(result *[]int, dictionary map[string]int,
	bytesString *[]byte, dictSize *int,
	currentBytesString *[]byte,
	currentSymbol byte){
	*result = append(*result, dictionary[string(*bytesString)])
	dictionary[string(*currentBytesString)] = *dictSize
	*dictSize++
	*currentBytesString = ([]byte{currentSymbol})
	*bytesString = *currentBytesString
}

func addToDecompressDictionary(bytesString *[]byte,currentBytesString *[]byte, dictionary map[int][]byte, dictSize *int,){
	if len(*bytesString) > 0 {
		*bytesString = append(*bytesString, (*currentBytesString)[0])
		dictionary[*dictSize] = *bytesString
		*dictSize++
	}
	*bytesString = *currentBytesString
}

func getCompressDictionary() map[string]int{
	dictSize := 128

	dictionary := make(map[string]int, dictSize)
	for i := 0; i < dictSize; i++ {
		dictionary[string([]byte{byte(i)})] = i;
	}

	return dictionary
}

func getDecompressDictionary() map[int][]byte{
	dictSize := 128
	dictionary := make(map[int][]byte, dictSize)
	for i := 0; i < dictSize; i++ {
		dictionary[i] = []byte{byte(i)}
	}
	return dictionary
}

type BadSymbolError int

func (e BadSymbolError) Error() string {
	return fmt.Sprint("Error, bad symbol ", int(e))
}


func decompress(compressed []int) (string, error) {
	dictSize := 128
	dictionary:=getDecompressDictionary()
	var result strings.Builder
	var bytesString []byte
	var currentBytesString []byte
	for _, currentSymbol := range compressed {
		if currentSymbolCode, ok := dictionary[currentSymbol];
		ok { currentBytesString = currentSymbolCode[:len(currentSymbolCode):len(currentSymbolCode)]
		} else if currentSymbol == dictSize && len(bytesString) > 0 {
			currentBytesString = append(bytesString, bytesString[0])
		} else {
			return result.String(), BadSymbolError(currentSymbol)
		}
		result.Write(currentBytesString)

		addToDecompressDictionary(&bytesString, &currentBytesString, dictionary, &dictSize)
	}
	return result.String(), nil
}


func readFile(fileName string) string{
	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Print(err)
	}
	return string(file)
}

func main() {
	text := readFile("text.txt")
	compressed := compress(text)
	fmt.Println(compressed)
	decompressed, err := decompress(compressed)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(decompressed)
}