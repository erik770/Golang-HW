package main

import (
	"github.com/erik770/Golang-HW/pkg/calc"
	"github.com/erik770/Golang-HW/pkg/readwrite"
	"log"
)

func main() {
	expression, err := readwrite.ScanFromInput()
	if err != nil {
		log.Fatalf("READ ERR")
	}

	res, err := calc.Calculate(expression)
	if err != nil {
		log.Fatalf("CALC ERR")
		return
	}

	err = readwrite.WriteToOutput(res)
	if err != nil {
		log.Fatalf("WRITE ERR")
	}
}
