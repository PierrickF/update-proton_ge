package update_proton_ge

import (
	"log"
	"os/exec"
)

// GetLatest runs shell commands to fetch a URL and strip it down to a version name.
// The output of the shell commands is captured in a slice of bytes,
// which is then converted to a string.
func GetLatest() string {

	// Define shell commands
	curlUrl := "https://api.github.com/repos/GloriousEggroll/proton-ge-custom/releases/latest"
	curl := exec.Command("curl", "-s", curlUrl)
	grep := exec.Command("grep", "-E", "browser_download_url.*.tar.gz")
	cut1 := exec.Command("cut", "-d", "\"", "-f4")
	cut2 := exec.Command("cut", "-d", "/", "-f8")

	// Set up pipes between commands
	curlOutput, err := curl.StdoutPipe()
	if err != nil {
		log.Fatalf("Error piping curl to grep: %s", err)
	}
	grep.Stdin = curlOutput

	grepOutput, err := grep.StdoutPipe()
	if err != nil {
		log.Fatalf("Error piping grep to cut: %s", err)
	}
	cut1.Stdin = grepOutput

	cut1Output, err := cut1.StdoutPipe()
	if err != nil {
		log.Fatalf("Error piping cut to cut: %s", err)
	}
	cut2.Stdin = cut1Output

	// Start commands
	err = curl.Start()
	if err != nil {
		log.Fatalf("Error starting curl: %s", err)
	}

	err = grep.Start()
	if err != nil {
		log.Fatalf("Error starting grep: %s", err)
	}

	err = cut1.Start()
	if err != nil {
		log.Fatalf("Error starting cut: %s", err)
	}

	// Start, retreive output, and terminate command
	latest, err := cut2.Output()
	if err != nil {
		log.Fatalf("Error running cut: %s", err)
	}

	// Terminate commands
	err = curl.Wait()
	if err != nil {
		log.Fatalf("Error terminating curl: %s", err)
	}

	err = grep.Wait()
	if err != nil {
		log.Fatalf("Error terminating grep: %s", err)
	}

	err = cut1.Wait()
	if err != nil {
		log.Fatalf("Error terminating cut: %s", err)
	}
	return string(latest)
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
