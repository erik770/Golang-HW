package uniq

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Flags struct {
	count      bool
	duplicate  bool
	unique     bool
	ignoreReg  bool
	fieldsSkip int
	skipChars  int
}

func ReadOptions() (input, output string, flags Flags) {
	flag.BoolVar(&flags.count, "c", false, "Count number of repeats")
	flag.BoolVar(&flags.duplicate, "d", false, "Only duplicate strings")
	flag.BoolVar(&flags.unique, "u", false, "Only unique strings")
	flag.BoolVar(&flags.ignoreReg, "i", false, "Ignore register")
	flag.IntVar(&flags.fieldsSkip, "f", -1, "Skip first num_fields")
	flag.IntVar(&flags.skipChars, "s", -1, "Skip first num_chars in sting")
	flag.Parse()
	input = flag.Arg(0)
	output = flag.Arg(1)
	return input, output, flags
}

func ReadFile(filePath string) (res []string, err error) {
	var openErr error
	input := os.Stdin

	if filePath != "" {
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
		_, err := fmt.Fprintln(os.Stderr, "reading standard input:", err)
		if err != nil {
			return nil, err
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
	if flags.unique && flags.count || flags.unique && flags.duplicate || flags.count && flags.duplicate {
		return errors.New("недопустимый набор флагов, гайдлайн по флагам:\n" +
			"uniq [-count | -duplicate | -unique] [-ignoreReg] [-fieldsSkip num] [-skipChars chars] [input_file [output_file]]\n\n")
	}
	return nil
}

func IgnoringOption(data []string, RegisterIgnore bool, StringsToIgnore int, CharsToIgnore int) (res []string) {
	if RegisterIgnore {
		for _, elem := range data {
			elem = strings.ToLower(elem)
		}
	}

	if StringsToIgnore > 0 {
		data = data[StringsToIgnore:]
	}

	stringSkipCounter := 0
	if CharsToIgnore > 0 {
		for _, elem := range data {
			if len(elem) > CharsToIgnore {
				break
			}
			stringSkipCounter++
			CharsToIgnore -= len(elem)
		}
		data = data[stringSkipCounter:]
		data[0] = data[0][CharsToIgnore:]
	}
	return data
}

func createCounterMap(data []string) (counterMap map[string]int) {
	counterMap = make(map[string]int)

	var prevString string
	for index, elem := range data {
		if index == 0 {
			prevString = elem
			continue
		}

		if prevString == elem {
			counterMap[elem]++
		}
		prevString = elem
	}
	return counterMap
}

func Uniq(data []string, flags Flags) []string {
	var res []string

	if flags.fieldsSkip != -1 || flags.skipChars != -1 || flags.ignoreReg {
		data = IgnoringOption(data, flags.ignoreReg, flags.fieldsSkip, flags.skipChars)
	}

	switch {
	case flags.count:
		counterMap := make(map[string]int)

		var prevString string
		for index, elem := range data {
			if index == 0 {
				prevString = elem
				continue
			}

			if prevString == elem {
				counterMap[elem]++
			} else {
				counterMap[prevString]++
				resStr := strconv.Itoa(counterMap[prevString]) + " " + prevString
				res = append(res, resStr)
				counterMap[prevString] = 0
			}
			prevString = elem
		}
		resStr := strconv.Itoa(counterMap[prevString]) + " " + prevString
		res = append(res, resStr)

	case flags.unique:
		counterMap := createCounterMap(data)
		for _, elem := range data {
			if counterMap[elem] == 0 {
				res = append(res, elem)
			}
		}

	case flags.duplicate:
		counterMap := createCounterMap(data)
		for fields, counter := range counterMap {
			if counter != 0 {
				res = append(res, fields)
			}
		}

	default:
		var prevString string
		for _, elem := range data {
			if prevString != elem {
				res = append(res, elem)
			}
			prevString = elem
		}
	}

	//}
	return res
}
