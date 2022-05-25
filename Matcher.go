package main

import "unicode"

const ALPHABET = "A"
const NUMERICAL = "N"
const SPACE = "W"
const SPECIAL = "S"

type Matcher struct {
	patterns     []string
	keep_special bool
	keep_unique  bool
}

func new_matcher(init_capacity int, keep_special bool, keep_unique bool) *Matcher {
	return &Matcher{
		patterns:     make([]string, 0, init_capacity),
		keep_special: keep_special,
		keep_unique:  keep_unique,
	}
}

func (mat *Matcher) add_pattern(inp ...string) {
	mat.patterns = append(mat.patterns, mat.to_pat(inp[0]))
}

func (mat *Matcher) f_add_pattern(fpath string) {
	lines := read_file_arr(fpath)
	for _, line := range lines {
		mat.add_pattern(line)
	}
}

func (mat *Matcher) dir_add_pattern(dirpath string) {}

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
