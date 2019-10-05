package analysis

import (
	"sort"
	"strings"
	"unicode"
)

type pair struct {
	Key   string
	Value int
}

func comparator(c rune) bool {
	return c == ' ' || unicode.IsPunct(c)
}

// AnaliseFrequency make frequency analise of text. Returns 10 words or less with highest frequency. Words with similar
// frequency sorted by alphabet.
func AnaliseFrequency(text string) []string {
	var maxResultLength = 10

	// Split text to array and create map of words with frequency
	words := strings.FieldsFunc(text, comparator)
	frequencyMap := map[string]int{}
	for _, word := range words {
		frequencyMap[strings.ToLower(word)]++
	}

	// make slice with capacity = len(frequencyMap) and length = 0 to avoid allocations
	pairs := make([]pair, 0, len(frequencyMap))
	for key, value := range frequencyMap {
		pairs = append(pairs, pair{key, value})
	}

	sort.Slice(pairs, func(i, j int) bool {
		if pairs[i].Value > pairs[j].Value {
			return true
		}
		if pairs[i].Value < pairs[j].Value {
			return false
		}
		return pairs[i].Key < pairs[j].Key
	})

	resultLength := len(pairs)
	if resultLength > maxResultLength {
		resultLength = maxResultLength
	}

	// make slice with capacity = resultLength and length = 0 to avoid allocations
	result := make([]string, 0, resultLength)
	for i := range pairs[:resultLength] {
		result = append(result, pairs[i].Key)
	}

	return result
}
