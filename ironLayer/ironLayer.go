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
var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to `file`")
var memprofile = flag.String("memprofile", "", "write memory profile to `file`")
var zStart = flag.Float64("zstart", 1.0, "height to start")
var input = flag.String("input", "", "input file")
var output = flag.String("output", "", "output file")
var util = NewUtil()

func main() {
	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
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

	defer o.Close()
	bb := bytes.Buffer{}
	var curInst *Instruction
	bb.WriteString("; First ironing buffer\n")
	bb.WriteString("G1 E-15 F6000\n")
	for scanner.Scan() {
		line := scanner.Text()

		curInst = util.GenGcode(line)
		state.Update(curInst)

		pipeInstruction(o, &bb, curInst, zStart)

		if line == ";BEFORE_LAYER_CHANGE" {
			flush(o, &bb, zStart)
		}
	}

}

func flush(o *os.File, bb *bytes.Buffer, zStart *float64) {
	if state.ZPosition() >= *zStart {
		pipeInstruction(o, bb, util.GenGcode("G1 E15 F2400"), zStart)
		pipeInstruction(o, bb, util.GenGcode("G1 F3000"), zStart)
		o.WriteString(bb.String())
		if len(DEBUG) > 0 {
			fmt.Println(bb.String())
		}
		bb.Reset()
		pipeInstruction(o, bb, util.GenGcode("G1 E-15 F6000"), zStart)
	}

}

func isExtrudeMove(inst *Instruction) bool {
	if ((*inst).HasCoordinate("X") || (*inst).HasCoordinate("Y")) && !(*inst).HasCoordinate("Z") && (*inst).HasCoordinate("E") {
		return true
	}
	return false
}

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

var DEBUG = os.Getenv("DEBUG")

func pipeInstruction(o *os.File, bb *bytes.Buffer, inst *Instruction, zStart *float64) {

	if state.ZPosition() >= *zStart {
		if isExtrudeMove(inst) && DetectTravel(inst) {
			o.WriteString((*inst).MovementOnly() + "\n")
			bb.WriteString((*inst).MovementOnly() + "\n")
			if len(DEBUG) > 0 {
				fmt.Println((*inst).MovementOnly())
			}
		} else {
			o.WriteString((*inst).Gcode() + "\n")
			bb.WriteString((*inst).Gcode() + "\n")
			if len(DEBUG) > 0 {
				fmt.Println((*inst).Gcode())
			}
		}

	} else {
		o.WriteString((*inst).Gcode() + "\n")
		if len(DEBUG) > 0 {
			fmt.Println((*inst).Gcode())
		}
	}

}
