package utils

import (
	"fmt"
	. "github.com/gnydick/gogcode/gcode/structs"
	"log"
	"math/rand"

	r "regexp"
	"strconv"
	"strings"
)

type Util struct {
	commandRe *r.Regexp
	paramRe   *r.Regexp
}

func NewUtil() *Util {
	return &Util{
		commandRe: r.MustCompile(`^(?P<command>[GM])(?P<value>-*[0-9.]+)$`),
		paramRe:   r.MustCompile(`^(?P<param>[A-Z])(?P<value>-*[0-9]+\.*[0-9]*)$`),
	}
}

func beautifyLine(line string) (tokens []string) {
	fields := strings.Fields(line)

	for _, field := range fields {
		if strings.HasPrefix(field, ";") {
			break
		} else if strings.HasSuffix(field, ";") {
			tokens = append(tokens, field[0:len(field)-2])
		} else {
			tokens = append(tokens, field)
		}
	}
	return
}

func IsExtrudeMove(inst *Instruction) bool {
	if (*inst).Command == "G1" {
		if ((*inst).HasCoordinate("X") || (*inst).HasCoordinate("Y")) && !(*inst).HasCoordinate("Z") && (*inst).HasCoordinate("E") {
			return true
		}
		return false
	} else {
		return false
	}
}

func (u Util) GenGcode(line string) *Instruction {
	tokens := beautifyLine(line)
	instruction := NewInstruction()
	if len(tokens) >= 1 {
		for _, commandMatches := range u.commandRe.FindAllStringSubmatchIndex(tokens[0], -1) {
			command := make([]byte, 0)
			instruction.Command = string(u.commandRe.ExpandString(command, "$command$value", tokens[0], commandMatches))
		}
		if len(tokens) >= 2 {
			for _, token := range tokens[1:] {

				for _, paramMatches := range u.paramRe.FindAllStringSubmatchIndex(token, -1) {
					param := make([]byte, 0)
					paramString := string(u.paramRe.ExpandString(param, "$param$value", token, paramMatches))
					value, _err := strconv.ParseFloat(string(u.paramRe.ExpandString(param, "$value", token, paramMatches)), 64)
					if _err != nil {
						log.Fatal(_err.Error())
					}

					instruction.Position[paramString[0:1]] = value
				}
			}
		}
	}
	return &instruction
}

func DetectTravel(gcode *Instruction) bool {

	if (*gcode).Command == "G1" {
		if (*gcode).HasCoordinate("X") || (*gcode).HasCoordinate("Y") {
			if (*gcode).HasCoordinate("Z") || (*gcode).HasCoordinate("E") {
				return false
			} else {
				return true
			}
		}
	}
	return false
}

func AddZHop(line *string, hop float32) string {
	sb := strings.Builder{}
	// TODO
	// get positioning mode before changing and changing back
	sb.WriteString("G91 ; set relative positioning ; added by gogcode\n")
	sb.WriteString(fmt.Sprintf("G1 Z%f ; hop! ; added by gogcode\n", hop))
	sb.WriteString("G90 ; set absolute positioning ; added by gogcode\n")
	sb.WriteString(fmt.Sprintf("%s\n", *line))
	sb.WriteString("G91 ; set relative positioning ; added by gogcode\n")
	sb.WriteString(fmt.Sprintf("G1 Z-%f ; hop! ; added by gogcode\n", hop))
	sb.WriteString("G90 ; set absolute positioning ; added by gogcode\n")
	return sb.String()
}

func GenRandMove(startX float64, startY float64, radius *float64, speed *int) {

	/*
		starting point is (startX, startY)
		pick a randum number, multiply it by 2 * radius
		add that to startX - radius


	*/

	distance := rand.Float64() * *radius * 2

	xDelta := distance * rand.Float64()

	xPos := startX - *radius + xDelta

	yDelta := distance * rand.Float64()

	yPos := startY - *radius + yDelta

	mmPerMin := *speed * 60
	output := fmt.Sprintf("G1 X%f Y%f F%d;", xPos, yPos, mmPerMin)

	fmt.Println(output)

}
