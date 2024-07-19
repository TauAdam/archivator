package vlc

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

type encodingTable = map[rune]string
type BinaryChunk string

func (c BinaryChunk) ToHex() HexChunk {
	num, err := strconv.ParseUint(string(c), 2, chunkSize)
	if err != nil {
		panic("failed to convert binary to int" + err.Error())
	}

	res := strings.ToUpper(fmt.Sprintf("%X", num))

	if len(res) == 1 {
		res = "0" + res
	}
	return HexChunk(res)
}

type BinaryChunks []BinaryChunk

func (c BinaryChunks) ToHex() HexChunks {
	res := make(HexChunks, 0, len(c))

	for _, chunk := range c {
		hexChunk := chunk.ToHex()
		res = append(res, hexChunk)
	}

	return res
}

type HexChunk string
type HexChunks []HexChunk

func (c HexChunks) ToString() string {
	const separator = " "

	switch len(c) {
	case 0:
		return ""
	case 1:
		return string(c[0])
	}
	var buf strings.Builder

	for i, chunk := range c {
		buf.WriteString(string(chunk))
		if i < len(c)-1 {
			buf.WriteString(separator)
		}
	}
	return buf.String()
}

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
		if (i+1)%chunkSize == 0 {
			res = append(res, BinaryChunk(buf.String()))
			buf.Reset()
		}
	}

	if buf.Len() != 0 {
		lastChunk := buf.String()
		lastChunk += strings.Repeat("0", chunkSize-len(lastChunk))
		res = append(res, BinaryChunk(lastChunk))
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
