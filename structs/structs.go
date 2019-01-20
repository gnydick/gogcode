package structs

import (
	"fmt"
	"strings"
)

type MotionPositioning int

const (
	relativeMotion MotionPositioning = 91
	absoluteMotion MotionPositioning = 90
)

type ExtruderPositioning int

const (
	relativeExtruding ExtruderPositioning = 83
	absoluteExtruding ExtruderPositioning = 82
)

type ToolSet struct {
	tools  map[int]Tool
	active int
}

const (
	off     int = 0
	active  int = 1
	standby int = 2
)

type State struct {
	MovementPositioning string
	zPosition           float64
}

func (s State) ZPosition() float64 {
	return s.zPosition
}

func (s *State) Update(inst *Instruction) {
	switch command := (*inst).Command; command {
	case "G90":
		(*s).MovementPositioning = "absolute"
	case "G91":
		(*s).MovementPositioning = "relative"
	case "G1":
		s.updateZ(inst)
	}

}

func (s *State) updateZ(inst *Instruction) {
	if (*s).MovementPositioning == "absolute" {
		if (*inst).HasCoordinate("Z") {
			(*s).zPosition = (*inst).CoordinateValue("Z")
		}
	} else if (*s).MovementPositioning == "relative" {
		if (*inst).HasCoordinate("Z") {
			(*s).zPosition += (*inst).CoordinateValue("Z")
		}
	}
}

type Tool struct {
	Number      int
	ActiveTemp  int
	StandbyTemp int
	Positioning ExtruderPositioning
	Position    map[string]float64
}

type Instruction struct {
	Command           string
	ToolSet           ToolSet
	Position          map[string]float64
	MotionPositioning MotionPositioning
	OtherParams       map[string]float64
}

func (i Instruction) Marshal() string {
	sb := strings.Builder{}

	sb.WriteString("$i.Command ")

	return sb.String()

}

func (i Instruction) HasCoordinate(axis string) bool {

	if _, ok := i.Position[axis]; ok {
		return true

	} else {
		return false
	}

}

func (i Instruction) Coordinate(axis string) string {
	return fmt.Sprintf("%s%f", axis, i.Position[axis])
}

func (i Instruction) CoordinateValue(axis string) float64 {

	return i.Position[axis]

}

func (i Instruction) MovementOnly() string {
	var fields = make([]string, len(i.Position)+3)
	fields[0] = i.Command
	fieldNo := 0
	for k, v := range i.Position {
		if k != "E" {
			fieldNo++
			fields[fieldNo] = fmt.Sprintf("%s%f", k, v)
		}
	}
	fieldNo++
	fields[fieldNo] = "F6000"
	fieldNo++
	fields[fieldNo] = "; ironing move"

	return strings.Join(fields, " ")

}
func (i Instruction) Gcode() string {
	var fields = make([]string, len(i.Position)+2)
	fields[0] = i.Command
	fieldNo := 0
	for k, v := range i.Position {
		fieldNo++
		fields[fieldNo] = fmt.Sprintf("%s%f", k, v)
	}
	fieldNo++
	fields[fieldNo] = ";"
	return strings.Join(fields, " ")

}

func NewInstruction() Instruction {
	i := Instruction{
		Position: make(map[string]float64),
	}
	return i
}
