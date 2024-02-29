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

	type testCase struct {
		target string
		want   string
	}

	testCases := []testCase{
		{target: "releaseTag",
			want: "GE-Proton8-32"},
		{target: "tarballUrl",
			want: "https://github.com/GloriousEggroll/proton-ge-custom/releases/download/GE-Proton8-32/GE-Proton8-32.tar.gz"},
		{target: "checksumUrl",
			want: "https://github.com/GloriousEggroll/proton-ge-custom/releases/download/GE-Proton8-32/GE-Proton8-32.sha512sum"},
	}

	for _, tc := range testCases {
		got, err := update_proton_ge.ParseReleaseData(data, tc.target)

		if err != nil {
			t.Fatal(err)
		}

		if tc.want != got {
			t.Errorf("Want %q, got %q", tc.want, got)
		}
	}
}
