package update_proton_ge_test

import (
	"encoding/json"
	"testing"
	"update_proton_ge"
)

func TestDownloadReleaseDataReturnsValidJson(t *testing.T) {
	t.Parallel()

	data, err := update_proton_ge.DownloadReleaseData()

	if err != nil {
		t.Fatal(err)
	}

	want := true
	got := json.Valid(data)

	if want != got {
		t.Error("data is not valid JSON")
	}
}
