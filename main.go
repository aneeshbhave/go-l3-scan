package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	var dict_path, match_path string
	flag.StringVar(&dict_path, "d", "", "Path to dictionary file")
	flag.StringVar(&match_path, "p", "", "Path to match file or dir")

	flag.Parse()

	_, err := os.Stat(dict_path)
	if err != nil {
		if os.IsNotExist(err) {
			print_error(fmt.Sprintf("%v does not exist\n", dict_path))
		} else {
			print_error(fmt.Sprintf("There was an unknown error with file/dir %v", dict_path))
		}
	}

	mat := init_matcher(1000, true, true)
	mat.f_add_pattern(dict_path, " ")

	match_info, err := os.Stat(match_path)
	if err != nil {
		if os.IsNotExist(err) {
			print_error(fmt.Sprintf("%v does not exist\n", match_path))
		} else {
			print_error(fmt.Sprintf("There was an unknown error with file/dir %v", match_path))
		}
	}
	if match_info.IsDir() {
		mat.dir_match_pattern(match_path, callback_func)
	} else {
		mat.f_match_pattern(match_path, callback_func)
	}

}

func print_usage() {
	fmt.Println("Usage: ./l3-scan ...")
	os.Exit(-1)
}
func print_error(message string) {
	fmt.Printf("%v", message)
	os.Exit(-2)
}

func remove_duplicates(arr []string) []string {
	keys := make(map[string]bool)
	list := []string{}

	for _, e := range arr {
		if _, value := keys[e]; !value {
			keys[e] = true
			list = append(list, e)
		}
	}

	return list
}

func callback_func(arr ...any) {
	fmt.Printf("%v %d %d %v %v\n", arr[0], arr[1], arr[2], arr[3], arr[4])
}
