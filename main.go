package main

import (
	"fmt"
)

func main() {}

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

func callback_func(pos int, val string) {
	fmt.Printf("%d %v\n", pos, val)
}
