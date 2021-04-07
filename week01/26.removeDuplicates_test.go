package week01

import "testing"

// run: go test -run Test_removeDuplicates
func Test_removeDuplicates(t *testing.T) {
	cases := []struct {
		name     string
		input    []int
		expected int
	}{
		{
			name:     "case1",
			input:    []int{1, 1, 2},
			expected: 2,
		},
		{
			name:     "case2",
			input:    []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			expected: 5,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if got := removeDuplicates(c.input); got != c.expected {
				t.Errorf("removeDuplicates(%v)=%v, expected=%v", c.input, got, c.expected)
			}
		})
	}
}
