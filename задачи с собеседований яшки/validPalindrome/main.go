package main

import "unicode"

func isPalindrome(s string) bool {
	left, right := 0, len(s)-1

	for left < right {
		l := rune(s[left])
		r := rune(s[right])

		if !unicode.IsLetter(l) && !unicode.IsDigit(l) {
			left++
		} else if !unicode.IsLetter(r) && !unicode.IsNumber(r) {
			right--
		} else if unicode.ToLower(l) == unicode.ToLower(r) {
			left++
			right--
		} else {
			return false
		}
	}

	return true
}