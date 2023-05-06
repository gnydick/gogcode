package main

import (
	"bufio"
	"flag"
	"fmt"
	. "github.com/gnydick/gogcode/pkg/structs"
	. "github.com/gnydick/gogcode/pkg/utils"
	"os"
)

var tools = ToolSet{}
var input = flag.String("input", "", "input file")
var output = flag.String("output", "", "output file")

var startValue = flag.Float64("startValue", 0, "")
var endValue = flag.Float64("endValue", 0, "")

var objectCount = flag.Float64("objectCount", 0, "")

var changemap = make(map[string]float64)

func main() {

	flag.Parse()

	interval := (*endValue - *startValue) / *objectCount
	lastValue := *startValue

	//value_range := args[2]

	i, err := os.Open(*input)
	Check(err)
	scanner := bufio.NewScanner(i)

	o, err := os.Create(*output)
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
		for _, inst := range instructions {
			if inst.Command == "M486" {
				if _, ok := changemap[inst.OtherParams["S"]]; ok {
					continue
				} else {
					newValue := lastValue + interval
					changemap[inst.OtherParams["S"]] = newValue
					lastValue = newValue
				}
				bo.WriteString(`M201 E` + fmt.Sprintf("%f", changemap[inst.OtherParams["S"]]) + "\n")
			}
			new_line := (*inst).Gcode()
			bo.WriteString(new_line + "\n")
		}
	}
}
