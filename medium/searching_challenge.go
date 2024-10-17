// Given one sentence, like "Today is the greatest day ever!"
// Find the word that have the most occuring letters.
// E.g greatest have two e's two t's and it's before ever which have only two e's
package medium

import (
	"strings"
)

func SearchingChallenge(str string) string {

	examine := make(map[rune]int)
	maxWord := "-1"
	maxLength := 0
	wordList := strings.Split(str, " ")

	for _, element := range wordList {
		clear(examine)
		for _, letter := range element {
			examine[letter] += 1
		}

		tmp := 0
		for idx, _ := range examine {
			if examine[idx] > 1 {
				tmp += examine[idx]
			}
		}
		if tmp > maxLength {
			maxLength = tmp
			maxWord = element
		}

	}

	return maxWord
}
