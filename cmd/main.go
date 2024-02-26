package main

import (
	"os"
	"update_proton_ge"
)

func main() {
	installDir := "~/.steam/root/compatibilitytools.d/"

	releaseData, _ := update_proton_ge.DownloadReleaseData()
	latestRelease, _ := update_proton_ge.ParseReleaseData(releaseData)
	latestInstall, _ := update_proton_ge.FindLatestInstall(installDir)
	if !update_proton_ge.NewerReleaseIsAvailable(latestRelease, latestInstall) {
		println("Already up to date, nothing to do.")
		os.Exit(0)
	}
}
