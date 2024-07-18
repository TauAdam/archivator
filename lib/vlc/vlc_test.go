package vlc

import "testing"

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
