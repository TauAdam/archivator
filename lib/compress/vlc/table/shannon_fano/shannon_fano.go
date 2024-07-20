package shannon_fano

import "github.com/TauAdam/archivator/lib/compress/vlc/table"

type Generator struct {
}

// NewTable returns a new encoding table based on the Shannon-Fano algorithm
func (g Generator) NewTable(text string) table.EncodingTable {
	statistics := newCharStat(text)

	//	TODO: encoding table generation
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
