package main

import (
	"github.com/realm76/sergeant/pkg/parser"
	"log"
	"os"
)

func main() {
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}

	makeFile, err := parser.ParseFile(path)
	if err != nil {
		panic(err)
	}

	log.Printf("Makefile: %v\n", makeFile)
}
