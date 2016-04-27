package main

import (
	"bytes"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const populationSize = 10
const tournamentSize = 5
const solutionWord = "Hello, world!"
const mutationRate = 0.06
const crossRate = 0.5

func main() {
	rand.Seed(time.Now().UnixNano())
	pop := populate(populationSize)
	maxFitness := getFitness(solutionWord, solutionWord)
	generation := 0
	for getFitness(getMostFit(pop), solutionWord) < maxFitness {
		pop = evolve(pop)
		fmt.Printf("Generation %d: %s\n", generation, getMostFit(pop))
		generation++
	}

	fmt.Printf("Generation %d: %s\n", generation, getMostFit(pop))
}

func populate(size int) []string {
	population := []string{}
	for i := 0; i < populationSize; i++ {
		population = append(population, getRandomText(size))
	}

	return population
}

func evolve(pop []string) []string {
	newPopulation := []string{}
	for i := 0; i < populationSize; i++ {
		wordOne := tournament(pop)
		wordTwo := tournament(pop)
		newWord := crossOver(wordOne, wordTwo)
		newPopulation = append(newPopulation, newWord)
	}

	for i := 0; i < populationSize; i++ {
		newPopulation[i] = mutate(newPopulation[i])
	}

	return newPopulation
}

func tournament(pop []string) string {
	tournamentPopulation := []string{}

	for i := 0; i < tournamentSize; i++ {
		randIndex := rand.Intn(len(pop))
		tournamentPopulation = append(tournamentPopulation, pop[randIndex])
	}

	return getMostFit(tournamentPopulation)
}

func getRandomText(n int) string {
	word := []rune{}

	for i := 0; i < len(solutionWord); i++ {
		word = append(word, getRandomLetter())
	}

	return string(word)
}

func getRandomLetter() rune {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ,! ")
	return letterRunes[rand.Intn(len(letterRunes))]
}

func getMostFit(population []string) string {
	fitness := 0
	fittestWord := ""
	for _, w := range population {
		wordFitness := getFitness(w, solutionWord)
		if wordFitness > fitness {
			fitness = wordFitness
			fittestWord = w
		}
	}

	return fittestWord
}

func crossOver(s1, s2 string) string {
	var textBuffer bytes.Buffer

	for index := 0; index < len(s1); index++ {
		if rand.Float64() <= crossRate {
			textBuffer.WriteByte(s1[index])
		} else {
			textBuffer.WriteByte(s2[index])
		}
	}

	return textBuffer.String()
}

func getFitness(word, solution string) int {
	if len(word) != len(solution) {
		return 0
	}

	fitness := 0
	for i := 0; i < len(word); i++ {
		if string(word[i]) == string(solution[i]) {
			fitness += 7
		}

		if strings.Contains(solution, string(word[i])) {
			fitness += 3
		}
	}

	return fitness
}

func mutate(text string) string {
	mutatedText := text

	for index := 0; index < len(mutatedText); index++ {
		if rand.Float64() <= mutationRate {
			letter := getRandomLetter()
			mutatedText = mutatedText[:index] + string(letter) + mutatedText[index+1:]
		}
	}

	return mutatedText
}
