package vlc

import (
	"testing"
)

func Test_prepareText(t *testing.T) {
	tests := []struct {
		name string
		str  string
		want string
	}{
		{
			name: "basic case",
			str:  "Hello, World!",
			want: "!hello, !world!",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := prepareText(tt.str); got != tt.want {
				t.Errorf("prepareText() = %v, want =  %v", got, tt.want)
			}
		})
	}
}

func TestEncodeToBinary(t *testing.T) {
	tests := []struct {
		name string
		str  string
		want string
	}{
		{
			name: "basic case",
			str:  "!hi",
			want: "001000001101001",
		},
		{
			name: "case with spaces",
			str:  "!hi !bob",
			want: "001000001101001110010000000010100010000010",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EncodeToBinary(tt.str); got != tt.want {
				t.Errorf("EncodeBin() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEncode(t *testing.T) {
	tests := []struct {
		name string
		str  string
		want string
	}{
		{
			name: "basic case",
			str:  "Hi",
			want: "20 D2",
		},
		{
			name: "case with spaces",
			str:  "Hi Bob",
			want: "20 D3 90 0A 20 80",
		},
		{
			name: "basic case",
			str:  "gopher",
			want: "09 10 A7 50",
		},
		{
			name: "quote",
			str:  "Consistency is key",
			want: "20 58 C1 52 B3 60 28 1D 2B 80 34 08",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Encode(tt.str); got != tt.want {
				t.Errorf("Encode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDecode(t *testing.T) {
	tests := []struct {
		name        string
		encodedText string
		want        string
	}{
		{
			name:        "basic case",
			encodedText: "20 D2",
			want:        "Hi",
		},
		{
			name:        "case with spaces",
			encodedText: "20 D3 90 0A 20 80",
			want:        "Hi Bob",
		},
		{name: "basic case",
			encodedText: "09 10 A7 50",
			want:        "gopher",
		},
		{
			name:        "Quote",
			encodedText: "20 58 C1 52 B3 60 28 1D 2B 80 34 08",
			want:        "Consistency is key",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Decode(tt.encodedText); got != tt.want {
				t.Errorf("Decode() = %v, want %v", got, tt.want)
			}
		})
	}
}
