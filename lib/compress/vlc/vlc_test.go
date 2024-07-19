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
			str:  "!hi",
			want: "20 D2",
		},
		{
			name: "case with spaces",
			str:  "!hi !bob",
			want: "20 D3 90 0A 20 80",
		},
		{
			name: "basic case",
			str:  "gopher",
			want: "09 10 A7 50",
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
