package vlc

import (
	"reflect"
	"testing"
)

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

func TestBinaryChunks_ToHex(t *testing.T) {
	tests := []struct {
		name string
		c    BinaryChunks
		want HexChunks
	}{
		{
			name: "basic case",
			c:    BinaryChunks{"00100000", "11010010"},
			want: HexChunks{"20", "D2"},
		},
		{
			name: "case with 4 chunks",
			c:    BinaryChunks{"0010", "0000", "1101", "1001"},
			want: HexChunks{"02", "00", "0D", "09"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.ToHex(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToHex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewHexChunks(t *testing.T) {
	tests := []struct {
		name string
		str  string
		want HexChunks
	}{
		{
			name: "basic case",
			str:  "20 D2",
			want: HexChunks{"20", "D2"},
		},
		{
			name: "case with spaces",
			str:  "20 D2 00 0D 09",
			want: HexChunks{"20", "D2", "00", "0D", "09"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewHexChunks(tt.str); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewHexChunks() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHexChunk_ToBinary(t *testing.T) {
	tests := []struct {
		name string
		c    HexChunk
		want BinaryChunk
	}{
		{
			name: "basic case",
			c:    HexChunk("20"),
			want: BinaryChunk("00100000"),
		},
		{
			name: "case with 0",
			c:    HexChunk("00"),
			want: BinaryChunk("00000000"),
		},
		{
			name: "case with 9",
			c:    HexChunk("09"),
			want: BinaryChunk("00001001"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.ToBinary(); got != tt.want {
				t.Errorf("ToBinary() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHexChunks_ToBinary(t *testing.T) {
	tests := []struct {
		name string
		c    HexChunks
		want BinaryChunks
	}{
		{
			name: "basic case",
			c:    HexChunks{"20", "D2"},
			want: BinaryChunks{"00100000", "11010010"},
		},
		{name: "case with 0",
			c:    HexChunks{"00", "00"},
			want: BinaryChunks{"00000000", "00000000"},
		},
		{
			name: "case with 9",
			c:    HexChunks{"09", "09"},
			want: BinaryChunks{"00001001", "00001001"},
		}, {
			name: "case with letters",
			c:    HexChunks{"A1", "B2"},
			want: BinaryChunks{"10100001", "10110010"},
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.ToBinary(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToBinary() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBinaryChunks_Join(t *testing.T) {
	tests := []struct {
		name string
		c    BinaryChunks
		want string
	}{
		{
			name: "basic case",
			c:    BinaryChunks{"00100000", "11010010"},
			want: "0010000011010010",
		},
		{name: "case with 0",
			c:    BinaryChunks{"00000000", "00000000"},
			want: "0000000000000000",
		},
		{name: "case with 9",
			c:    BinaryChunks{"00001001", "00001001"},
			want: "0000100100001001",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.Join(); got != tt.want {
				t.Errorf("Join() = %v, want %v", got, tt.want)
			}
		})
	}
}
