package main

import (
	"bufio"
	"fmt"
	"strconv"

	. "github.com/gnydick/gogcode/gcode/structs"
	. "github.com/gnydick/gogcode/gcode/utils"
	"log"
	"os"
)

func main() {
	state := State{}
	DEBUG := os.Getenv("DEBUG")
	args := os.Args[1:]
	zStart, _ := strconv.ParseFloat(args[0], 32)
	input := args[1]
	output := args[2]
	i, err := os.Open(input)
	check(err)
	scanner := bufio.NewScanner(i)

	o, err := os.Create(output)
	check(err)

	defer o.Close()

	var prevInst *Instruction
	var curInst *Instruction
	for scanner.Scan() {

		line := scanner.Text()
		curInst = GenGcode(line)
		state.Update(curInst)

		if state.ZPosition() < zStart {
			o.Write([]byte(line + "\n"))
			if len(DEBUG) > 0 {
				fmt.Println(line)
			}
			curInst = nil
			prevInst = nil
			continue
		}
		if isExtrudeMove(curInst) {
			if prevInst != nil {
				if len(DEBUG) > 0 {
					fmt.Println("")
					fmt.Println((*curInst).Gcode())
					fmt.Println((*prevInst).MovementOnly())
					fmt.Println((*curInst).MovementOnly())
					fmt.Println("")
				}
				o.Write([]byte((*curInst).Gcode() + "\n"))
				o.Write([]byte((*prevInst).MovementOnly() + "\n"))
				o.Write([]byte((*curInst).MovementOnly() + "\n"))
				prevInst = curInst
				continue
			} else {
				o.Write([]byte(line + "\n"))
				if len(DEBUG) > 0 {
					fmt.Println(line)
				}
				prevInst = curInst
				continue
			}
		}
		o.Write([]byte(line + "\n"))
		if len(DEBUG) > 0 {
			fmt.Println(line)
		}
		curInst = nil
		prevInst = nil
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
