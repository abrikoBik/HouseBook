package util

import (
	"fmt"
	"testing"
)

func TestExtractNumber(t *testing.T) {
	t.Parallel()
	tests := []struct {
		str    string
		result int
	}{
		{"1", 1},
		{"1 1", 11},
		{"0 1 2 test 3", 123},
		{"123 0 321", 1230321},
		{"string", 0},
	}
	for _, c := range tests {
		name := fmt.Sprintf("%s,%d", c.str, c.result)
		t.Run(name, func(t *testing.T) {
			result, err := ExtractNumber(c.str)
			if err != nil {
				t.Fatal(err)
			}
			if result != c.result {
				t.Errorf("ExtractNumber returned %d; wanted %d", result, c.result)
			}
		})
	}
}
