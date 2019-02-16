package main

import (
	"bufio"
	"fmt"
	. "github.com/gnydick/gogcode/structs"
	. "github.com/gnydick/gogcode/utils"
	"os"
)

var tools = ToolSet{}

func main() {
	args := os.Args[1:]
	input := args[0]
	output := args[1]

	i, err := os.Open(input)
	check(err)
	scanner := bufio.NewScanner(i)

	o, err := os.Create(output)
	check(err)

	defer o.Close()

	for scanner.Scan() {
		instruction := NewInstruction()

		line := scanner.Text()
		GenGcode(&instruction, line)
		travel := DetectTravel(&instruction)

		if travel {
			o.WriteString(AddZHop(&line, .6))
		} else {
			o.WriteString(line + "\n")
		}
	}

}
func check(e error) {
	if e != nil {
		panic(e)
	}
}
