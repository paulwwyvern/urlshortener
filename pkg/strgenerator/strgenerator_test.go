package strgenerator

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenerator(t *testing.T) {
	tests := []struct {
		name  string
		chars string
		seed  int64
		len   int
		want  []string
	}{
		{
			name:  "Test #1 OnlyDigits",
			chars: Digits,
			len:   10,
			seed:  42,
			want: []string{
				"5780357683",
				"9758232154",
				"2445412979",
				"2250984727",
			},
		}, {
			name:  "Test #2 OnlyLetters",
			chars: LowercaseLatin,
			len:   10,
			seed:  42,
			want: []string{
				"hrukpttuez",
				"ptneuvunhu",
				"ksqvgzadxl",
				"gghejkmvez",
			},
		}, {
			name:  "Test #3 Random",
			chars: "ab",
			len:   1,
			seed:  999,
			want: []string{
				"a",
				"a",
				"b",
			},
		}, {
			name:  "Test #4 Random",
			chars: LowercaseLatin + UppercaseLatin + Digits,
			len:   10,
			seed:  999,
			want: []string{
				"aMRZFw7y9i",
				"rKMCqJsQNa",
				"uTYJLCEoGw",
				"yb35z436xb",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			generator := NewGenerator(tt.chars, tt.len, tt.seed)
			for _, w := range tt.want {
				got := generator.Generate()
				assert.Equal(t, w, got)
			}
		})
	}
}
