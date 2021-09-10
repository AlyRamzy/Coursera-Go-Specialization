package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func isCorrect(s string) bool {
	str := strings.ToLower(s)

	// if exists, remove trailing Line Feed (added on Windows and on Unix/OSX)
	if strings.HasSuffix(str, "\n") {
		str = strings.TrimSuffix(str, "\n")
	}

	if strings.HasPrefix(str, "i") == false {
		return false
	}

	if strings.Contains(str, "a") == false {
		return false
	}

	if strings.HasSuffix(str, "n") == false {
		return false
	}

	return true
}
func main() {
	// fmt.Scan place only characters before first SPACE into a string variable so we have to use buffer.
	// When user presses ENTER, that adds End Of Line (CRLF (\r\n) on Windows and LF(\n) on Unix/OSX) at the end of the string!

	stdinReader := bufio.NewReader(os.Stdin)

	if inputString, err := stdinReader.ReadString('\n'); err != nil {
		fmt.Println("Invalid input. Error: ", err)
	} else {
		if isCorrect(inputString) {
			fmt.Println("Found!")
		} else {
			fmt.Println("Not found!")
		}
	}
}
