package update_proton_ge_test

import (
	"os"
	"testing"
	"update_proton_ge"
)

func TestParseReleaseDataReturnsCorrectVersionName(t *testing.T) {
	t.Parallel()

	data, err := os.ReadFile("testdata/release-data.json")

	if err != nil {
		t.Fatal(err)
	}

	want := "GE-Proton8-32"
	got, err := update_proton_ge.ParseReleaseData(data)

	if err != nil {
		t.Fatal(err)
	}

	if want != got {
		t.Errorf("Want %q, got %q", want, got)
	}
}
