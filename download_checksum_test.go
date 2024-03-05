package update_proton_ge_test

import (
	"os"
	"testing"
	"update_proton_ge"
)

func TestDownloadChecksumDownloadsCorrectRelease(t *testing.T) {
	t.Parallel()

	// create the target directory if it doesn't exist already
	_ = update_proton_ge.MakeTempDirectory()

	testRelease := "GE-Proton8-32"
	testUrl := "https://github.com/GloriousEggroll/proton-ge-custom/releases/download/GE-Proton8-32/GE-Proton8-32.sha512sum"

	if err := update_proton_ge.DownloadChecksum(testRelease, testUrl); err != nil {
		t.Fatal(err)
	}

	if _, err := os.Stat("/tmp/update-proton-ge/GE-Proton8-32.sha512sum"); err != nil {
		t.Error(err)
	}

	// cleanup to avoid false positive next time the test runs
	if err := os.RemoveAll("/tmp/update-proton-ge/"); err != nil {
		t.Fatal(err)
	}
}
