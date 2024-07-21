package shannon_fano

import (
	"reflect"
	"testing"
)

func Test_findBestPosition(t *testing.T) {
	tests := []struct {
		name string
		args []code
		want int
	}{
		{
			name: "1 element",
			args: []code{
				{Quantity: 2},
			},
			want: 0,
		},
		{
			name: "1 element",
			args: []code{
				{Quantity: 2},
				{Quantity: 2},
			},
			want: 1,
		},
		{
			name: "2 elements",
			args: []code{
				{Quantity: 2},
				{Quantity: 2},
			},
			want: 1,
		},
		{
			name: "3 elements",
			args: []code{
				{Quantity: 2},
				{Quantity: 1},
				{Quantity: 1},
			},
			want: 1,
		},
		{
			name: "4 elements",
			args: []code{
				{Quantity: 1},
				{Quantity: 3},
				{Quantity: 2},
				{Quantity: 1},
			},
			want: 2,
		},
		{
			name: "5 elements",
			args: []code{
				{Quantity: 2},
				{Quantity: 2},
				{Quantity: 1},
				{Quantity: 1},
				{Quantity: 1},
			},
			want: 2,
		},
		{
			name: "implicit position",
			args: []code{
				{Quantity: 1},
				{Quantity: 1},
				{Quantity: 1},
			},
			want: 1,
		},
		{
			name: "implicit position",
			args: []code{
				{Quantity: 2},
				{Quantity: 2},
				{Quantity: 1},
				{Quantity: 1},
			},
			want: 1,
		},
		{
			name: "implicit position",
			args: []code{
				{Quantity: 2},
				{Quantity: 2},
				{Quantity: 3},
				{Quantity: 1},
				{Quantity: 1},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findBestPosition(tt.args); got != tt.want {
				t.Errorf("findBestPosition() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_assignCodes(t *testing.T) {
	type args struct {
		codes []code
	}
	tests := []struct {
		name string
		args args
		want []code
	}{
		{name: "2 elements",
			args: args{
				codes: []code{
					{Quantity: 2},
					{Quantity: 2},
				},
			}, want: []code{
				{Quantity: 2, Bits: 0, Size: 1},
				{Quantity: 2, Bits: 1, Size: 1},
			}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if assignCodes(tt.args.codes); !reflect.DeepEqual(tt.args.codes, tt.want) {
				t.Errorf("assignCodes() = %v, want %v", tt.args.codes, tt.want)
			}
		})
	}
}
