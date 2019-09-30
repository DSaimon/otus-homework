package analysis

import (
	"fmt"
	"sort"
	"strings"
)

type Pair struct {
	Key   string
	Value int
}

func AnaliseFrequency(text string) []string {
	var maxResultLength = 10

	arr := strings.Split(text, " ")
	frequencyMap := map[string]int{}

	for _, word := range arr {
		frequencyMap[word] = frequencyMap[word] + 1
	}

	var pairs []Pair
	for key, value := range frequencyMap {
		pairs = append(pairs, Pair{key, value})
	}

	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].Value > pairs[j].Value
	})

	fmt.Printf("%v\n", pairs)


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