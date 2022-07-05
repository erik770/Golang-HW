package readwrite

import (
	"bufio"
	"fmt"
	"os"
)

func ScanFromInput() (scannedString string, openErr error) {
	input := os.Stdin
	if filepath := os.Args[1]; filepath != "" {
		input, openErr = os.Open(filepath)
		if openErr != nil {
			return scannedString, openErr
		}
		defer input.Close()
	}
	scanner := bufio.NewScanner(input)
	scanner.Scan()
	scannedString = scanner.Text()
	return scannedString, openErr
}

func WriteToOutput(str string) (err error) {
	output := os.Stdout
	if filepath := os.Args[2]; filepath != "" {
		output, err = os.Create(filepath)
		if err != nil {
			return err
		}
		defer output.Close()
	}
	_, err = fmt.Fprintln(output, str)
	if err != nil {
		return err
	}
	return err
}
