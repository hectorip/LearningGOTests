package word

import "unicode"

func IsPalindrome(s string) bool {
	var letters []rune
	for _, r := range s {
		if unicode.IsLetter(r) {
			letters = append(letters, unicode.ToLower(r))
		}
	}
	l := len(letters)
	for i := range letters {
		if letters[i] != letters[l-1-i] {
			return false
		}
}
	return true
}

