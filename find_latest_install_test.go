package update_proton_ge_test

import (
	"testing"
	"update_proton_ge"
)

func TestFindLatestInstallReturnsCorrectVersionName(t *testing.T) {
	t.Parallel()

	installDir := "testdata/local-installs/"

	want := "GE-Proton8-32"
	got, err := update_proton_ge.FindLatestInstall(installDir)

	if err != nil {
		t.Fatal(err)
	}

	if want != got {
		t.Errorf("Want %q, got %q", want, got)
	}
}
