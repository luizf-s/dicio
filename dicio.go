package main

import (
	"fmt"
	"net/http"
	"os"
)

type Description struct {
	WordClass   string
	Definitions []string
}

type WordData struct {
	Descriptions    []Description
	Etimology, Word string
}

func main() {
	word := os.Args[1]
	resp, err := http.Get(fmt.Sprintf("http://dicio.com.br/%s/", word))
	defer resp.Body.Close()

	if err != nil {
		fmt.Printf("Error on request: %s\n", err)
		os.Exit(1)
	}

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Error trying to get %s definition: %v\n", word, err)
		os.Exit(1)
	}

	wordData := GetWordData(resp.Body, word)
	PrintWordData(wordData)
}
