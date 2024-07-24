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

func New(tlbGen table.Generator) EncoderDecoder {
	return EncoderDecoder{
		tableGenerator: tlbGen,
	}
}

const chunkSize = 8

// Encode encodes the input string to VLC
// Where all the magic happens
func (ed EncoderDecoder) Encode(str string) []byte {
	newTable := ed.tableGenerator.NewTable(str)

	binStr := EncodeToBinary(str, newTable)

	return buildEncodedFile(newTable, binStr)
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

// Decode decodes the input bytes to string
func (ed EncoderDecoder) Decode(encodedBytes []byte) string {
	encodingTable, data := parseArchive(encodedBytes)

	return encodingTable.Decode(data)
}

// parseArchive parses the input bytes to the encoding table and the encoded data
func parseArchive(data []byte) (table.EncodingTable, string) {
	const (
		tableSizeBytesCount = 4
		dataSizeBytesCount  = 4
	)

	tableSizeBinary, data := data[:tableSizeBytesCount], data[tableSizeBytesCount:]
	data = data[dataSizeBytesCount:]

	tableSize := binary.BigEndian.Uint32(tableSizeBinary)

	tblBinary, data := data[:tableSize], data[tableSize:]

	tbl := decodeGobTable(tblBinary)

	body := NewBinChunks(data).Join()
	//log.Printf("data size: %d, body size: %d", dataSize, len(body))
	return tbl, body
}

// decodeGobTable decodes the input bytes to the encoding table
func decodeGobTable(tblBinary []byte) table.EncodingTable {
	var tbl table.EncodingTable
	err := gob.NewDecoder(bytes.NewReader(tblBinary)).Decode(&tbl)
	if err != nil {
		log.Fatalf("deserialization error: %v", err)
	}
	return tbl
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
