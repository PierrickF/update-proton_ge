package update_proton_ge_test

import (
	"os"
	"testing"
	"update_proton_ge"
)

func TestMakeTempDirectoryCreatesCorrectDirectory(t *testing.T) {
	t.Parallel()

	err := update_proton_ge.MakeTempDirectory()

	if err != nil {
		t.Fatal(err)
	}

	if _, err := os.Stat("/tmp/update-proton-ge/"); err != nil {
		t.Error(err)
	}

	// cleanup to avoid false positive next time the test runs
	_ = os.Remove("/tmp/update-proton-ge/")
}
