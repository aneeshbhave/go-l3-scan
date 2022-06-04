package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	var dict_path, match_path, output_file, pad_str string
	var training, keep_special, is_raw bool
	var init_size int

	flag.StringVar(&dict_path, "d", "", "Path to dictionary file")
	flag.StringVar(&match_path, "p", "", "Path to match file or dir")
	flag.StringVar(&output_file, "o", "", "Path to output to")
	flag.StringVar(&pad_str, "pad", "", "Padding string")

	flag.BoolVar(&training, "t", false, "Training mode instead of matching mode (default is match mode)")
	flag.BoolVar(&keep_special, "s", false, "Keep special characters instead of replacing them with pattern")
	flag.BoolVar(&is_raw, "r", false, "Is the truth file raw pattern data or actual data")

	flag.IntVar(&init_size, "i", 250, "Guesstimate of how many patterns may appear in a file (for optimization)")

	flag.Parse()

	//Assign callback function based on what -o parameter is
	callback_func := callback_stdout
	if output_file != "" {
		callback_func = callback_file
	}

	//Initialize matcher struct
	mat := init_matcher(init_size, keep_special)

	//Obtain patterns from file
	dict_exists, dict_is_dir := path_describe(dict_path)
	if !dict_exists {
		print_error(fmt.Sprintf("%v does not exist\n", dict_path))
	}
	if dict_is_dir {
		mat.dir_add_pattern(dict_path, pad_str, is_raw)
	} else {
		mat.f_add_pattern(dict_path, pad_str, is_raw)
	}

	//Scenario: Training
	if training {
		if output_file == "" {
			tfunc_stdout(mat.patterns)
			os.Exit(0)
		}
		tfunc_file(mat.patterns, output_file)
		os.Exit(0)
	} else {
		//Scenario: Matching
		match_exists, match_is_dir := path_describe(match_path)
		if !match_exists {
			print_error(fmt.Sprintf("%v does not exist\n", match_path))
		}
		if match_is_dir {
			mat.dir_match_pattern(match_path, callback_func)
		} else {
			mat.f_match_pattern(match_path, callback_func)
		}
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

func callback_stdout(arr ...any) {
	fmt.Printf("%v %d %d %v %v\n", arr[0], arr[1], arr[2], arr[3], arr[4])
}

//TODO Rewrite callback_file function
func callback_file(arr ...any) {
	fmt.Printf("%v %d %d %v %v\n", arr[0], arr[1], arr[2], arr[3], arr[4])
}

//TODO Implement tfunc_stdout function
func tfunc_stdout(arr [][]rune) {
	for _, val := range arr {
		fmt.Println(string(val))
	}
}

//TODO Implement tfunc_file function
func tfunc_file(arr ...any) {

}
