package main

import (
	"fmt"
	"io"
	"strings"

	"golang.org/x/net/html"
)

func GetWordData(htmlRaw io.Reader, word string) WordData {
	node, err := html.Parse(htmlRaw)
	if err != nil {
		fmt.Printf("Error parsing html: %s\n", err)
	}
	description := findDescription(node)
	return parseDescription(description, word)
}

func findDescription(node *html.Node) *html.Node {
	if isDescription(node) {
		return node
	}
	for child := node.FirstChild; child != nil; child = child.NextSibling {
		rowElementSearch := findDescription(child)
		if isDescription(rowElementSearch) {
			return rowElementSearch
		}
	}
	return &html.Node{}
}

func isDescription(node *html.Node) bool {
	return node.Type == html.ElementNode &&
		node.Data == "p" &&
		hasItemPropDescription(node)
}

func hasItemPropDescription(node *html.Node) bool {
	for _, attr := range node.Attr {
		if attr.Key == "itemprop" && attr.Val == "description" {
			return true
		}
	}
	return false
}

func parseDescription(descriptionNode *html.Node, word string) WordData {
	var wordData WordData
	wordData.Word = word

	for element := descriptionNode.FirstChild; element != nil; element = element.NextSibling {
		if isEtimology(element) {
			wordData.Etimology = getInnerText(element)
		}
		if isWordClass(element) {
			var description Description
			description.WordClass = getInnerText(element)
			for elementInner := element.NextSibling;
				isDefinition(elementInner); // TODO: elementInner may be nihil
				elementInner = elementInner.NextSibling {
				if strings.ReplaceAll(getInnerText(elementInner), " ", "") != "" {
					description.Definitions = append(description.Definitions, getInnerText(elementInner))
				}
			}
			wordData.Descriptions = append(wordData.Descriptions, description)
		}
	}

	return wordData
}

func isWordClass(node *html.Node) bool {
	if node == nil {
		fmt.Println("isWordClass: nihil found")
	}
	return node.Data == "span" && hasClassCl(node)
}

func isDefinition(node *html.Node) bool {
	if node == nil {
		return false
	}
	return !isWordClass(node) && !isEtimology(node)
}

func isEtimology(node *html.Node) bool {
	return node.Data == "span" && hasClassEtim(node)
}

func hasClassEtim(node *html.Node) bool {
	for _, attr := range node.Attr {
		if attr.Key == "class" && attr.Val == "etim" {
			return true
		}
	}
	return false
}

func hasClassCl(node *html.Node) bool {
	for _, attr := range node.Attr {
		if attr.Key == "class" && attr.Val == "cl" {
			return true
		}
	}
	return false
}

func getInnerText(node *html.Node) string {
	var builder strings.Builder
	if node.Type == html.TextNode {
		return node.Data
	}
	for children := node.FirstChild; children != nil; children = children.NextSibling {
		builder.WriteString(getInnerText(children))
	}
	return builder.String()
}
