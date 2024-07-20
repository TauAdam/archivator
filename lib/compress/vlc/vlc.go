package vlc

import (
	"strings"
	"unicode"
)

type encodingTable map[rune]string

const chunkSize = 8

// Encode encodes the input string to VLC
// Where all the magic happens
func Encode(str string) string {
	text := prepareText(str)

	binStr := EncodeToBinary(text)

	chunks := splitByChunks(binStr, chunkSize)

	return chunks.ToHex().ToString()
}

// prepareText removes all upper case characters from the input string
// and converts it to lowercase(uppercase letters to `! + lower case letter`.
// i.g.: "Hello, World!" -> "!hello, world!"
func prepareText(str string) string {
	var buf strings.Builder

	for _, char := range str {
		if unicode.IsUpper(char) {
			buf.WriteRune('!')
			buf.WriteRune(unicode.ToLower(char))
		} else {
			buf.WriteRune(char)
		}
	}
	return buf.String()
}

// restoreText converts the input string to Uppercase
// if the character is followed by '!'
func restoreText(str string) string {
	var buf strings.Builder

	var isUpper bool
	for _, char := range str {
		if isUpper {
			buf.WriteRune(unicode.ToUpper(char))
			isUpper = false
			continue
		}
		if char == '!' {
			isUpper = true
			continue
		} else {
			buf.WriteRune(char)
		}
	}
	return buf.String()
}

// EncodeToBinary encodes the input string to binary without spaces
func EncodeToBinary(str string) string {
	var buf strings.Builder
	for _, char := range str {
		buf.WriteString(encodeCharToBinary(char))
	}
	return buf.String()
}

func encodeCharToBinary(char rune) string {
	table := newEncodingTable()
	res, ok := table[char]
	if !ok {
		panic("unknown character" + string(char))
	}
	return res
}

func newEncodingTable() encodingTable {
	return encodingTable{
		' ': "11",
		't': "1001",
		'n': "10000",
		's': "0101",
		'r': "01000",
		'd': "00101",
		'!': "001000",
		'c': "000101",
		'm': "000011",
		'g': "0000100",
		'b': "0000010",
		'v': "00000001",
		'k': "0000000001",
		'q': "000000000001",
		'e': "101",
		'o': "10001",
		'a': "011",
		'i': "01001",
		'h': "0011",
		'l': "001001",
		'u': "00011",
		'f': "000100",
		'p': "0000101",
		'w': "0000011",
		'y': "0000001",
		'j': "000000001",
		'x': "00000000001",
		'z': "000000000000",
	}
}

// Decode decodes the input string from VLC
// "09 10 A7 50" -> "gopher"
func Decode(encodedText string) string {
	hexChunks := NewHexChunks(encodedText)

	binChunks := hexChunks.ToBinary()

	binString := binChunks.Join()

	tree := newEncodingTable().DecodingTree()

	return tree.Decode(binString)
}
