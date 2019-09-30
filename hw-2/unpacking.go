package unpacking

import (
	"fmt"
	"strconv"
	)

func unpackString(str string) string {
	for _, char := range str {
		stringChar := string(char)
		_, err := strconv.ParseInt(stringChar, 10, 0)

		if err != nil {
			fmt.Printf("%s %v\n", stringChar, err)
		}
	}

	return ""
}


