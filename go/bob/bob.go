package bob

import (
	"regexp"
	"strings"
	"unicode"
)

// Hey returns a string based on the remark provided
func Hey(remark string) string {
	remark = strings.TrimSpace(remark)
	anyLettersRegex, _ := regexp.Compile("[a-zA-Z]")
	anyLetters := anyLettersRegex.MatchString(remark)
	silence := len(remark) == 0

	if silence {
		return "Fine. Be that way!"
	} else if !anyLetters && isQuestion(remark) {
		return "Sure."
	} else if !anyLetters {
		return "Whatever."
	} else if isAllCaps(remark) && isQuestion(remark) {
		return "Calm down, I know what I'm doing!"
	} else if isAllCaps(remark) {
		return "Whoa, chill out!"
	} else if isQuestion(remark) {
		return "Sure."
	}
	return "Whatever."
}

func isQuestion(s string) bool {
	if s[len(s)-1:] == "?" {
		return true
	}

	return false
}

func isAllCaps(s string) bool {
	for _, r := range s {
		if unicode.IsLetter(r) {
			if unicode.IsUpper(r) == false {
				return false
			}
		}
	}

	return true
}
