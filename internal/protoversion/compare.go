package protoversion

// Compare returns an integer comparing two versions.
// The result will be 0 if a == b, -1 if a < b, and +1 if a > b.
func Compare(a, b Version) int {
	switch {
	case a.Major < b.Major:
		return -1
	case a.Major > b.Major:
		return 1
	case a.Stability > b.Stability:
		return -1
	case a.Stability < b.Stability:
		return 1
	case a.Minor < b.Minor:
		return -1
	case a.Minor > b.Minor:
		return 1
	default:
		return 0
	}
}
