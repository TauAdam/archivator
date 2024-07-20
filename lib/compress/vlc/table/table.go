package table

import "strings"

type EncodingTable map[rune]string

// Generator is an interface for encoding table generators
type Generator interface {
	NewTable(text string) EncodingTable
}
type decodingTree struct {
	Value string
	Left  *decodingTree
	Right *decodingTree
}

// add adds a new node to the decoding tree
// based on the code and value. if the code is '0' the node will be added to the left, otherwise to the right
func (t *decodingTree) add(code string, value rune) {
	currentNode := t
	for _, bit := range code {
		switch bit {
		case '0':
			if currentNode.Left == nil {
				currentNode.Left = &decodingTree{}
			}
			currentNode = currentNode.Left
		case '1':
			if currentNode.Right == nil {
				currentNode.Right = &decodingTree{}
			}
			currentNode = currentNode.Right
		}
	}
	currentNode.Value = string(value)
}

// Decode decodes the input string based on the decoding tree
func (t *decodingTree) Decode(str string) string {
	var buf strings.Builder

	currentNode := t
	for _, char := range str {
		if currentNode.Value != "" {
			buf.WriteString(currentNode.Value)
			currentNode = t
		}
		switch char {
		case '0':
			currentNode = currentNode.Left
		case '1':
			currentNode = currentNode.Right
		}
	}
	if currentNode.Value != "" {
		buf.WriteString(currentNode.Value)
		currentNode = t
	}
	return buf.String()
}

// Decode decodes the input string based on the encoding table
func (t EncodingTable) decodingTree() decodingTree {
	res := decodingTree{}

	for char, code := range t {
		res.add(code, char)
	}
	return res
}
