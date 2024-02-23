package update_proton_ge

import (
	"os"
)

// FindLatestInstall takes in the install directory for Proton GE on the local
// computer, and returns the name of the last file found in said directory.
func FindLatestInstall(installDir string) (name string, err error) {
	files, err := os.ReadDir(installDir)

	if err != nil {
		return "", err
	}

	return files[len(files)-1].Name(), nil
}
