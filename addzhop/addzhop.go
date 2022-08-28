package main

import (
	"bufio"
	"os"

	. "github.com/gnydick/gogcode/structs"
	. "github.com/gnydick/gogcode/utils"
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

	for scanner.Scan() {

		util := NewUtil()
		line := scanner.Text()
		instruction := util.GenGcode(line)
		travel := DetectTravel(instruction)

		if travel {
			o.WriteString(AddZHop(&line, .6))
		} else {
			o.WriteString(line + "\n")
		}
	}

}
