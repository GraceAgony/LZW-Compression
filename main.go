package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func compress(file *os.File) string{
	dictSize := 128
	dictionary:= getCompressDictionary()
	//var result []int
	var bytesString []byte
	var fileName = "compressedFile.csv"
	compressedFile, err := os.Create(fileName)
	if(err != nil){
		fmt.Println("err")
	}
	writer:= csv.NewWriter(compressedFile)
	defer func() {
		writer.Flush()
		compressedFile.Close()
	}()
	reader := bufio.NewReader(file)
	for{
		inputChar, _, err := reader.ReadRune()
		if(err != nil){
			break
		}
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
		fmt.Println(fmt.Sprintf("%d", dictionary[string(bytesString)]))
		if(err != nil){
			fmt.Println("err")
		}
	}
	return fileName
}

func addToCompressDictionary(writer *csv.Writer, dictionary map[string]int,
	bytesString *[]byte, dictSize *int,
	currentBytesString *[]byte,
	currentSymbol byte){
	//*result = append(*result, dictionary[string(*bytesString)])
	err := writer.Write(strings.Split(fmt.Sprintf("%d", dictionary[string(*bytesString)]), ","))
	fmt.Println(fmt.Sprintf("%d", dictionary[string(*bytesString)]))
	if(err != nil){
		fmt.Println("err")
	}
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
	reader := csv.NewReader(compressedFile)
	var str string
	for {
		inputChar,  err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		//if(inputChar != ","){
		//	str = str + fmt.Sprintf("%d", inputChar);
		//	continue;
		//}
		str = strings.Join(inputChar, "")
		int64, err := strconv.ParseInt(str, 10, 32)
		currentSymbol := int(int64)
		if currentSymbolCode, ok := dictionary[currentSymbol];
			ok { currentBytesString = currentSymbolCode[:len(currentSymbolCode):len(currentSymbolCode)]
		} else if currentSymbol == dictSize && len(bytesString) > 0 {
			currentBytesString = append(bytesString, bytesString[0])
		} else {
			return result.String(), BadSymbolError(currentSymbol)
		}
		//result.Write(currentBytesString)
		fmt.Println(currentBytesString)
		addToDecompressDictionary(&bytesString, &currentBytesString, dictionary, &dictSize)

		str = ""
	}
	return result.String(), nil
}



func main() {
	file, err := os.Open("text.txt")
	compressed := compress(file)
	defer file.Close()
	_, err = decompress(compressed)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println(decompressed)
}