package uniq

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"log"
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

func ValidateFlags(flags Flags) (err error) {
	if (flags.unique && flags.count) || (flags.unique && flags.duplicate) || (flags.count && flags.duplicate) || flags.fieldsSkip < -1 || flags.skipChars < -1 {
		log.Printf("недопустимый набор флагов, гайдлайн по флагам:\n" +
			"uniq [-count | -duplicate | -unique] [-ignoreReg] [-fieldsSkip num (pos number)]" +
			" [-skipChars chars (pos number)] [input_file [output_file]]\n\n")
		return errors.New("NO VALIDE FLAGS")
	}

	return nil
}
func MyParseFlags() (flags Flags, err error) {
	flag.BoolVar(&flags.count, "c", false, "Count number of repeats")
	flag.BoolVar(&flags.duplicate, "d", false, "Only duplicate strings")
	flag.BoolVar(&flags.unique, "u", false, "Only unique strings")
	flag.BoolVar(&flags.ignoreReg, "i", false, "Ignore register")
	flag.IntVar(&flags.fieldsSkip, "f", -1, "Skip first num_fields")
	flag.IntVar(&flags.skipChars, "s", -1, "Skip first num_chars in sting")
	flag.Parse()
	if err := ValidateFlags(flags); err != nil {
		return flags, err
	}
	return flags, nil
}

func ReadInputOutputPaths() (input, output string) {
	flag.Parse()
	input = flag.Arg(0)
	output = flag.Arg(1)
	return input, output
}

func ReadFromInput(filePath string) (res []string, err error) {
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

func WriteFile(fileName string, data []string) (err error) {
	output := os.Stdout

	if fileName != "" {
		output, err = os.Create(fileName)
		if err != nil {
			return err
		}
		defer output.Close()
	}

	for _, elem := range data {
		_, err = fmt.Fprintln(output, elem)
		if err != nil {
			break
		}

	}
	return nil
}

func registerIgnoreOption(data []string) []string {
	for index := range data {
		data[index] = strings.ToLower(data[index])
	}
	return data
}

func stringsIgnoreOption(data []string, stringsToIgnore int) []string {
	data = data[stringsToIgnore:]
	return data
}

func charsIgnoreOption(data []string, charsToIgnore int) []string {
	stringSkipCounter := 0
	for _, str := range data {
		if len(str) > charsToIgnore {
			break
		}
		stringSkipCounter++
		charsToIgnore -= len(str)
	}

	data = data[stringSkipCounter:]
	data[0] = data[0][charsToIgnore:]
	return data
}

func createCounterMap(data []string) (counterMap map[string]int) {
	counterMap = make(map[string]int)

	var prevString string
	for index, str := range data {
		if index == 0 {
			prevString = str
			counterMap[str] = 0
			continue
		}

		_, exist := counterMap[str]

		if !exist {
			counterMap[str] = 0
		}
		if prevString == str {
			counterMap[str]++
		}

		prevString = str
	}
	return counterMap
}

func addNumOfRepeatsToString(repeatingString string, numberOfRepeats int) string {
	return strconv.Itoa(numberOfRepeats) + " " + repeatingString
}

func Uniq(data []string, flags Flags) []string {
	var res []string

	if flags.ignoreReg {
		data = registerIgnoreOption(data)
	}
	if flags.fieldsSkip != -1 {
		data = stringsIgnoreOption(data, flags.fieldsSkip)
	}
	if flags.skipChars != -1 {
		data = charsIgnoreOption(data, flags.skipChars)
	}

	switch {
	case flags.count:
		counterMap := make(map[string]int)

		var prevString string
		for index, str := range data {
			if index == 0 {
				prevString = str
				continue
			}

			if prevString == str {
				counterMap[str]++
				prevString = str
				continue
			}

			counterMap[prevString]++
			res = append(res, addNumOfRepeatsToString(prevString, counterMap[prevString]))
			counterMap[prevString] = 0

			prevString = str
		}
		counterMap[prevString]++
		resStr := strconv.Itoa(counterMap[prevString]) + " " + prevString
		res = append(res, resStr)

	case flags.unique:
		counterMap := createCounterMap(data)
		for _, str := range data {
			if counterMap[str] == 0 {
				res = append(res, str)
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
		for _, str := range data {
			if prevString != str {
				res = append(res, str)
			}
			prevString = str
		}
	}
	return res
}
