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

func Test_splitByChunks(t *testing.T) {
	type args struct {
		binStr    string
		chunkSize int
	}
	tests := []struct {
		name string
		args args
		want BinaryChunks
	}{
		{
			name: "basic case",
			args: args{
				binStr:    "001000001101001",
				chunkSize: 8,
			},
			want: BinaryChunks{"00100000", "11010010"},
		},
		{
			name: "case with chunk size 4",
			args: args{
				binStr:    "001000001101001",
				chunkSize: 4,
			},
			want: BinaryChunks{"0010", "0000", "1101", "0010"},
		},
		{
			name: "case with chunk size 2",
			args: args{
				binStr:    "001000001101001",
				chunkSize: 2,
			},
			want: BinaryChunks{"00", "10", "00", "00", "11", "01", "00", "10"},
		},
		{
			name: "case with chunk size 6",
			args: args{
				binStr:    "001000001101001",
				chunkSize: 6,
			},
			want: BinaryChunks{"001000", "001101", "001000"},
		},
		{
			name: "case with chunk size 3",
			args: args{
				binStr:    "001000001101001",
				chunkSize: 3,
			},
			want: BinaryChunks{"001", "000", "001", "101", "001"},
		},
		{
			name: "case with chunk size 5",
			args: args{
				binStr:    "001000001101001",
				chunkSize: 5,
			},
			want: BinaryChunks{"00100", "00011", "01001"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := splitByChunks(tt.args.binStr, tt.args.chunkSize); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("splitByChunks() = %v, want %v", got, tt.want)
			}
		})
	}
}
