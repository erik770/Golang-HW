package main

import (
	"github.com/erik770/Golang-HW/pkg/uniq"
	"log"
)

func main() {
	input, output, flags := uniq.ReadOptions()

	data, err := uniq.ReadFile(input)
	if err != nil {
		log.Fatalf("READ ERR")
		return
	}

	err = uniq.ValidateFlags(flags)
	if err != nil {
		log.Fatalf("VALIDATE ERR")
		return
	}

	err = uniq.WriteFile(output, uniq.Uniq(data, flags))
	if err != nil {
		log.Fatalf("WRITE ERR")
		return
	}
}
