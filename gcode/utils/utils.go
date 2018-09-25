package utils

import (
	"fmt"
	"github.com/gnydick/gogcode/gcode/structs"
	"strings"
)

func beautifyLine(line string) (tokens []string) {
	fields := strings.Fields(line)

	for _, field := range fields {
		if strings.HasPrefix(field, ";") {
			break
		} else {
			tokens = append(tokens, field)
		}
	}
	return
}

func DetectTravel(gcode structs.GcodeState, line string) (travel bool) {
	inMove := false
	tokens := beautifyLine(line)
	travel = false

	for _, token := range tokens {
		if inMove {
			if strings.HasPrefix(token, "Z") || strings.HasPrefix(token, "E") {
				travel = false
				break
			}
		} else if token == "G1" {
			inMove = true
			travel = true
		}
	}
	return
}

func BuildState(gcode structs.GcodeState, terms []string) {

}

func AddZHop(line string, hop float32) string {
	sb := strings.Builder{}
	// TODO
	// get positioning mode before changing and changing back
	sb.WriteString("G91 ; set relative positioning ; added by gogcode\n")
	sb.WriteString(fmt.Sprintf("G1 Z%f ; hop! ; added by gogcode\n", hop))
	sb.WriteString("G90 ; set absolute positioning ; added by gogcode\n")
	sb.WriteString(fmt.Sprintf("%s\n", line))
	sb.WriteString("G91 ; set relative positioning ; added by gogcode\n")
	sb.WriteString(fmt.Sprintf("G1 Z-%f ; hop! ; added by gogcode\n", hop))
	sb.WriteString("G90 ; set absolute positioning ; added by gogcode\n")
	return sb.String()
}
