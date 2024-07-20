package vlc

import (
	"reflect"
	"testing"
)

func Test_encodingTable_DecodingTree(t *testing.T) {
	tests := []struct {
		name string
		t    encodingTable
		want DecodingTree
	}{
		{name: "basic case",
			t: encodingTable{
				'a': "11",
				'b': "1001",
				'c': "0101",
			},
			want: DecodingTree{
				Left: &DecodingTree{
					Right: &DecodingTree{
						Left: &DecodingTree{
							Right: &DecodingTree{
								Value: "c",
							},
						},
					},
				},
				Right: &DecodingTree{
					Left: &DecodingTree{
						Left: &DecodingTree{
							Right: &DecodingTree{Value: "b"},
						},
					},
					Right: &DecodingTree{
						Value: "a",
					},
				},
			},
		},
		{
			name: "empty encoding table",
			t:    encodingTable{},
			want: DecodingTree{},
		},
		{
			name: "single character encoding",
			t: encodingTable{
				'd': "0",
			},
			want: DecodingTree{
				Left: &DecodingTree{
					Value: "d",
				},
			},
		},
		{
			name: "encoding with common prefix",
			t: encodingTable{
				'e': "10",
				'f': "11",
			},
			want: DecodingTree{
				Right: &DecodingTree{
					Left: &DecodingTree{
						Value: "e",
					},
					Right: &DecodingTree{
						Value: "f",
					},
				},
			},
		},
		{
			name: "complex encoding structure",
			t: encodingTable{
				'g': "000",
				'h': "001",
				'i': "010",
				'j': "011",
			},
			want: DecodingTree{
				Left: &DecodingTree{
					Left: &DecodingTree{
						Left: &DecodingTree{
							Value: "g",
						},
						Right: &DecodingTree{
							Value: "h",
						},
					},
					Right: &DecodingTree{
						Left: &DecodingTree{
							Value: "i",
						},
						Right: &DecodingTree{
							Value: "j",
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.t.DecodingTree(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DecodingTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
