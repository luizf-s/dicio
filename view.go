package main

import (
	"fmt"
)

func PrintWordData(wordData WordData) {
	fmt.Printf("%s\n", bold(wordData.Word))
	for _, description := range wordData.Descriptions {
		fmt.Printf("  %s\n", green(description.WordClass))
		for _, definition := range description.Definitions {
			fmt.Printf("  · %s\n", definition)
		}
		fmt.Println()
	}
	fmt.Printf("%s\n", cyan(wordData.Etimology))
}

func bold(raw string) string {
	return fmt.Sprintf("\u001b[1;91m%s\u001b[0;0m", raw)
}

func green(raw string) string {
	return fmt.Sprintf("\u001b[32m%s\u001b[0m", raw)
}

func cyan(raw string) string {
	return fmt.Sprintf("\u001b[36m%s\u001b[0m", raw)
}
