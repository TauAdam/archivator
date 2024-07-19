package vlc

import (
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"
)

const hexChunkSeparator = " "

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
			buf.WriteString(hexChunkSeparator)
		}
	}
	return buf.String()
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

func NewHexChunks(str string) HexChunks {
	chunks := strings.Split(str, hexChunkSeparator)
	res := make(HexChunks, 0, len(chunks))

	for _, chunk := range chunks {
		res = append(res, HexChunk(chunk))
	}

	return res
}
