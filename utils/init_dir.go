package utils

import (
	"fmt"
	"os"
)

func InitDirectory(s string) error {
	dir, err := os.Stat(s)
	if os.IsNotExist(err) {
		return os.MkdirAll(s, 0700)
	}
	if err != nil {
		return err
	}
	if !dir.IsDir() {
		return fmt.Errorf("%s already exists but is not a directory", s)
	}
	return nil
}
