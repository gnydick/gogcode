package utils

import (
	"fmt"
	. "github.com/gnydick/gogcode/pkg/structs"
	"log"
	"regexp"
	"strconv"
)

var gcodeRegex = `(?P<command>[gm]\d+)`
var paramRegex = `(?:[a-z]\d+(?:\.\d+)*)|(?:\s+[a-z]"[\w \t]+")`
var paramsRegex = `(?P<parameters>(?:(?:\s+` + paramRegex + `))*)`
var commentRegex = `(?:;\s*(?P<comment>.*))`
var lineRegex = `^(?i)` + gcodeRegex + paramsRegex + `\s*` + commentRegex + `?$`

func ParseLine(line string) {
	//gcodeRe := regexp.MustCompile(`^(?i)(?P<command>[gm]\d+)(?P<parameters>(?:(?:\s+[a-z]\d+(?:\.\d+)*)|(?:\s+[a-z]"[\w \t]+"))*)\s*(?:;\s*(?P<comment>.*))?$`)
	gcodeRe := regexp.MustCompile(lineRegex)
	matches := gcodeRe.FindAllStringSubmatch(line, -1)
	//paramsRe := regexp.MustCompile(paramsRegex)

	if len(matches) == 1 {
		for i, submatch := range matches[0] {
			if i <= 1 {
				continue // skip the first element, which is the entire match
			}
			paramRe := regexp.MustCompile(paramRegex)
			params := paramRe.FindAllStringSubmatch(submatch, -1)
			for i, param := range params {

				fmt.Printf("\tParam: %s\n", i, param)
			}
		}

	}
	//if len(match) == 4 {
	//	params := line[match[2]:match[3]]
	//	paramsRe.FindAllStringSubmatchIndex(match[2], -1)
	//	for _, param := range params {
	//		println(param)
	//	}
	//	command := match[1]
	//	parameters := strings.FieldsFunc(match[2], func(r rune) bool {
	//		return r == ' ' && !strings.Contains("XYZFIJKPQRT", string(match[0][strings.Index(match[0], string(r))-1]))
	//	})
	//	comment := match[3]
	//	fmt.Printf("Command: %s\nParameters: %v\nComment: %s\n", command, parameters, comment)
	//}

}

func (u *Util) GenGcode(line string) []*Instruction {
	ParseLine(line)
	tokens := beautifyLine(line, u)
	var instructions []*Instruction
	instruction := NewInstruction()
	instructions = append(instructions, &instruction)
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
	return instructions
}
