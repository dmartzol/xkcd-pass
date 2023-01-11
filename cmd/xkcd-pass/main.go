package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

var (
	dictionaryPath string
	wordLanguage   string
	separator      string
	wordCount      int
	wordLength     int
)

func init() {
	flag.StringVar(&dictionaryPath, "d", "en", "path to file with dictionary of words")
	flag.StringVar(&wordLanguage, "l", "en", "language")
	flag.StringVar(&separator, "s", "-", "separator")
	flag.IntVar(&wordCount, "c", 4, "number of words to use")
	flag.IntVar(&wordLength, "x", 6, "max word length")
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
	if dictionaryPath == "" {
		PrintDefaultsWithError("dictionary path is required")
	}

	// chose a wordList
	file, err := os.Open(dictionaryPath)
	if err != nil {
		log.Fatalf("unable to open file: %v", err)
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}
	wordList := strings.Split(string(content), "\n")

	randomSeed := time.Now().UTC().UnixNano()
	rand.Seed(randomSeed)

	var words []string
	for i := 0; i < wordCount; i++ {
		randomIndex := rand.Intn(len(wordList))
		words = append(words, wordList[randomIndex])
	}

	fmt.Printf("%v\n", strings.Join(words, "-"))
}
