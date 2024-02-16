package update_proton_ge

import (
	"encoding/json"
	"fmt"
	"os/exec"
)

func ParseGithubData(data []byte) (string, error) {
	var release struct {
		ReleaseTag string `json:"tag_name"`
	}
	err := json.Unmarshal(data, &release)
	if err != nil {
		return "", err
	}

	return release.ReleaseTag, nil
}

// GetLatest runs shell commands to fetch a URL

// and strip it down to a version name.
// The output of the shell commands is captured in a slice of bytes,
// which is then converted to a string.
func GetLatest() (string, error) {

	// Define shell commands
	curlUrl := "https://api.github.com/repos/GloriousEggroll/proton-ge-custom/releases/latest"
	curl := exec.Command("curl", "-s", curlUrl)

	curlOutput, err := curl.Output()
	if err != nil {
		return "", fmt.Errorf("Curl error: %w", err)
	}

	return ParseGithubData(curlOutput)

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

// MakeTempDir creates a temporary directory which will be used to download data.
func MakeTempDir() error {
	mkdir := exec.Command("mkdir", "/tmp/proton-ge-custom")
	return mkdir.Run()
}
