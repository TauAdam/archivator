package vlc

import (
	"bytes"
	"encoding/binary"
	"encoding/gob"
	"github.com/TauAdam/archivator/lib/compress/vlc/table"
	"log"
	"strings"
)

type EncoderDecoder struct {
	tableGenerator table.Generator
}

func New() EncoderDecoder {
	return EncoderDecoder{}
}

const chunkSize = 8

// Encode encodes the input string to VLC
// Where all the magic happens
func (ed EncoderDecoder) Encode(str string) []byte {
	newTable := ed.tableGenerator.NewTable(str)

	binStr := EncodeToBinary(str, newTable)

	chunks := splitByChunks(binStr, chunkSize)

	return chunks.Bytes()
}

func buildEncodedFile(tbl table.EncodingTable, data string) []byte {
	encodedTable := encodeTable(tbl)

	var buf bytes.Buffer
	buf.Write(encodeNumbers(len(encodedTable)))
	buf.Write(encodeNumbers(len(tbl)))
	buf.Write(encodedTable)
	buf.Write(splitByChunks(data, chunkSize).Bytes())

	return buf.Bytes()
}

func encodeTable(tbl table.EncodingTable) []byte {
	var buf bytes.Buffer
	err := gob.NewEncoder(&buf).Encode(tbl)
	if err != nil {
		log.Fatalf("serialization error: %v", err)
	}
	return buf.Bytes()
}
func encodeNumbers(num int) []byte {
	res := make([]byte, 4)
	binary.BigEndian.PutUint32(res, uint32(num))

	return res
}

// Decode decodes the input bytes from VLC
// "09 10 A7 50" -> "gopher"
func (_ EncoderDecoder) Decode(encodedBytes []byte) string {
	binString := NewBinChunks(encodedBytes).Join()

	tree := newEncodingTable().DecodingTree()

	return restoreText(tree.Decode(binString))
}

// EncodeToBinary encodes the input string to binary without spaces
func EncodeToBinary(str string, table table.EncodingTable) string {
	var buf strings.Builder
	for _, char := range str {
		buf.WriteString(encodeCharToBinary(char, table))
	}
	return buf.String()
}

func encodeCharToBinary(char rune, table table.EncodingTable) string {
	res, ok := table[char]
	if !ok {
		panic("unknown character" + string(char))
	}
	return res
}
