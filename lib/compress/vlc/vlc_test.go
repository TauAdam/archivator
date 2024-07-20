package vlc

import (
	"reflect"
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
		want []byte
	}{
		{
			name: "basic case",
			str:  "Hi",
			want: []byte{32, 210},
		},
		{
			name: "case with spaces",
			str:  "Hi Bob",
			want: []byte{32, 211, 144, 10, 32, 128},
		},
		{
			name: "basic case",
			str:  "gopher",
			want: []byte{9, 16, 167, 80},
		},
		{
			name: "quote",
			str:  "Consistency is key",
			want: []byte{32, 88, 193, 82, 179, 96, 40, 29, 43, 128, 52, 8},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			encoder := New()
			if got := encoder.Encode(tt.str); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Encode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDecode(t *testing.T) {
	tests := []struct {
		name  string
		bytes []byte
		want  string
	}{
		{
			name:  "basic case",
			bytes: []byte{32, 210},
			want:  "Hi",
		},
		{
			name:  "case with spaces",
			bytes: []byte{32, 211, 144, 10, 32, 128},
			want:  "Hi Bob",
		},
		{name: "basic case",
			bytes: []byte{9, 16, 167, 80},
			want:  "gopher",
		},
		{
			name:  "Quote",
			bytes: []byte{32, 88, 193, 82, 179, 96, 40, 29, 43, 128, 52, 8},
			want:  "Consistency is key",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			decoder := New()
			if got := decoder.Decode(tt.bytes); got != tt.want {
				t.Errorf("Decode() = %v, want %v", got, tt.want)
			}
		})
	}
}
