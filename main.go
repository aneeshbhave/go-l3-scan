package main

import (
	"fmt"
)

func main() {
	mat := init_matcher(1000, true, true)
	mat.f_add_pattern("./Data/TRUTH.TXT", " ")
	mat.dir_match_pattern("./Data/TextBlobs3000", callback_func)
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

func callback_func(fname string, i int, j int, pat string, match string) {
	fmt.Printf("%v %d %d %v %v\n", fname, i, j, pat, match)
}
