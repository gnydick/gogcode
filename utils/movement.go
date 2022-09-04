package utils

import (
	json "encoding/json"
	"fmt"
	"math/rand"
	r "regexp"
	"strings"

	. "github.com/gnydick/gogcode/structs"
)

type Axis int

const (
	X Axis = iota
	Y
	Z
	C
	E0
	E1
	E2
	E3
)

func (a Axis) String() string {
	names := [...]string{
		"X", "Y", "Z", "E1", "E2",
	}
	return names[a]
}

type Transform struct {
	axis     Axis
	distance float64
}

func NewTransform(axis Axis, distance float64) Transform {

	return Transform{
		axis:     axis,
		distance: distance,
	}
}

func (t Transform) String() string {
	return fmt.Sprintf("%s%f", t.axis, t.distance)
}

type Util struct {
	commandRe *r.Regexp
	paramRe   *r.Regexp
	toolRe    *r.Regexp
	GcodeMeta *GcodeMeta
}

func NewUtil() *Util {
	var objects []string
	gcodeMeta := GcodeMeta{
		Objects: objects,
	}

	return &Util{
		commandRe: r.MustCompile(`^(?P<command>[GM])(?P<value>-*[0-9.]+)$`),
		paramRe:   r.MustCompile(`^(?P<param>[A-Z])(?P<value>-*[0-9]+\.*[0-9]*)$`),
		toolRe:    r.MustCompile(`^(?P<tool>[T])(?P<value>-*[0-9.]+)$`),
		GcodeMeta: &gcodeMeta,
	}
}

func Move(transforms []Transform, speed float64, comment ...string) string {
	strs := make([]string, len(transforms))
	for i := 0; i < len(transforms); i++ {
		strs[i] = transforms[i].String()
	}
	return fmt.Sprintf("G1 %s ;%s", strings.Join(strs[:], " "), comment)

}

func beautifyLine(line string, util *Util) (tokens []string) {
	if strings.HasPrefix(line, ";") {
		i := *util
		i.processInfo(&line)
	}
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

// exact object declaration below
// ; object:{"name":"3DBenchy","id":"3DBenchy.stl id:0 copy 0","object_center":[137.499500,143.998000,0.000000],"boundingbox_center":[137.499500,143.998000,24.000000],"boundingbox_size":[60.000999,31.004000,48.000000]}
func (u *Util) processInfo(gcodeInfo *string) {

	if strings.HasPrefix(*gcodeInfo, ";") && len(*gcodeInfo) > 1 {
		infoObject := strings.Trim((*gcodeInfo)[1:len(*gcodeInfo)], " ")
		// record new object, read from json
		if strings.HasPrefix(infoObject, "object:") {
			objectJson := infoObject[7:len(infoObject)]
			var object map[string]string
			json.Unmarshal([]byte(objectJson), &object)
			u.GcodeMeta.Objects = append(u.GcodeMeta.Objects, object["id"])
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
