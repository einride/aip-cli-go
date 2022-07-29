package protoshell

import (
	"testing"

	"gotest.tools/v3/assert"
)

func TestCompleteResourceName(t *testing.T) {
	for _, tt := range []struct {
		name       string
		toComplete string
		pattern    string
		completion string
		ok         bool
	}{
		{
			name:       "mismatch",
			toComplete: "user",
			pattern:    "shippers/{shipper}/shipments/{shipment}",
			completion: "",
			ok:         false,
		},

		{
			name:       "empty input",
			toComplete: "",
			pattern:    "shippers/{shipper}/shipments/{shipment}",
			completion: "shippers/",
			ok:         true,
		},

		{
			name:       "partial segment with singleton pattern",
			toComplete: "ship",
			pattern:    "shippers",
			completion: "shippers",
			ok:         true,
		},

		{
			name:       "empty input with singleton pattern",
			toComplete: "",
			pattern:    "shippers",
			completion: "shippers",
			ok:         true,
		},

		{
			name:       "empty input with variable pattern",
			toComplete: "",
			pattern:    "{shipper}",
			completion: "",
			ok:         false,
		},

		{
			name:       "resource ID input with variable pattern",
			toComplete: "foo",
			pattern:    "{shipper}",
			completion: "foo",
			ok:         true,
		},

		{
			name:       "partial segment",
			toComplete: "shi",
			pattern:    "shippers/{shipper}/shipments/{shipment}",
			completion: "shippers/",
			ok:         true,
		},

		{
			name:       "full segment",
			toComplete: "shippers",
			pattern:    "shippers/{shipper}/shipments/{shipment}",
			completion: "shippers/",
			ok:         true,
		},

		{
			name:       "full segment with slash",
			toComplete: "shippers/",
			pattern:    "shippers/{shipper}/shipments/{shipment}",
			completion: "shippers/",
			ok:         true,
		},

		{
			name:       "partial resource ID",
			toComplete: "shippers/test",
			pattern:    "shippers/{shipper}/shipments/{shipment}",
			completion: "shippers/test",
			ok:         true,
		},

		{
			name:       "full resource ID",
			toComplete: "shippers/test/",
			pattern:    "shippers/{shipper}/shipments/{shipment}",
			completion: "shippers/test/shipments/",
			ok:         true,
		},

		{
			name:       "partial second collection",
			toComplete: "shippers/foo/ship",
			pattern:    "shippers/{shipper}/shipments/{shipment}",
			completion: "shippers/foo/shipments/",
			ok:         true,
		},

		{
			name:       "second collection mismatch",
			toComplete: "shippers/foo/sit",
			pattern:    "shippers/{shipper}/shipments/{shipment}",
			completion: "",
			ok:         false,
		},

		{
			name:       "partial second resource ID",
			toComplete: "shippers/foo/shipments/bar",
			pattern:    "shippers/{shipper}/shipments/{shipment}",
			completion: "shippers/foo/shipments/bar",
			ok:         true,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			actual, ok := CompleteResourceName(tt.toComplete, tt.pattern)
			assert.Equal(t, tt.ok, ok)
			assert.Equal(t, tt.completion, actual)
		})
	}
}
