package main

import (
	"bufio"
	"fmt"
	"github.com/erik770/Golang-HW/pkg/calc"
	"log"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	expression := scanner.Text()
	res, err := calc.Calculate(expression)
	if err != nil {
		log.Fatalf("CALC ERR", err)
		return
	}
	fmt.Println(res)
}
