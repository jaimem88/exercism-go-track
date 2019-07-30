package luhn

import (
	"strconv"
	"strings"
)

// Valid checks if the input is valid according to the Luhn algorithm
func Valid(in string) bool {
	in = strings.ReplaceAll(in, " ", "")
	l := len(in)
	if l < 2 {
		return false
	}
	var sum int
	// even number digits
	if l%2 == 0 {
		// calculate luhn for even length input
		sum = luhnSum(in, false)
	} else {
		// calculate luhn for odd length input
		sum = luhnSum(in, true)
	}
	return sum%10 == 0
}

// luhnSum -
// if even length double every number in an even position, if >9 substract 9
// for odd length entries the above condition is true for every number in an odd position
// the last digit in the entry is not changed
// return the sum of all the digits including the sum
func luhnSum(in string, odd bool) int {
	var sum int
	for k, c := range in {
		d, err := strconv.Atoi(string(c))
		if err != nil {
			return -1
		}
		if k == len(in)-1 {
			sum += d
		} else if odd && k%2 != 0 {
			sum += checkDouble(d)
		} else if !odd && k%2 == 0 {
			sum += checkDouble(d)
		} else {
			sum += d
		}
	}
	return sum
}

// checkDouble doubles i, if i > 9 it substracts 9
func checkDouble(i int) int {
	t := i * 2
	if t > 9 {
		t -= 9
	}
	return t
}
