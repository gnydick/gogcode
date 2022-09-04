package utils

import (
	"log"
	"strconv"

	. "github.com/gnydick/gogcode/structs"
)

func (u *Util) GenGcode(line string) []*Instruction {
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
