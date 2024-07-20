package vlc

import (
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"
)

type BinaryChunk string

func (c BinaryChunk) Byte() byte {
	num, err := strconv.ParseUint(string(c), 2, chunkSize)
	if err != nil {
		panic("Error converting binary to byte" + err.Error())
	}
	return byte(num)
}

type BinaryChunks []BinaryChunk

func (c BinaryChunks) Join() string {
	var buf strings.Builder
	for _, chunk := range c {
		buf.WriteString(string(chunk))
	}
	return buf.String()

}

func (c BinaryChunks) Bytes() []byte {
	res := make([]byte, 0, len(c))
	for _, chunk := range c {
		res = append(res, chunk.Byte())

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

func NewBinChunks(data []byte) BinaryChunks {
	res := make(BinaryChunks, 0, len(data))

	for _, code := range data {
		res = append(res, NewBinChunk(code))
	}

	return res
}

func NewBinChunk(code byte) BinaryChunk {
	return BinaryChunk(fmt.Sprintf("%08b", code))
}
