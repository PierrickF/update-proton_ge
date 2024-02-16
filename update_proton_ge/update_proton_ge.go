package update_proton_ge

import "os/exec"

// GetLatest runs shell commands to fetch a URL and strip it down to a version name.
// The output of the shell commands is captured in a slice of bytes,
// which is then converted to a string.
func GetLatest() string {
	latest, _ := exec.Command("curl -s https://api.github.com/repos/GloriousEggroll/proton-ge-custom/releases/latest | grep browser_download_url | cut -d\" -f4 | grep .tar.gz | cut -d / -f8").Output()

	return string(latest)
}

// CompareVersions takes in two strings and compare them.
// If they are equal, the software is up to date.
func CompareVersions(latest, local string) bool {
	return latest == local
}
