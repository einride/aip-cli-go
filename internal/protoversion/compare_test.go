package protoversion

import (
	"testing"

	"gotest.tools/v3/assert"
)

func TestCompare(t *testing.T) {
	for _, tt := range []struct {
		name     string
		a, b     Version
		expected int
	}{
		{
			name:     "identical",
			a:        Version{Major: 2, Stability: Beta, Minor: 2},
			b:        Version{Major: 2, Stability: Beta, Minor: 2},
			expected: 0,
		},

		{
			name:     "lower minor",
			a:        Version{Major: 2, Stability: Beta, Minor: 1},
			b:        Version{Major: 2, Stability: Beta, Minor: 2},
			expected: -1,
		},

		{
			name:     "higher minor",
			a:        Version{Major: 2, Stability: Beta, Minor: 2},
			b:        Version{Major: 2, Stability: Beta, Minor: 1},
			expected: 1,
		},

		{
			name:     "lower stability",
			a:        Version{Major: 2, Stability: Alpha, Minor: 2},
			b:        Version{Major: 2, Stability: Beta, Minor: 2},
			expected: -1,
		},

		{
			name:     "higher stability",
			a:        Version{Major: 2, Stability: Beta, Minor: 2},
			b:        Version{Major: 2, Stability: Alpha, Minor: 2},
			expected: 1,
		},

		{
			name:     "lower minor higher stability",
			a:        Version{Major: 2, Stability: Beta, Minor: 1},
			b:        Version{Major: 2, Stability: Alpha, Minor: 2},
			expected: 1,
		},

		{
			name:     "stable vs alpha",
			a:        Version{Major: 1},
			b:        Version{Major: 1, Stability: Beta, Minor: 1},
			expected: 1,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			actual := Compare(tt.a, tt.b)
			assert.Equal(t, tt.expected, actual)
		})
	}
}
