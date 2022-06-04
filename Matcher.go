package main

//*---------------------------------------------
//TODO Write code for dir_add_pattern function
//TODO Fix the dir_ls_rec function
//TODO Change mode of error handling
//TODO Handle Error for f_match_pattern
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
}

func init_matcher(init_capacity int, keep_special bool) *Matcher {
	return &Matcher{
		patterns:     make([][]rune, 0, init_capacity),
		keep_special: keep_special,
	}
}

func (mat *Matcher) add_pattern(inp string, pad string, is_raw bool) {
	to_app := pad + inp + pad
	if !is_raw {
		to_app = mat.to_pat(to_app)
	}
	mat.patterns = append(mat.patterns, []rune(to_app))
}

func (mat *Matcher) f_add_pattern(fpath string, pad string, is_raw bool) {
	lines, err := read_file_arr(fpath)
	if err != nil {
		log.Panicf("Error reading file at %v", fpath)
	}
	for _, line := range lines {
		mat.add_pattern(line, pad, is_raw)
	}
}

func (mat *Matcher) dir_add_pattern(dirpath string, pad string, is_raw bool) {
	files := dir_ls_rec(dirpath)
	for _, file := range files {
		mat.f_add_pattern(file, pad, is_raw)
	}
}

func (mat *Matcher) match_pattern(haystack string, fname string, callback_func func(...any)) {
	aho := new(goahocorasick.Machine)
	r_haystack := []rune(mat.to_pat(haystack))

	if err := aho.Build(mat.patterns); err != nil {
		log.Panicf("Error building trie")
	}

	matches := aho.MultiPatternSearch(r_haystack, false)
	for _, e := range matches {
		pattern := string(e.Word)
		i, j := e.Pos, e.Pos+len(pattern)
		word := haystack[i:j]
		callback_func(fname, i, j, pattern, word)
	}
}

func (mat *Matcher) f_match_pattern(fpath string, callback_func func(...any)) {
	data, err := read_file_str(fpath)
	if err != nil {
		return
	}
	mat.match_pattern(data, fpath, callback_func)
}

func (mat *Matcher) dir_match_pattern(dirpath string, callback_func func(...any)) {
	files := dir_ls_rec(dirpath)
	for _, file := range files {
		mat.f_match_pattern(file, callback_func)
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
