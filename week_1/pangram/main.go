// Checking if a string is a pangram
package pangram

import "strings"

// IsPangram checks if given string is a pangram
func IsPangram(s string) bool {

	s = strings.ToLower(s)

	seen_alpha := map[rune]bool{
		'a': false, 'b': false, 'c': false, 'd': false, 'e': false, 'f': false,
		'g': false, 'h': false, 'i': false, 'j': false, 'k': false, 'l': false,
		'm': false, 'n': false, 'o': false, 'p': false, 'q': false, 'r': false,
		's': false, 't': false, 'u': false, 'v': false, 'w': false, 'x': false,
		'y': false, 'z': false,
	}

	for _, c := range s {
		if _, ok := seen_alpha[c]; ok {
			seen_alpha[c] = true
		}
	}

	for _, seen := range seen_alpha {
		if !seen {
			return false
		}
	}

	return true
}
