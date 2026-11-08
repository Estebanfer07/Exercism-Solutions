package anagram

import (
	"slices"
	"strings"
)

func Detect(subject string, candidates []string) []string {

	anagrams := []string{}

	for _, candidate := range candidates {

		if len(candidate) != len(subject) || strings.ToLower(candidate) == strings.ToLower(subject) {
			continue
		}

		subjectChars := strings.Split(strings.ToLower(subject), "")
		for _, char := range candidate {

			charIndex := slices.Index(subjectChars, strings.ToLower(string(char)))
			if charIndex == -1 {
				break
			}
			subjectChars = slices.Delete(subjectChars, charIndex, charIndex+1)

			if len(subjectChars) == 0 {
				anagrams = append(anagrams, candidate)
			}

		}

	}

	return anagrams

}
