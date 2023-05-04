package main

import (
	bufio "bufio"
	flag "flag"
	fmt "fmt"
	. "github.com/gnydick/gogcode/pkg/structs"
	. "github.com/gnydick/gogcode/pkg/utils"
	os "os"
)

var tools = ToolSet{}

var state = State{}
var cpuProfile = flag.String("cpuProfile", "", "write cpu profile to `file`")
var memProfile = flag.String("memProfile", "", "write memory profile to `file`")
var gcodeCommand = flag.String("gcodeCommand", "", "gcode command")
var rangeStart = flag.Float64("rangeStart", 1.0, "value to start")
var rangeEnd = flag.Float64("rangeEnd", 1.0, "value to end")

var numObjects = flag.Int("numObjects", 1, "number of objects")
var input = flag.String("input", "", "input file")
var output = flag.String("output", "", "output file")

func main() {
	flag.Parse()

	//var object_map map[string]float32

	//var interval = (*rangeEnd - *rangeStart) / float64(*numObjects)

	var value = *rangeStart
	i, err := os.Open(*input)
	Check(err)
	scanner := bufio.NewScanner(i)

	o, err := os.Create(*output)
	Check(err)

	bo := bufio.NewWriter(o)

	defer func() {
		i.Close()

	}()

	util := NewUtil()
	for scanner.Scan() {

		line := scanner.Text()
		instructions := util.GenGcode(line)
		for _, instruction := range instructions {
			if instruction.Command == "M486" {
				bo.WriteString(fmt.Sprintf("%s%f", value))
			}
			bo.WriteString((*instruction).Gcode() + "\n")

		}
	}
	println("Hello!")
	bo.Flush()
	o.Close()
}
