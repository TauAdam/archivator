package vlc

import (
	"strings"
	"unicode"
	"unicode/utf8"
)

type encodingTable = map[rune]string
type BinaryChunks []string

const chunkSize = 8

func Encode(str string) string {
	_ = prepareText(str)

	binStr := EncodeToBinary(str)

	// TODO: split by chunks
	_ = splitByChunks(binStr, chunkSize)
	return ""
}

// prepareText removes all upper case characters from the input string
// and converts it to lowercase(uppercase letters to `! + lower case letter`.
// i.g.: "Hello, World!" -> "!hello, world!"
func prepareText(text string) string {
	var buf strings.Builder
	for _, char := range text {
		if unicode.IsUpper(char) {
			buf.WriteRune('!')
			buf.WriteRune(unicode.ToLower(char))
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

// splitByChunks splits the input string by chunks of given size
// i.g.: "101010101", 3 -> ["101", "010", "101"]
func splitByChunks(binStr string, chunkSize int) BinaryChunks {
	// Better than len(binStr)
	numberOfRunes := utf8.RuneCountInString(binStr)
	amount := numberOfRunes / chunkSize

	if numberOfRunes%chunkSize != 0 {
		amount++
	}
	res := make(BinaryChunks, 0, amount)

	var buf strings.Builder

	for i, char := range binStr {
		buf.WriteString(string(char))
		if chunkSize == i+1 {
			res = append(res, buf.String())
			buf.Reset()
		}
	}

	if buf.Len() != 0 {
		lastChunk := buf.String()
		lastChunk += strings.Repeat("0", chunkSize-len(lastChunk))
		res = append(res, lastChunk)
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
