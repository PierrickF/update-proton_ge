package update_proton_ge

import (
	"os/exec"
)

// GetLatest runs shell commands to fetch a URL and strip it down to a version name.
// The output of the shell commands is captured in a slice of bytes,
// which is then converted to a string.
func GetLatest() (string, error) {
	curlUrl := "https://api.github.com/repos/GloriousEggroll/proton-ge-custom/releases/latest"
	curl := exec.Command("curl", "-s", curlUrl)
	grep := exec.Command("grep", "-E", "browser_download_url.*.tar.gz")
	cut1 := exec.Command("cut", "-d", "\"", "-f4")
	cut2 := exec.Command("cut", "-d", "/", "-f8")

	grep.Stdin, _ = curl.StdoutPipe()
	cut1.Stdin, _ = grep.StdoutPipe()
	cut2.Stdin, _ = cut1.StdoutPipe()

	_ = curl.Start()
	_ = grep.Start()
	_ = cut1.Start()

	latest, err := cut2.Output()

	_ = curl.Wait()
	_ = grep.Wait()
	_ = cut1.Wait()

	return string(latest), err
}

// GetLocal runs shell commands to find a file name.
// The output of the shell commands is captured in a slice of bytes,
// which is then converted to a string.
func GetLocal() (string, error) {
	//TODO: make the test pass with the relative path ~
	ls := exec.Command("ls", "/home/pierrick/.steam/root/compatibilitytools.d/")
	sort := exec.Command("sort", "-r")
	head := exec.Command("head", "-1")

	sort.Stdin, _ = ls.StdoutPipe()
	head.Stdin, _ = sort.StdoutPipe()

	_ = ls.Start()
	_ = sort.Start()

	local, err := head.Output()

	_ = ls.Wait()
	_ = sort.Wait()

	return string(local), err
}

// CompareVersions takes in two strings and compare them.
// If they are equal, the software is up to date.
func CompareVersions(latest, local string) bool {
	return latest == local
}
