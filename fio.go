package main

import (
	"bufio"
	"io/ioutil"
	"log"
	"os"
)

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
