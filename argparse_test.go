package argparse_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/zikes/argparse"
)

func TestParse(t *testing.T) {
	tests := []struct {
		input  string
		output []string
	}{
		{
			// Test quotes
			`this "is a" test`,
			[]string{"this", "is a", "test"},
		},
		{
			// Test escaping
			`this \" "is a" te\ st`,
			[]string{"this", "\"", "is a", "te st"},
		},
	}
	for _, test := range tests {
		got := argparse.Parse(test.input)
		if diff := cmp.Diff(test.output, got); diff != "" {
			t.Errorf("Parse mismatch (-want +got):\n%s", diff)
		}
	}
}
