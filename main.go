package main

import (
	"fmt"
	"io/fs"
	"path/filepath"
)

type Walker struct {
	paths []string
}

func main() {
	w := Walker{}
	filepath.WalkDir("/home/aneesh", w.walk)
	fmt.Println(w.paths)
}

func (w *Walker) walk(s string, d fs.DirEntry, err error) error {
	if err != nil {
		return err
	}
	if !d.IsDir() {
		w.paths = append(w.paths, s)
	}
	return nil
}
