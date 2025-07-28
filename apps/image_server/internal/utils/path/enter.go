package path

import "os"

func GetRootPath() (path string) {
	path, err := os.Getwd()
	if err != nil {
		return ""
	}
	return path
}
