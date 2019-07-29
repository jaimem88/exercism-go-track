/*
package bob implements Bob's behaviour as a lackadaisical teenager.
this implementation uses regular expresions to find the correct response
*/
package bob

import (
	"regexp"
	"strings"
)

var (
	letterRegexp  = regexp.MustCompile("[a-zA-Z]")
	controlRegexp = regexp.MustCompile("^[\n\r\t\\s]+$")
)

// Hey replies based on the remark
func Hey(remark string) string {
	remark = strings.TrimSpace(remark)

	if controlRegexp.MatchString(remark) || len(remark) == 0 {
		return "Fine. Be that way!"
	}
	if remark[len(remark)-1] == '?' {
		if tmp := strings.ToUpper(remark); tmp == remark && letterRegexp.MatchString(remark) {
			return "Calm down, I know what I'm doing!"
		}
		return "Sure."
	} else if letterRegexp.MatchString(remark) {
		if tmp := strings.ToUpper(remark); tmp == remark {
			return "Whoa, chill out!"
		}
	}

	return "Whatever."
}
