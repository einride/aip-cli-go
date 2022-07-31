package protoversion

import (
	"fmt"
	"strconv"
	"strings"
)

// Version represents a protobuf package version.
type Version struct {
	// Major version of the package.
	Major int
	// Stability of the package.
	Stability StabilityLevel
	// Minor version of the package.
	Minor int
}

// Parse a version string.
func Parse(s string) (Version, error) {
	if !strings.HasPrefix(s, "v") {
		return Version{}, fmt.Errorf("missing prefix 'v'")
	}
	withoutV := strings.TrimPrefix(s, "v")
	if strings.HasPrefix(withoutV, "0") {
		return Version{}, fmt.Errorf("version must not start with 0")
	}
	withoutMajor := strings.TrimLeft(withoutV, "0123456789")
	majorStr := withoutV[:len(withoutV)-len(withoutMajor)]
	major, err := strconv.Atoi(majorStr)
	if err != nil {
		return Version{}, fmt.Errorf("invalid major version: %w", err)
	}
	if len(withoutMajor) == 0 {
		return Version{Major: major}, nil
	}
	var withoutStabilityLevel string
	stabilityLevel := Stable
	switch {
	case strings.HasPrefix(withoutMajor, "alpha"):
		stabilityLevel = Alpha
		withoutStabilityLevel = strings.TrimPrefix(withoutMajor, "alpha")
	case strings.HasPrefix(withoutMajor, "beta"):
		stabilityLevel = Beta
		withoutStabilityLevel = strings.TrimPrefix(withoutMajor, "beta")
	}
	if len(withoutStabilityLevel) == 0 {
		return Version{Major: major, Stability: stabilityLevel}, nil
	}
	minor, err := strconv.Atoi(withoutStabilityLevel)
	if err != nil {
		return Version{}, fmt.Errorf("invalid minor version: %w", err)
	}
	return Version{
		Major:     major,
		Stability: stabilityLevel,
		Minor:     minor,
	}, nil
}
