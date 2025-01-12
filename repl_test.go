package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "Martiwka, is CrazY!!!",
			expected: []string{"martiwka,", "is", "crazy!!!"},
		},
	}

	for _, c := range cases {
		clean := cleanInput(c.input)

		if len(clean) != len(c.expected) {
			t.Errorf("wrong count: %d", len(clean))
		}

		for i, actual := range clean {
			expected := c.expected[i]

			if expected != actual {
				t.Errorf("Expected: %s, Actual: %s", expected, actual)
			}
		}
	}
}
