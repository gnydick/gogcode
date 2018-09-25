package main

import (
	"bufio"
	"github.com/gnydick/gogcode/gcode/structs"
	"github.com/gnydick/gogcode/gcode/utils"
	"os"
)

var tools = structs.ToolSet{}
var state = structs.GcodeState{}

func main() {
	state.ToolSet = tools
	args := os.Args[1:]
	filename := args[0]
	f, err := os.Open(filename)
	check(err)
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		if utils.DetectTravel(state, scanner.Text()) == true {
			print(utils.AddZHop(scanner.Text(), .6))
		} else {
			println(scanner.Text())
		}
	}

}
func check(e error) {
	if e != nil {
		panic(e)
	}
}
