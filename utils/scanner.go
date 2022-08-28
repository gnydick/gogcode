package utils

import (
	"bufio"
	"os"
)

func Scan(outputChan *chan string, input *string) {

	i, err := os.Open(*input)
	Check(err)
	scanner := bufio.NewScanner(i)
	defer i.Close()
	for scanner.Scan() {

		line := scanner.Text()
		*outputChan <- line

	}
}
