package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/dmartzol/xkcd-pass/wordlists"
)

var (
	wordLanguage string
	separator    string
	wordCount    int
)

func init() {
	flag.StringVar(&wordLanguage, "l", "en", "language")
	flag.StringVar(&separator, "s", "-", "separator")
	flag.IntVar(&wordCount, "c", 4, "number of words to use")
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

	// chose a wordList
	wordList, ok := wordlists.Wordlists[wordLanguage]
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

	fmt.Printf("%v\n", strings.Join(words, "-"))
}
