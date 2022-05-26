package main

//*---------------------------------------------
//TODO Write code for dir_add_pattern function
//TODO Fix the dir_ls_rec function
//TODO Change mode of error handling
//*---------------------------------------------

import (
	"log"
	"unicode"

	goahocorasick "github.com/anknown/ahocorasick"
)

//Used to convert text to proprietary patterns
const ALPHABET = "A"
const NUMERICAL = "N"
const SPACE = "W"
const SPECIAL = "S"

type Matcher struct {
	patterns     [][]rune
	keep_special bool
	keep_unique  bool
}

func init_matcher(init_capacity int, keep_special bool, keep_unique bool) *Matcher {
	return &Matcher{
		patterns:     make([][]rune, 0, init_capacity),
		keep_special: keep_special,
		keep_unique:  keep_unique,
	}
}

func (mat *Matcher) add_pattern(inp string, pad string) {
	mat.patterns = append(mat.patterns, []rune(mat.to_pat(pad+inp+pad)))
}

func (mat *Matcher) f_add_pattern(fpath string, pad string) {
	lines, err := read_file_arr(fpath)
	if err != nil {
		log.Panicf("Error reading file at %v", fpath)
	}
	for _, line := range lines {
		mat.add_pattern(line, pad)
	}
}

func (mat *Matcher) dir_add_pattern(dirpath string) {}

func (mat *Matcher) match_pattern(haystack string, callback_func func(int, string)) {
	aho := new(goahocorasick.Machine)
	r_haystack := []rune(mat.to_pat(haystack))

	if err := aho.Build(mat.patterns); err != nil {
		log.Panicf("Error building trie")
	}

	matches := aho.MultiPatternSearch(r_haystack, false)
	for _, e := range matches {
		callback_func(e.Pos, string(e.Word))
	}
}

func (mat Matcher) to_pat(inp string) string {
	pat := ""
	for _, c := range inp {
		if unicode.IsLetter(c) {
			pat += ALPHABET
			continue
		} else if unicode.IsNumber(c) {
			pat += NUMERICAL
			continue
		} else if unicode.IsSpace(c) {
			pat += SPACE
			continue
		} else {
			if mat.keep_special {
				pat += string(c)
				continue
			}
			pat += SPECIAL
			continue
		}
	}
	return pat
}
