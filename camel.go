package camel

import (
	"unicode"
)

const (
	nonletter = 0
	lower     = 1
	upper     = 2
)

func Strip(s string) string {
	if len(s) == 0 {
		return s
	}
	res := ""
	lastKind := 0
	for _, ch := range s {
		kind := 0
		if unicode.IsLetter(ch) {
			if unicode.IsUpper(ch) {
				kind = upper
			} else {
				kind = lower
			}
		} else {
			kind = nonletter
		}

		if (len(res) > 0) && (((lastKind == nonletter) && (kind != nonletter)) || ((lastKind == upper) && (kind == nonletter)) || ((lastKind == lower) && (kind != lower))) {
			res = res + "_"
		}
		res = res + string(unicode.ToLower(ch))
		lastKind = kind
	}
	return res
}
