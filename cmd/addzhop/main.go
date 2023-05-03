package main

import (
	"bufio"
	"os"

	. "github.com/gnydick/gogcode/pkg/structs"
	. "github.com/gnydick/gogcode/pkg/utils"
)

var tools = ToolSet{}

func main() {
	args := os.Args[1:]
	input := args[0]
	output := args[1]

	i, err := os.Open(input)
	Check(err)
	scanner := bufio.NewScanner(i)

	o, err := os.Create(output)
	Check(err)

	defer func() {
		if err := o.Close(); err != nil {
			panic(err)
		}
	}()

	util := NewUtil()
	for scanner.Scan() {

		line := scanner.Text()
		instructions := util.GenGcode(line)
		for _, instruction := range instructions {
			travel := DetectTravel(instruction)

			if travel {
				o.WriteString(AddZHop(&line, .6))
			} else {
				o.WriteString(line + "\n")
			}
		}
	}
	println("Hello!")
}
