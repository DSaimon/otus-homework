package analysis

import (
	"fmt"
	"strings"
)

func AnaliseFrequency(text string) {
	arr := strings.Split(text, " ")
	fmt.Printf("%v\n", arr)
}