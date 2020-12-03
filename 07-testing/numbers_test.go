package numbers

import "testing"

func Test_Sum(t *testing.T) {
	// creating test cases is useful to loop through the possible scenarios
	// you want to test without having to do the setup all over again.
	testCases := []struct {
		name   string
		input  []int
		output int
	}{
		{
			name:   "testing with 1, 3, 5",
			input:  []int{1, 3, 5},
			output: 9,
		},
		{
			name:   "testing with 2, 4, 6",
			input:  []int{2, 4, 6},
			output: 12,
		},
		{
			name:   "testing with 1",
			input:  []int{1},
			output: 1,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// we use the elipsis here to pass individual array elements
			result := Sum(tc.input...)
			if result != tc.output {
				t.Errorf("expected %d but got %d", tc.output, result)
			}
		})
	}
}
