package update_proton_ge_test

import (
	"os"
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
