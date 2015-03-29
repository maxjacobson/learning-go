package collatz

import (
	"testing"

	"github.com/maxjacobson/learning-go/collatz"
)

func TestFrom(t *testing.T) {
	chain := collatz.From(13)
	if chain.Length != 10 {
		t.Errorf("From(13) == %q, want 10", chain.Length)
	}
}
