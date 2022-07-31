package protoversion

import (
	"testing"

	"gotest.tools/v3/assert"
)

func TestParse(t *testing.T) {
	for _, tt := range []struct {
		version       string
		expected      Version
		errorContains string
	}{
		{
			version:  "v1",
			expected: Version{Major: 1},
		},

		{
			version:  "v2",
			expected: Version{Major: 2, Stability: Stable},
		},

		{
			version:  "v1beta1",
			expected: Version{Major: 1, Stability: Beta, Minor: 1},
		},

		{
			version:  "v1beta2",
			expected: Version{Major: 1, Stability: Beta, Minor: 2},
		},

		{
			version:  "v2alpha3",
			expected: Version{Major: 2, Stability: Alpha, Minor: 3},
		},

		{
			version:       "foo",
			errorContains: "missing prefix 'v'",
		},
	} {
		t.Run(tt.version, func(t *testing.T) {
			actual, err := Parse(tt.version)
			if tt.errorContains != "" {
				assert.ErrorContains(t, err, tt.errorContains)
			} else {
				assert.Equal(t, tt.expected, actual)
			}
		})
	}
}
