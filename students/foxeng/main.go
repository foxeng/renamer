package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
)

func rename(old string) (string, bool) {
	re := regexp.MustCompile(`(.*)([0-9])\.txt`)
	newName := re.ReplaceAllString(old, "${1}.txt.${2}")
	return newName, newName != old
}

func main() {
	targets := os.Args[1:]

	renamer := func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		dir, file := filepath.Split(path)
		newName, match := rename(file)
		if !match {
			return nil
		}
		fmt.Println("Renaming", path, "to", dir+newName)
		return os.Rename(path, dir+newName)
	}

	for _, target := range targets {
		filepath.Walk(target, renamer)
	}
}
