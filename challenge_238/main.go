package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

const fileName string = "dict.txt"
const noOfGuesses int = 4

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}

func main() {
	words := getLinesFromFile(fileName)

	difficulty, err := strconv.Atoi(os.Args[1])
	failOnError(err, "Failed to parse difficulty.")

	filteredWords := filterBasedOnLength(words, difficulty)
	var shuffledWords []string

	if len(filteredWords) > 1 {
		shuffledWords = shuffleWords(filteredWords)
	} else {
		fmt.Println("Too few words!")
		os.Exit(1)
	}

	fmt.Println("Guess the password!")
	correctAnswer := getCorrectAnswer(shuffledWords)
	printWords(shuffledWords)

	for i := 0; i < noOfGuesses; i++ {
		guess := getPlayerGuess()
		matches := getNoOfMatchingChars(correctAnswer, guess)

		fmt.Printf("\nMatches: %d", matches)
	}
}

func getLinesFromFile(path string) []string {
	var lines []string
	file, err := os.Open(path)
	failOnError(err, "Failed to open file")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, strings.ToLower(scanner.Text()))
	}

	return lines
}

func filterBasedOnLength(words []string, length int) []string {
	var filteredWords []string

	for _, word := range words {
		if len(word) == length {
			filteredWords = append(filteredWords, word)
		}
	}

	return filteredWords
}

func shuffleWords(words []string) []string {
	rand.Seed(time.Now().UTC().UnixNano())
	wordCount := len(words)

	for i := 0; i < wordCount; i++ {
		randIndex := rand.Intn(wordCount)
		temp := words[randIndex]
		words[randIndex] = words[i]
		words[i] = temp
	}

	return words
}

func getCorrectAnswer(words []string) string {
	rand.Seed(time.Now().UTC().UnixNano())
	return words[rand.Intn(len(words))]
}

func printWords(words []string) {
	for _, word := range words {
		fmt.Printf("- %s \n", word)
	}
}

func getPlayerGuess() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("\nGuess:")
	guess, err := reader.ReadString('\n')
	failOnError(err, "Error while parsing user input.")

	return guess
}

func getNoOfMatchingChars(correctWord, wordToCompare string) int {
	var matchingChars int

	fmt.Printf("Length of correct word: %d \n Length of player guess: %d", len(correctWord), len(wordToCompare))

	if len(correctWord) != len(wordToCompare) {
		fmt.Println("Word is of incorrect length!")
	} else {
		for i := 0; i < len(wordToCompare); i++ {
			if wordToCompare[i] == correctWord[i] {
				matchingChars++
			}
		}
	}

	return matchingChars
}
