package update_proton_ge_test

import (
	"os"
	"os/exec"
	"testing"
	"update_proton_ge"
)

func TestParseGithubDataReturnsCorrectVersionName(t *testing.T) {
	t.Parallel()

	data, err := os.ReadFile("testdata/gh-data.json")

	if err != nil {
		t.Fatal(err)
	}

	want := "GE-Proton8-32"
	got, err := update_proton_ge.ParseGithubData(data)

	if err != nil {
		t.Fatal(err)
	}

	if want != got {
		t.Errorf("Want %q, got %q", want, got)
	}

}

func TestGetLocal(t *testing.T) {
	t.Parallel()

	got, err := update_proton_ge.GetLocal()

	if err != nil {
		t.Errorf("Failed to run shell commands: %s", err)
	}

	if got == "" {
		t.Error("Want a version name, got an empty string")
	}
}

func TestCompareVersions(t *testing.T) {
	t.Parallel()

	testLatest := "meow-8.34"
	testLocal := "meow-8.34"

	want := true
	got := update_proton_ge.CompareVersions(testLatest, testLocal)

	if want != got {
		t.Errorf("Compare %q and %q. Wanted %t, got %t.",
			testLatest,
			testLocal,
			want,
			got)
	}
}

func TestMakeTempDir(t *testing.T) {
	t.Parallel()

	rm := exec.Command("rm", "-rf", "/tmp/proton-ge-custom/")
	_ = rm.Run()

	err := update_proton_ge.MakeTempDir()

	if err != nil {
		t.Errorf("Failed to run mkdir : %s", err)
	}
}
