package update_proton_ge

import "os"

func MakeTempDirectory() error {
	err := os.Mkdir("/tmp/update-proton-ge/", 0660)

	if err != nil {
		return err
	}

	return nil
}
