package analysis

import (
	"sort"
	"strings"
)

type pair struct {
	Key   string
	Value int
}

// AnaliseFrequency make frequency analise of text. Returns 10 words or less with highest frequency
func AnaliseFrequency(text string) []string {
	var maxResultLength = 10

	arr := strings.Split(text, " ")
	frequencyMap := map[string]int{}

	for _, word := range arr {
		frequencyMap[word] = frequencyMap[word] + 1
	}

	var pairs []pair
	for key, value := range frequencyMap {
		pairs = append(pairs, pair{key, value})
	}

	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].Value > pairs[j].Value
	})

	var result []string
	var resultLength int
	if len(pairs) > maxResultLength {
		resultLength = maxResultLength
	} else {
		resultLength = len(pairs)
	}

	for i := 0; i < resultLength; i++ {
		result = append(result, pairs[i].Key)
	}

	return result
}
