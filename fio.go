package main

import (
	"bufio"
	"io"
	"io/fs"
	"io/ioutil"
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

func read_file_str(fpath string) (string, error) {
	data, err := ioutil.ReadFile(fpath)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func read_file_arr(fpath string) ([]string, error) {
	file, err := os.Open(fpath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, nil
}

func read_file_runearr_arr(fpath string) ([][]rune, error) {
	arr := [][]rune{}

	fp, err := os.OpenFile(fpath, os.O_RDONLY, 0660)
	if err != nil {
		return nil, err
	}

	reader := bufio.NewReader(fp)
	for {
		line, err := reader.ReadBytes('\n')
		if err != nil || err == io.EOF {
			break
		}
		arr = append(arr, []rune(string(line)))
	}

	return arr, nil
}

func path_describe(path string) (bool, bool) {
	f_stat, err := os.Stat(path)
	if err != nil {

		if os.IsNotExist(err) {
			return false, false
		}
		return false, false
	}
	return true, f_stat.IsDir()
}
