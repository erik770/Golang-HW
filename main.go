package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	// "strconv"
)

func ReadFile(filePath string) (res []string, err error) {
	var input *os.File
	var openErr error

	if filePath == "" {
		input = os.Stdin
	} else {
		input, openErr = os.Open(filePath)
		if openErr != nil {
			return res, openErr
		}
		defer input.Close()
	}

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		res = append(res, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		_, err2 := fmt.Fprintln(os.Stderr, "reading standard input:", err)
		if err2 != nil {
			return nil, err2
		}
	}
	return res, err
}

func WriteFile(fileName string, data []string, flags Flags) (err error) {
	var output *os.File

	if fileName != "" {
		output, err = os.Create(fileName)
		if err != nil {
			return err
		}
		defer output.Close()
	} else {
		output = os.Stdout
	}

	for _, elem := range data {
		_, err = fmt.Fprintln(output, elem)
		if err != nil {
			break
		}

	}
	return
}

func ValidateFlags(flags Flags) (err error) {
	if flags.u && flags.c || flags.u && flags.d || flags.c && flags.d {
		return errors.New("недопустимый набор флагов, гайдлайн по флагам:\n" +
			"uniq [-c | -d | -u] [-i] [-f num] [-s chars] [input_file [output_file]]\n\n")
	}
	return nil
}

// Пока думаю как сделать логику для флагов
// чтобы не для каждого флага своя а как-то оптимально
func ApplyFlags(data []string, flags Flags) []string {
	var res []string

	// if flags.c {
	// 	counter := 0
	// 	prevString := "init prv string"
	// 	for _, elem := range data {
	// 		if prevString == elem || prevString == "init prv string" {
	// 			counter++
	// 		} else {

	// 			res = append(res, strconv.Itoa(counter)+" "+elem)
	// 			counter = 1
	// 		}
	// 		if prevString == "\n" {
	// 			prevString = elem
	// 			continue
	// 		}
	// 		prevString = elem
	// 	}
	// }
	if !(flags.c || flags.u || flags.d || flags.i || flags.f != -1 || flags.s != -1) {
		var prevString string
		for _, elem := range data {
			if prevString != elem {
				res = append(res, elem)
			}
			prevString = elem
		}
	}
	return res
}

func main() {
	input, output, flags := myParseFlags()
	data, _ := ReadFile(input)
	err := ValidateFlags(flags)
	if err != nil {
		log.Fatal(err)
	}
	WriteFile(output, ApplyFlags(data, flags), flags)

}
