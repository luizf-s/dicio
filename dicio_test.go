package main

import (
	"bytes"
	"io/ioutil"
	"testing"
)

var testCases = []struct {
	word, input string
	expect      WordData
}{
	{
		word:  "léxico",
		input: "test_assets/lexico.html",
		expect: WordData{
			Word: "léxico",
			Descriptions: []Description{
				Description{
					WordClass: "substantivo masculino",
				},
				Description{
					WordClass: "adjetivo",
				},
			},
			Etimology: "Etimologia (origem da palavra léxico). Do grego leksikós.é.ón.",
		},
	},
	{
		word: "ferrar",
		input: "test_assets/ferrar.html",
		expect: WordData {
			Word: "ferrar",
			Descriptions: []Description {
				Description{ WordClass: "verbo transitivo" },
			},
		},
	},
}

func TestGetWordData(t *testing.T) {
	for _, testCase := range testCases {
		file, err := ioutil.ReadFile(testCase.input)
		if err != nil {
			t.Errorf("Error trying to open test case file: %s", err)
		}
		htmlReader := bytes.NewReader(file)
		wordData := GetWordData(htmlReader, testCase.word)
		if !isSame(wordData, testCase.expect) {
			t.Errorf("GetWordData(%s) = %v, want %v", testCase.input, wordData, testCase.expect)
		}
	}
}

func isSame(left, right WordData) bool {
	return left.Etimology == right.Etimology &&
		left.Word == right.Word &&
		hasSameWordClasses(left, right)
}

func hasSameWordClasses(left, right WordData) bool {
	for i, description := range left.Descriptions {
		if right.Descriptions[i].WordClass != description.WordClass {
			return false
		}
	}
	return true
}
