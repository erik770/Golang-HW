package main

import (
	"log"
)

func main() {
	input, output, flags := myParseFlags()
	data, _ := ReadFile(input)
	err := ValidateFlags(flags)
	if err != nil {
		log.Fatal(err)
	}
	WriteFile(output, ApplyFlags(data, flags), flags)

}
