package b3propagator

// Version is the current release version of the B3 propagator.
func Version() string {
	return "1.11.1"
	// This string is updated by the pre_release.sh script during release
}

// SemVersion is the semantic version to be supplied to tracer/meter creation.
func SemVersion() string {
	return "semver:" + Version()
}
