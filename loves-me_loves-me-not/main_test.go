package lovesmelovesmenot

import (
	"testing"
)

func TestLovesMe(t *testing.T) {
	var tests = []struct {
		p      int
		expect string
	}{
		{1, "LOVES ME"},
		{3, "Loves me, Loves me not, LOVES ME"},
		{6, "Loves me, Loves me not, Loves me, Loves me not, Loves me, LOVES ME NOT"},
	}

	for _, tt := range tests {
		t.Run(tt.expect, func(t *testing.T) {
			result := LovesMe(tt.p)
			if result != tt.expect {
				t.Errorf("got %q, want %q", result, tt.expect)
			}
		})
	}
}
