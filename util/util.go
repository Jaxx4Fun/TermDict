package util

import (
	"fmt"
	"os"
)

func MkdirIfNotExist(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.MkdirAll(path, 0777)
		if err != nil {
			return fmt.Errorf("failed to mkdir %s, %v", path, err)
		}
	}
	return nil
}
