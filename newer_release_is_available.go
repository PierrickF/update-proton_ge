package update_proton_ge

// NewerReleaseIsAvailable takes in the name of the latest Proton GE release,
// and the name of the latest version of Proton GE installed.
// It compares them, and returns whether or not their are equal.
// If they are different, a new version is available.
func NewerReleaseIsAvailable(release, install string) bool {
	return release != install
}
