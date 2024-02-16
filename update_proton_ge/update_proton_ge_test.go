package update_proton_ge_test

import (
	"main/update_proton_ge"
	"testing"
)

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
