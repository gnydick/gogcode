package utils

import (
	. "github.com/gnydick/gogcode/pkg/structs"
	"regexp"
)

var gcodeRegex = `(?P<command>[gm]\d+)`
var paramRegex = `(?:[a-z]-*\d*\.\d+)|(?:[a-z]".+")`

var commentRegex = `(?:;\s*(?P<comment>.*))`
var lineRegex = `^(?i)(?P<command>[mgt]\d+)?\s*(?P<parameters>(?:(?:\s*)(?:[a-z]-?\d*(?:\.\d*)??|[a-z]"[\w-_. ]+(?:\.\d*)?"))+){0,1}(?:\s*);\s*(?P<comment>.*)?$`

var x = 0

func ParseLine(line string) (instruction *Instruction) {
	var inst Instruction
	lineRe := regexp.MustCompile(lineRegex)
	matches := lineRe.FindAllStringSubmatch(line, -1)
	parameters := make(map[string]string)

	if len(matches) == 1 {

		command := matches[0][1]
		params := matches[0][2]
		comment := matches[0][3]

		paramRe := regexp.MustCompile(`(?i)` + paramRegex)
		matchedParams := paramRe.FindAllString(params, -1)
		for _, param := range matchedParams {
			parameters[param[0:1]] = param[1:]
		}

		inst.Command = command
		inst.OtherParams = parameters
		inst.ToolSet = ToolSet{}
		inst.MotionPositioning = 0
		inst.Position = nil
		inst.Comment = comment
	}

	return &inst
}

func (u *Util) GenGcode(line string) []*Instruction {
	inst := ParseLine(line)
	instructions := []*Instruction{}
	instructions = append(instructions, inst)
	return instructions
}
