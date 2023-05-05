/* Write a function that checks if a given string (case insensitive) is a palindrome.
A palindrome is a word, number, phrase, or other sequence of symbols that reads the same backwards as forwards,
such as madam or racecar, the date and time 12/21/33 12:21, and the sentence: "A man, a plan, a canal – Panama". */

package kata

import "strings"

func IsPalindrome(s string) bool {
	runes := []rune(strings.ToLower(s))
	length := len(runes)

	for i := 0; i < length/2; i++ {
		if runes[i] != runes[length-1-i] {
			return false
		}
	}

	return true
}
