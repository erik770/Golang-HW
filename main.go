package main

import (
	"github.com/erik770/Golang-HW/pkg/uniq"
	"log"
)

func main() {
	flags, err := uniq.MyParseFlags()
	if err != nil {
		log.Fatalf("WRONG FLAGS ERR")
	}
	input, output := uniq.ReadInputOutputPaths()

	data, err := uniq.ReadFromInput(input)
	if err != nil {
		log.Fatalf("READ ERR")
	}

	err = uniq.WriteFile(output, uniq.Uniq(data, flags))
	if err != nil {
		log.Fatalf("WRITE ERR")
	}
}
