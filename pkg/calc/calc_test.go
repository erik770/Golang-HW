package calc

import (
	"github.com/stretchr/testify/require"
	"log"
	"testing"
)

func TestCalculate(t *testing.T) {
	type calculateOutput struct {
		resValue string
		err      error
	}
	cases := map[string]struct {
		in       string
		expected calculateOutput
	}{
		"simple sum": {
			in: "5+5",
			expected: calculateOutput{
				"10",
				nil,
			},
		},
		"simple sub": {
			in: "50-30",
			expected: calculateOutput{
				"20",
				nil,
			},
		},
		"simple mult": {
			in: "33*3",
			expected: calculateOutput{
				"99",
				nil,
			},
		},
		"simple div": {
			in: "770/10",
			expected: calculateOutput{
				"77",
				nil,
			},
		},
		"hard expression": {
			in: "596*3+65-(48+589)/7*58/2-(8*(6+50))",
			expected: calculateOutput{
				"-1234",
				nil,
			},
		},
	}

	for name, tc := range cases {
		tc := tc
		t.Run(name, func(t *testing.T) {
			var actual calculateOutput
			actual.resValue, actual.err = Calculate(tc.in)

			require.Equal(t, tc.expected, actual)
		})
		log.Println("SUCCESS")
	}
}