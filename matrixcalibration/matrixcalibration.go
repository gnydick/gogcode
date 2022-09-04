package main

import (
	"bufio"
	"flag"
	"log"
	"os"
	"runtime/pprof"

	. "github.com/gnydick/gogcode/structs"
	. "github.com/gnydick/gogcode/utils"
)

var meta = GcodeMeta{}
var state = State{}
var cpuProfile = flag.String("cpuProfile", "", "write cpu profile to `file`")
var memProfile = flag.String("memProfile", "", "write memory profile to `file`")

var input = flag.String("input", "", "input file")
var output = flag.String("output", "", "output file")
var util = NewUtil()

var xgcode = flag.String("xgcode", "M572 S", "what command to use along the x-axis")

// xRange := flag.String("xRange", "0:100", "starting and ending value `0:100`")
// ygcode := flag.String("ygcode", "M566 E", "what command to use along the y-axis")
// yRange := flag.String("yRange", "0:100", "starting and ending value `0:100`")
// zgcode := flag.String("zgcode", "M204 P", "what command to use along the z-axis")
// zRange := flag.String("zRange", "0:100", "starting and ending value `0:100`")
func main() {
	flag.Parse()
	println(*input)
	println(*output)
	println(*xgcode)
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

	//bb := bytes.Buffer{}
	var curInst *Instruction
	for scanner.Scan() {
		line := scanner.Text()

		curInst = util.GenGcode(line)
		state.Update(curInst)

	}
	for _, obj := range util.GcodeMeta.Objects {
		println(obj)
	}

}
