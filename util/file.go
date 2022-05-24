package util

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
)

func GetFiles(dir string) ([]string, error) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, fmt.Errorf("cannot open directory: %w", err)
	}

	var paths []string
	for _, file := range files {
		if !file.IsDir() && filepath.Ext(file.Name()) == ".sql" {
			paths = append(paths, file.Name())
		}
	}
	return paths, nil
}

// baseに含まれてtargetに含まれない要素を返す
func Diff(base []string, target []string) []string {
	mb := make(map[string]bool, len(target))
	for _, x := range target {
		mb[x] = true
	}
	var diff []string
	for _, x := range base {
		if _, found := mb[x]; !found {
			diff = append(diff, x)
		}
	}
	return diff
}

func GetFileNameWithoutExt(path string) string {
	return filepath.Base(path[:len(path)-len(filepath.Ext(path))])
}

func ReadFileAsString(filename string) (string, error) {
	data, err := ioutil.ReadFile(filename)
	return string(data), err
}
