package vlc

type DecodingTree struct {
	Value string
	Left  *DecodingTree
	Right *DecodingTree
}

// Add adds a new node to the decoding tree
// based on the code and value. if the code is '0' the node will be added to the left, otherwise to the right
func (t *DecodingTree) Add(code string, value rune) {
	currentNode := t
	for _, bit := range code {
		switch bit {
		case '0':
			if currentNode.Left == nil {
				currentNode.Left = &DecodingTree{}
			}
			currentNode = currentNode.Left
		case '1':
			if currentNode.Right == nil {
				currentNode.Right = &DecodingTree{}
			}
			currentNode = currentNode.Right
		}
	}
	currentNode.Value = string(value)
}

func (t encodingTable) DecodingTree() DecodingTree {
	res := DecodingTree{}

	for char, code := range t {
		res.Add(code, char)
	}
	return res
}
