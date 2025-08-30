package utilities

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

func ReadStrings(outputText, errorMsg string) (result string, err error) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println(outputText)

	if ok := scanner.Scan(); !ok {
		return "", errors.New(errorMsg)
	}
	text := scanner.Text()
	return text, nil
}
