package shannon_fano

import (
	"fmt"
	"github.com/TauAdam/archivator/lib/compress/vlc/table"
	"math"
	"sort"
	"strings"
)

type Generator struct {
}

// NewTable returns a new encoding table based on the Shannon-Fano algorithm
func (g Generator) NewTable(text string) table.EncodingTable {
	occurrences := newCharOccurrences(text)

	return build(occurrences).Export()
}

type Occurrences map[rune]int

// newCharOccurrences returns a map of character frequencies in the text
func newCharOccurrences(text string) Occurrences {
	res := make(Occurrences)
	for _, char := range text {
		res[char]++
	}
	return res
}

type code struct {
	Char     rune
	Quantity int
	Bits     uint32
	Size     int
}
type encodingTable map[rune]code

// Export returns the encoding table as a map of runes to binary strings
func (t encodingTable) Export() map[rune]string {
	res := make(map[rune]string)

	for key, value := range t {
		bytesStr := fmt.Sprintf("%b", value.Bits)

		// 11 -> 0011
		if sizeDiff := value.Size - len(bytesStr); sizeDiff > 0 {
			bytesStr = strings.Repeat("0", sizeDiff) + bytesStr
		}

		res[key] = bytesStr
	}
	return res
}

// build creates a new encoding table based on the character occurrences
func build(stats Occurrences) encodingTable {
	codes := make([]code, 0, len(stats))

	for char, quantity := range stats {
		codes = append(codes, code{Char: char, Quantity: quantity})
	}
	sort.Slice(codes, func(i, j int) bool {
		if codes[i].Quantity != codes[j].Quantity {
			return codes[i].Quantity > codes[j].Quantity
		}
		return codes[i].Char < codes[j].Char
	})
	assignCodes(codes)

	res := make(encodingTable)
	for _, c := range codes {
		res[c.Char] = c
	}

	return res
}

// assignCodes assigns binary codes to the characters based on their frequencies
func assignCodes(codes []code) {
	if len(codes) <= 1 {
		return
	}

	pos := findBestPosition(codes)

	for i := 0; i < len(codes); i++ {
		// move to the left
		codes[i].Bits <<= 1
		codes[i].Size++

		if i >= pos {
			codes[i].Bits |= 1

		}
	}
	assignCodes(codes[:pos])
	assignCodes(codes[pos:])
}

// findBestPosition finds the best position to split the codes
func findBestPosition(codes []code) int {
	sum := 0
	for _, code := range codes {
		sum += code.Quantity
	}

	left := 0
	prevDiff := math.MaxInt
	pos := 0
	for i := 0; i < len(codes)-1; i++ {
		left += codes[i].Quantity
		right := sum - left

		diff := abs(left - right)

		if diff >= prevDiff {
			break
		}
		prevDiff = diff
		pos = i + 1

	}
	return pos
}

// abs returns module of the integer
func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}
