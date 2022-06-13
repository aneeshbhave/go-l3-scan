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
	flag.StringVar(&pad_str, "pad", " ", "Padding string")

	flag.BoolVar(&training, "t", false, "Training mode instead of matching mode (default is match mode)")
	flag.BoolVar(&keep_special, "s", false, "Keep special characters instead of replacing them with pattern")
	flag.BoolVar(&is_raw, "r", false, "Is the truth file raw pattern data or actual data")

	flag.IntVar(&init_size, "i", 250, "Guesstimate of how many patterns may appear in a file (for optimization)")

	flag.Parse()

	//!DEBUG LINES
	fmt.Printf("%v\n%v\n%v\n%v\n", dict_path, match_path, output_file, pad_str)
	fmt.Printf("%v\n%v\n%v\n", training, keep_special, is_raw)
	fmt.Printf("%v\n", init_size)

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
		if is_raw {
			mat.dir_add_raw(dict_path)
		} else {
			mat.dir_add_pattern(dict_path, pad_str)
		}
	} else {
		if is_raw {
			mat.f_add_raw(dict_path)
		} else {
			mat.f_add_pattern(dict_path, pad_str)
		}
	}

	//Scenario: Training
	if training {
		if output_file == "" {
			tfunc_stdout(mat.patterns)
		}
		tfunc_file(mat.patterns, output_file)
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

func callback_stdout(arr ...any) {
	fmt.Printf("%v %d %d %v %v\n", arr[0], arr[1], arr[2], arr[3], arr[4])
}

//TODO Rewrite callback_file function
func callback_file(arr ...any) {
	fmt.Printf("%v %d %d %v %v\n", arr[0], arr[1], arr[2], arr[3], arr[4])
}

//TODO Implement tfunc_stdout function
//!SHOULD EXIT PROGRAM
func tfunc_stdout(arr [][]rune) {
	for _, val := range arr {
		fmt.Println(string(val))
	}
}

//TODO Implement tfunc_file function
//!SHOULD EXIT PROGRAM
func tfunc_file(arr ...any) {
	fmt.Printf("%#v", arr)
}
