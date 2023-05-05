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
	//value_range := args[2]

	i, err := os.Open(input)
	Check(err)
	scanner := bufio.NewScanner(i)

	o, err := os.Create(output)
	Check(err)

	bo := bufio.NewWriter(o)

	//defer func() {
	//	if err := o.Close(); err != nil {
	//		panic(err)
	//	}
	//}()

	defer func() {
		i.Close()
		bo.Flush()
		o.Close()
	}()

	util := NewUtil()
	for scanner.Scan() {

		line := scanner.Text()
		instructions := util.GenGcode(line)
		for _, instruction := range instructions {
			if instruction.Command == "M486" {
				bo.WriteString((*instruction).Gcode() + "\n")
			}

		}
	}
}
