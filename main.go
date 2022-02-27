package main

import (
	"github.com/erik770/Golang-HW/pkg/uniq"
	"log"
)

func main() {
	input, output, flags := uniq.ReadOptions()
	data, _ := uniq.ReadFile(input)
	err := uniq.ValidateFlags(flags)
	if err != nil {
		log.Fatal(err)
	}
	uniq.WriteFile(output, uniq.Uniq(data, flags), flags)

}
