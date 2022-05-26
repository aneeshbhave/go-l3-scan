package main

import (
	"fmt"
)

func main() {
	mat := init_matcher(10, true, true)
	mat.add_pattern("Aneesh", " ")
	mat.match_pattern("Aneesh Bhavee helloo whatup eeeee", callback_func)
}

func callback_func(pos int, val string) {
	fmt.Printf("%d %v\n", pos, val)
}
