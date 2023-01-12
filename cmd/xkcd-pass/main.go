package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"math"
	"math/rand"
	"os"
	"strings"
	"time"
)

var (
	dictionaryPath string
	separator      string
	wordCount      int
	maxWordLength  int
	minWordLength  int
	verbose        bool
)

func init() {
	flag.StringVar(&dictionaryPath, "d", "", "path to file with dictionary of words")
	flag.StringVar(&separator, "s", "-", "separator")
	flag.IntVar(&wordCount, "c", 4, "number of words to use")
	flag.IntVar(&maxWordLength, "M", 5, "max word length")
	flag.IntVar(&minWordLength, "m", 2, "min word length")
	flag.BoolVar(&verbose, "v", false, "logs information on screen")
}

func PrintDefaultsWithError(errorMessage string) {
	log.Printf("invalid input parameters: %v", errorMessage)
	fmt.Println("Usage: xkcd-pass [OPTIONS]")
	flag.PrintDefaults()
	os.Exit(1)
}

func main() {
	flag.Parse()
	// flag validation
	if wordCount <= 0 {
		PrintDefaultsWithError("number of words should be > 0")
	}
	if maxWordLength <= 0 {
		PrintDefaultsWithError("max word length must be > 0")
	}
	if minWordLength <= 0 {
		PrintDefaultsWithError("min word length must be > 0")
	}
	if !(minWordLength < maxWordLength) {
		PrintDefaultsWithError("min word length must be lower than max word length")
	}
	if dictionaryPath == "" {
		PrintDefaultsWithError("dictionary path is required")
	}

	// read word list from file
	file, err := os.Open(dictionaryPath)
	if err != nil {
		log.Fatalf("unable to open file: %v", err)
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}
	allWords := strings.Split(string(content), "\n")

	// selecting first only valid words
	var validWords []string
	for _, candidate := range allWords {
		if len(candidate) < minWordLength {
			continue
		}
		if len(candidate) > maxWordLength {
			continue
		}
		candidate = strings.ToLower(candidate)
		validWords = append(validWords, candidate)
	}

	// seeding with random number
	randomSeed := time.Now().UTC().UnixNano()
	rand.Seed(randomSeed)

	// selecting the words that we will use in the password
	var chosenWords []string
	for i := 0; i < wordCount; i++ {
		randomIndex := rand.Intn(len(validWords))
		chosenWords = append(chosenWords, validWords[randomIndex])
	}

	fmt.Printf("%v", strings.Join(chosenWords, "-"))
	if verbose {
		fmt.Println()
		fmt.Printf("all words %v\n", len(allWords))
		fmt.Printf("sample space %v\n", len(validWords))
		entropy := float64(wordCount) * math.Log2(float64(len(validWords)))
		fmt.Printf("entropy %v\n", math.Round(entropy))
	}
}
