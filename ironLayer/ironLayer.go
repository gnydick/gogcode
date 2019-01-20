package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"

	"runtime/pprof"

	. "github.com/gnydick/gogcode/gcode/structs"
	. "github.com/gnydick/gogcode/gcode/utils"
	"log"
	"os"
)

var state = State{}
var cpuProfile = flag.String("cpuProfile", "", "write cpu profile to `file`")
var memProfile = flag.String("memProfile", "", "write memory profile to `file`")
var zStart = flag.Float64("zstart", 1.0, "height to start")
var input = flag.String("input", "", "input file")
var output = flag.String("output", "", "output file")
var util = NewUtil()

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
	check(err)
	scanner := bufio.NewScanner(i)

	o, err := os.Create(*output)
	check(err)

	bo := bufio.NewWriter(o)

	defer func() {
		i.Close()
		bo.Flush()
		o.Close()
	}()

	bb := bytes.Buffer{}
	var curInst *Instruction
	bb.WriteString("; First ironing buffer\n")
	bb.WriteString("G1 E-15 F6000\n")
	for scanner.Scan() {
		line := scanner.Text()

		curInst = util.GenGcode(line)
		state.Update(curInst)

		pipeInstruction(bo, &bb, curInst, zStart)

		if line == ";BEFORE_LAYER_CHANGE" {
			flush(bo, &bb, zStart)
		}
	}

}

func flush(bo *bufio.Writer, bb *bytes.Buffer, zStart *float64) {
	if state.ZPosition() >= *zStart {
		pipeInstruction(bo, bb, util.GenGcode("G1 E15 F2400"), zStart)
		pipeInstruction(bo, bb, util.GenGcode("G1 F3000"), zStart)
		bo.WriteString(bb.String())
		if len(DEBUG) > 0 {
			fmt.Println(bb.String())
		}
		bb.Reset()
		pipeInstruction(bo, bb, util.GenGcode("G1 E-15 F6000"), zStart)
	}
}

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

var DEBUG = os.Getenv("DEBUG")

func pipeInstruction(bo *bufio.Writer, bb *bytes.Buffer, inst *Instruction, zStart *float64) {

	if state.ZPosition() >= *zStart {
		if IsExtrudeMove(inst) {
			bo.WriteString((*inst).Gcode() + "\n")
			bb.WriteString((*inst).MovementOnly() + "\n")
			if len(DEBUG) > 0 {
				fmt.Println("output: " + (*inst).Gcode() + "\n")
				fmt.Println("buffer: " + (*inst).MovementOnly() + "\n")
			}
		} else {
			bo.WriteString((*inst).Gcode() + "\n")
			bb.WriteString((*inst).Gcode() + "\n")
			if len(DEBUG) > 0 {
				fmt.Println("output: " + (*inst).Gcode() + "\n")
				fmt.Println("buffer: " + (*inst).Gcode() + "\n")
			}
		}

	} else {
		bo.WriteString((*inst).Gcode() + "\n")
		if len(DEBUG) > 0 {
			fmt.Println("output: " + (*inst).Gcode() + "\n")
		}
	}

}
