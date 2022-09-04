package main

import (
	"bufio"
	"bytes"
	"flag"
	"log"
	"os"
	"runtime/pprof"
	"strings"

	. "github.com/gnydick/gogcode/structs"
	. "github.com/gnydick/gogcode/utils"
)

var meta = GcodeMeta{}
var state = State{}
var cpuProfile = flag.String("cpuProfile", "", "write cpu profile to `file`")
var memProfile = flag.String("memProfile", "", "write memory profile to `file`")

var input = flag.String("input", "", "input file")
var output = flag.String("output", "", "output file")
var xcount = flag.Int("xcount", 2, "number of objects along the x axis")
var ycount = flag.Int("ycount", 2, "number of objects along the y axis")
var xGcode = flag.String("xGcode", "M572 S", "what command to use along the x-axis")
var xRange = flag.String("xRange", "0:100", "starting and ending value `0:100`")
var yGcode = flag.String("ygcode", "M566 E", "what command to use along the y-axis")
var yRange = flag.String("yRange", "0:100", "starting and ending value `0:100`")
var zGcode = flag.String("zgcode", "M204 P", "what command to use along the z-axis")
var zRange = flag.String("zRange", "0:100", "starting and ending value `0:100`")

func main() {

	flag.Parse()
	if *cpuProfile != "" {
		f, err := os.Create(*cpuProfile)
		if err != nil {
			log.Fatal("could not create CPU profile: ", err)
		}
		if err := pprof.StartCPUProfile(f); err != nil {
			log.Fatal("could not start CPU profile: ", err)
		}
		defer pprof.StopCPUProfile()
	}

	i, err := os.Open(*input)
	Check(err)
	scanner := bufio.NewScanner(i)

	o, err := os.Create(*output)
	Check(err)

	bo := bufio.NewWriter(o)

	defer func() {
		i.Close()
		bo.Flush()
		o.Close()
	}()
	x := *xcount
	y := *ycount
	if x == y {
	}

	commandMatrix := make([][]float64, x, y)

	bb := bytes.Buffer{}
	var curInsts []*Instruction
	util := NewUtil()
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "; printing object ") {
			objectId := line[18:]
			if objectId == objectId {
			}
			if *xcount**ycount == len(util.GcodeMeta.Objects) {
				println("have square number of objects")

			}
		}
		curInsts = util.GenGcode(line)

		state.Update(curInsts)
		for _, instruction := range curInsts {
			bb.WriteString(instruction.Gcode())
		}

	}
	for _, obj := range util.GcodeMeta.Objects {
		println(obj)

	}
	println(len(util.GcodeMeta.Objects))

}