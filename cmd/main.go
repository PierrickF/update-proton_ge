package main

import (
	"os"
	"update_proton_ge"
)

func main() {
	installDir := "~/.steam/root/compatibilitytools.d/"
	tmpDir := "/tmp/update-proton-eg/"

	releaseData, _ := update_proton_ge.DownloadReleaseData()
	latestRelease, _ := update_proton_ge.ParseReleaseData(releaseData, "releaseTag")
	latestInstall, _ := update_proton_ge.FindLatestInstall(installDir)
	println("Latest release: %q", latestRelease)
	println("Latest install: %q", latestInstall)
	if !update_proton_ge.NewerReleaseIsAvailable(latestRelease, latestInstall) {
		println("Already up to date, nothing to do.")
		os.Exit(0)
	}
	println("New release available, installing...")
	update_proton_ge.MakeTempDirectory()
	tarballUrl, _ := update_proton_ge.ParseReleaseData(releaseData, "tarballUrl")
	update_proton_ge.DownloadTarball(latestRelease, tarballUrl)
	checksumUrl, _ := update_proton_ge.ParseReleaseData(releaseData, "checksumUrl")
	update_proton_ge.DownloadChecksum(latestRelease, checksumUrl)
	update_proton_ge.CheckTarball(tmpDir+latestRelease+".tar.gz", tmpDir+latestRelease+".sha512sum")
	update_proton_ge.MakeSteamDirectory()
	update_proton_ge.ExtractTarball()
	update_proton_ge.RemoveTempDirectory()
	println("All done.")
}
