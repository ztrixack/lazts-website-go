package utils

import "path/filepath"

func GetContentDir(name ...string) string {
	dir := append([]string{".", "contents"}, name...)
	return filepath.Join(dir...)
}
