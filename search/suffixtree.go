// Using: http://stackoverflow.com/questions/22570839/create-a-suffix-tree-in-golang
package search

import (
	"index/suffixarray"
	"regexp"
	"strings"
)

func SuffixTree(words []string, pattern string) []string {
	// use \x00 to start each string
	joinedStrings := "\x00" + strings.Join(words, "\x00")
	sa := suffixarray.New([]byte(joinedStrings))

	// User has typed in "he"
	match, err := regexp.Compile("\x00" + pattern + "[^\x00]*")
	if err != nil {
		panic(err)
	}
	ms := sa.FindAllIndex(match, -1)
	var matchedWords []string

	for _, m := range ms {
		start, end := m[0], m[1]
		matchedWords = append(matchedWords, joinedStrings[start+1:end])
	}
	return matchedWords
}
