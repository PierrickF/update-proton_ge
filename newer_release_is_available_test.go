package update_proton_ge_test

import (
	"github.com/google/go-cmp/cmp"
	"testing"
	"update_proton_ge"
)

func TestNewerReleaseIsAvailableDetectsDifferentStrings(t *testing.T) {
	t.Parallel()

	testData := [3]string{"GE-Proton8-30", "GE-Proton8-31", "GE-Proton8-32"}
	got := [3]bool{}

	for index, data := range testData {
		res := update_proton_ge.NewerReleaseIsAvailable("GE-Proton8-30", data)

		got[index] = res
	}

	want := [3]bool{false, true, true}

	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}
