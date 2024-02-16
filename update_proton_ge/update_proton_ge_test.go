package update_proton_ge_test

import (
	"main/update_proton_ge"
	"os/exec"
	"testing"
)

func TestGetLatest(t *testing.T) {
	t.Parallel()

	got := update_proton_ge.GetLatest()

	if got == "" {
		t.Error("Want a version name, got an empty string")
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
