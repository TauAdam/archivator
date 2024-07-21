package shannon_fano

import (
	"github.com/TauAdam/archivator/lib/compress/vlc/table"
	"math"
	"sort"
)

type Generator struct {
}

// NewTable returns a new encoding table based on the Shannon-Fano algorithm
func (g Generator) NewTable(text string) table.EncodingTable {
	statistics := newCharStat(text)
	_ = statistics

	//	TODO: encoding table generation
	return nil
}

type Stats map[rune]int

// newCharStat returns a map of character frequencies in the text
func newCharStat(text string) Stats {
	res := make(Stats)
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

func build(stats Stats) encodingTable {
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
	//assignCodes()
	return nil
}

func assignCodes(codes []code) {
	if len(codes) <= 1 {
		return
	}
	findBestPosition(codes)

}

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

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}
