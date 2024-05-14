package utils

import "path/filepath"

func GetContentDir(name ...string) string {
	dir := append([]string{".", "web", "contents"}, name...)
	return filepath.Join(dir...)
}

func GetContentPath(name ...string) string {
	dir := append([]string{"/", "static", "contents"}, name...)
	return filepath.Join(dir...)
}

func GetStaticDir(name ...string) string {
	dir := append([]string{".", "web", "static"}, name...)
	return filepath.Join(dir...)
}

func GetStaticPath(name ...string) string {
	dir := append([]string{"/", "static"}, name...)
	return filepath.Join(dir...)
}
