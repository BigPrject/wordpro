package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const size int = 26

type trieNode struct {
	//size of possible characters in trie

	children [size]*trieNode

	terminal bool
}

type Trie struct {
	root *trieNode
}

// create trie with emtpy
func StartTrie() *Trie {
	result := &Trie{root: &trieNode{}}
	return result

}

// insert a word and it will add it to the trie

func (t *Trie) insert(word string) {

	wordLength := len(word)
	currentNode := t.root
	for i := 0; i < wordLength; i++ {
		charIndex := word[i] - 'a'
		if currentNode.children[charIndex] == nil {

			var node *trieNode = &trieNode{}
			currentNode.children[charIndex] = node

		}
		currentNode = currentNode.children[charIndex]

	}
	currentNode.terminal = true
}

func ImportWords() *Trie {
	TestTrie := StartTrie()
	file, err := os.Open("words.txt")
	if err != nil {
		fmt.Println(err.Error())
	}
	reader := bufio.NewScanner(file)

	for reader.Scan() {
		TestTrie.insert(reader.Text())

	}
	log.Println("Loaded Words!!")
	return TestTrie

}
