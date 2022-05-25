package main

import (
	"bufio"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func dir_ls_rec(dir_path string) []string {
	w := Walker{}
	filepath.WalkDir(dir_path, w.walk)
	return w.paths
}

type Walker struct {
	paths []string
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

func read_file_str(fpath string) string {
	data, err := ioutil.ReadFile(fpath)
	if err != nil {
		log.Panicf("Could not read data from %v", fpath)
	}

	return string(data)
}

func read_file_arr(fpath string) []string {
	file, err := os.Open(fpath)
	if err != nil {
		log.Panicf("Could not read data from %v", fpath)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}
