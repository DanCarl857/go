package bob

import (
	"strings"
)

const (
	calm     = "Calm down, I know what I'm doing"
	sure     = "Sure."
	whoa     = "Whoa, chill out!"
	fine     = "Fine. Be that way!"
	whatever = "Whatever."
)

// Hey should have a comment documenting it.
func Hey(remark string) string {

	q := isQuestion(remark)
	s := shouting(remark)
	st := statement(remark)

	// if remark is a question
	if q {
		// if remark is a question being shouted
		if s {
			return calm
		}
		return sure
	} else if s {
		return whoa
	} else if st {
		return whatever
	} else {
		return fine
	}
}

// function to check if a string is a question
func isQuestion(remark string) bool {
	i := strings.Index(remark, "?")
	l := len(remark)

	// Check if remark ends with a question mark
	if i > -1 && i == l {
		return true
	}
	return false
}

// function to check if a string is in all caps
func shouting(remark string) bool {
	for _, r := range remark {
		if r >= 'A' || r <= 'Z' {
			return true
		}
	}
	return false
}

func statement(remark string) bool {
	i := strings.Index(remark, "!")
	l := len(remark)

	if i > -1 && i == l {
		return true
	}
	return false
}
