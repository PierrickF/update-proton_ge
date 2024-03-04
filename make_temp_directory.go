package update_proton_ge

import "os"

func MakeTempDirectory() error {
	if err := os.Mkdir("/tmp/update-proton-ge/", 0777); err != nil {
		return err
	}

	return nil
}
