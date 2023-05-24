package tools

import "path/filepath"

func GetDeployPath(myPath string) string {
	abs, err := filepath.Abs("")
	if err != nil {
		panic(any(err))
	}
	return filepath.Join(abs, myPath)
}
