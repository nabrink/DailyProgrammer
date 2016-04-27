package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp/syntax"
	"strings"
)

const fileName string = "palindrom.txt"

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}

func getText(path string) string {
	var lines string
	file, err := os.Open(path)
	failOnError(err, "Failed to open file")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = lines + strings.ToLower(scanner.Text())
	}

	return lines
}

func main() {
	text := getText(fileName)
	text = strings.ToLower(strings.Map(func(r rune) rune {
		if syntax.IsWordChar(r) {
			return r
		}
		return -1
	}, text))

	length := len(text)

	if length%2 == 0 {
		firstPart := text[0 : length/2]
		secondPart := reverse(text[length/2 : length])

		if firstPart != secondPart {
			fmt.Println("Text is not a palindrom")
		} else {
			fmt.Println("Text is a palindrom")
		}

	} else {
		firstPart := text[0 : length/2]
		secondPart := reverse(text[(length/2)+1 : length])

		if firstPart != secondPart {
			fmt.Println("Text is not a palindrom")
		} else {
			fmt.Println("Text is a palindrom")
		}
	}

}
