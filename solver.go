package main

import (
	"encoding/json"
	"fmt"
)

type Word struct {
	Word string  `json:"word:"`
	Path [][]int `json:"path:"`
}

type wordResponse struct {
	Words []Word `json:"words"`
}

func solve(wordTrie *Trie, boggle string) string {

	var board [4][4]uint8
	var words = make(map[string]Word)
	var pathTaken [][]int
	var visited [4][4]bool

	//index value for the string
	index := 0
	for i := 0; i < len(board); i++ {

		for j := 0; j < len(board[i]); j++ {
			board[i][j] = boggle[index]
			index++
		}

	}

	// to loop through every starting position on the board
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			dfssearch(wordTrie, wordTrie.root, board, i, j, visited, "", pathTaken, &words)
		}
	}
	var resp wordResponse
	for _, wordType := range words {

		resp.Words = append(resp.Words, wordType)

	}
	output, err := json.MarshalIndent(resp, "", "")
	if err != nil {
		fmt.Println("error occured", err)
	}
	return string(output)
}

func dfssearch(t *Trie, currentNode *trieNode, board [4][4]uint8, row int, column int, visited [4][4]bool, path string, pathTaken [][]int, foundWords *map[string]Word) {
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
	temp := []int{row, column}
	pathTaken = append(pathTaken, temp)
	if currentNode.terminal {

		current := Word{Word: path, Path: pathTaken}
		(*foundWords)[path] = current
		return
	}

	currentNode = currentNode.children[letter]

	searchSpace := [8][2]int{
		{1, 0}, {-1, 0}, {0, 1}, {0, -1}, // up, down, left, right
		{1, 1}, {1, -1}, {-1, 1}, {-1, -1}, // diagonals
	}
	for _, d := range searchSpace {
		nRow := int(row) + d[0]

		nCol := int(column) + d[1]

		dfssearch(t, currentNode, board, nRow, nCol, visited, path, pathTaken, foundWords)

	}

	visited[row][column] = false

}
