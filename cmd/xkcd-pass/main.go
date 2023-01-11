package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/dmartzol/xkcd-pass/wordlists"
)

var (
	dictionaryPath string
	wordLanguage   string
	wordCount      int
)

func init() {
	flag.StringVar(&dictionaryPath, "d", "", "input dictionry path")
	flag.StringVar(&wordLanguage, "l", "", "language")
	flag.IntVar(&wordCount, "c", 4, "number of words to use")
	//flag.IntVar(&iterations, "n", 1000, "number of iterations")
}

func PrintDefaultsWithError(errorMessage string) {
	log.Printf("invalid input parameters: %v", errorMessage)
	fmt.Println("Usage: poly [OPTIONS] -o output")
	flag.PrintDefaults()
	os.Exit(1)
}

func main() {
	flag.Parse()
	// flag validation
	if wordCount <= 0 {
		PrintDefaultsWithError("number of words should be > 0")
	}

	// chose a wordList
	wordList, ok := wordlists.Wordlists["en"]
	if !ok {
		log.Fatalf("list not found")
	}

	randomSeed := time.Now().UTC().UnixNano()
	rand.Seed(randomSeed)

	var words []string
	for i := 0; i < wordCount; i++ {
		randomIndex := rand.Intn(len(wordList))
		words = append(words, wordList[randomIndex])
	}
	fmt.Printf("%v-%v\n", words[0], words[1])
}
