package main

/*
import (
	"bufio"
	"fmt"
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
func startTrie() *Trie {
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

// search

func (t *Trie) search(word string) bool {

	currentNode := t.root
	wordLength := len(word)
	for i := 0; i < wordLength; i++ {
		charIndex := word[i] - 'a'
		if currentNode.children[charIndex] == nil {
			return false
		}
		currentNode = currentNode.children[charIndex]

	}
	return currentNode.terminal

}

func main() {

	testTrie := startTrie()
	file, err := os.Open("words.txt")
	if err != nil {
		fmt.Println(err.Error())
	}
	reader := bufio.NewScanner(file)

	for reader.Scan() {
		testTrie.insert(reader.Text())

	}
	var board [4][4]uint8
	boggle := "drpxgarimroeaitt"
	var words []string
	var visited [4][4]bool

	//index value for the string
	index := 0
	for i := 0; i < len(board); i++ {

		for j := 0; j < len(board[i]); j++ {
			board[i][j] = boggle[index]
			index++
		}

	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			dfssearch(testTrie, testTrie.root, board, i, j, visited, "", &words)
		}
	}
	fmt.Printf("words: %v\n", words)

}

func dfssearch(t *Trie, currentNode *trieNode, board [4][4]uint8, row int, column int, visited [4][4]bool, path string, foundWords *[]string) {
	if (row >= 4 || column >= 4) || (row < 0 || column < 0) {
		return
	}

	if visited[row][column] {
		return
	}
	visited[row][column] = true

	letter := board[row][column] - 'a'

	if currentNode.children[letter] == nil {
		visited[row][column] = false
		return
	}

	path += string(board[row][column])

	if currentNode.terminal {
		*foundWords = append(*foundWords, path)

	}

	currentNode = currentNode.children[letter]

	searchSpace := [8][2]int{
		{1, 0}, {-1, 0}, {0, 1}, {0, -1}, // up, down, left, right
		{1, 1}, {1, -1}, {-1, 1}, {-1, -1}, // diagonals
	}
	for _, d := range searchSpace {
		nRow := int(row) + d[0]

		nCol := int(column) + d[1]

		dfssearch(t, currentNode, board, nRow, nCol, visited, path, foundWords)

	}

	visited[row][column] = false

}
*/
