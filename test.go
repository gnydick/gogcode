package main

import (
	"fmt"
	"regexp"
)

func main() {
	re := regexp.MustCompile(`item\d+: (\d+),(\d+),(\d+)`)

	str := `item1: 1,2,3 item2: 4,5,6 item3: 7,8,9`

	matches := re.FindAllStringSubmatch(str, -1)

	for _, match := range matches {
		fmt.Printf("Match: %v\n", match)

		for i, submatch := range match {
			if i == 0 {
				continue // skip the first element, which is the entire match
			}
			fmt.Printf("\tSubmatch %d: %s\n", i, submatch)
		}
	}
}
