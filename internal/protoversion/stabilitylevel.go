package protoversion

// StabilityLevel represents an AIP stability level.
// See: https://google.aip.dev/181
type StabilityLevel int

const (
	// A Stable component must be fully-supported over the lifetime of the major API version.
	Stable StabilityLevel = iota

	// A Beta component must be considered complete and ready to be declared stable, subject to public testing.
	Beta

	// An Alpha component undergoes rapid iteration with a known set of users who must be tolerant of change.
	Alpha
)
