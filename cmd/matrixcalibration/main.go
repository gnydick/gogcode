package main

import (
	"bufio"
	"bytes"
	"flag"
	"log"
	"os"
	"runtime/pprof"
	strconv "strconv"
	"strings"

	. "github.com/gnydick/gogcode/pkg/structs"
	. "github.com/gnydick/gogcode/pkg/utils"
)

var meta = GcodeMeta{}
var state = State{}
var cpuProfile = flag.String("cpuProfile", "", "write cpu profile to `file`")
var memProfile = flag.String("memProfile", "", "write memory profile to `file`")

var input = flag.String("input", "", "input file")
var output = flag.String("output", "", "output file")
var xCount = flag.Int("xcount", 2, "number of objects along the x axis")
var yCount = flag.Int("ycount", 2, "number of objects along the y axis")
var xGcode = flag.String("xgcode", "M572 S", "what command to use along the x-axis")
var xRange = flag.String("xrange", "0:100", "starting and ending value `0:100`")
var yGcode = flag.String("ygcode", "M566 E", "what command to use along the y-axis")
var yRange = flag.String("yrange", "0:100", "starting and ending value `0:100`")
var zGcode = flag.String("zgcode", "M204 P", "what command to use along the z-axis")
var zRange = flag.String("zrange", "0:100", "starting and ending value `0:100`")

func main() {

	flag.Parse()
	xrange := strings.Split(*xRange, ":")

	xstart, e := strconv.ParseFloat(xrange[0], 64)
	Check(e)
	xend, e := strconv.ParseFloat(xrange[1], 64)
	Check(e)

	yrange := strings.Split(*yRange, ":")
	ystart, e := strconv.ParseFloat(yrange[0], 64)
	Check(e)

	yend, e := strconv.ParseFloat(yrange[1], 64)
	Check(e)

	xDelta := (xend - xstart) / float64(*xCount)
	yDelta := (yend - ystart) / float64(*yCount)

	xSettings := make([]float64, *xCount)
	ySettings := make([]float64, *yCount)

	for x := 0; x < *xCount; x++ {
		xSettings[x] = xstart + float64(x)*xDelta
		for y := 0; y < *yCount; y++ {
			ySettings[y] = ystart + float64(y)*yDelta
		}
	}

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

	bb := bytes.Buffer{}
	var curInsts []*Instruction
	util := NewUtil()
	for scanner.Scan() {
		line := scanner.Text()

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
