package update_proton_ge

import (
	"fmt"
	"os/exec"
)

func GetLatest() (string, error) {

	curlUrl := "https://api.github.com/repos/GloriousEggroll/proton-ge-custom/releases/latest"
	curl := exec.Command("curl", "-s", curlUrl)

	curlOutput, err := curl.Output()
	if err != nil {
		return "", fmt.Errorf("curl error: %w", err)
	}

	return ParseGithubData(curlOutput)

}
