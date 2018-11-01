package main

import (
	"fmt"
	"log"
)

type BadSymbolError int

func (e BadSymbolError) Error() string {
	return fmt.Sprint("Error, bad symbol ", int(e))
}

func Error(err error) {
	if(err != nil) {
		log.Fatal(err)
	}
}
