package vlc

import (
	"strings"
	"unicode"
)

func Encode(str string) string {
	_ = prepareText(str)

	// TODO: encode to binary
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
