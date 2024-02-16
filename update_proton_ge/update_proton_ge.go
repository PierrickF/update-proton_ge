package update_proton_ge

// CompareVersions takes in two strings and compare them.
// If they are equal, the software is up to date.
func CompareVersions(latest, local string) bool {
	return latest == local
}
