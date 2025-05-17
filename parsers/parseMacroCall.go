package parsers

import (
	"regexp"
	"strings"
)

type SyscallArg struct {
	Type  string // const, range, flags, struct, define, macro, vs.
	Value string // "./file", 1-100, flag1, vs.
}

type MacroCall struct {
	MacroName string
	FuncName  string
	Args      []SyscallArg
	OutputVar string
}

func ParserMacros(context []byte) []MacroCall {
	var macros []MacroCall

	// örnek eşleşme: $macro@function(args...) -> outputVar
	macroRe := regexp.MustCompile(`\$(\w+)@(\w+)\s*\((.*?)\)(?:\s*->\s*(\w+))?`)
	matches := macroRe.FindAllStringSubmatch(string(context), -1)

	for _, match := range matches {
		macroName := match[1]
		funcName := match[2]
		rawArgs := match[3]
		outputVar := match[4]

		args := parseSyscallArgs(rawArgs)

		macros = append(macros, MacroCall{
			MacroName: macroName,
			FuncName:  funcName,
			Args:      args,
			OutputVar: outputVar,
		})
	}

	return macros
}
func parseSyscallArgs(argStr string) []SyscallArg {
	var args []SyscallArg

	// Virgül ile ayır, ancak tırnak içindekileri bölme
	argRe := regexp.MustCompile(`\s*(\w+:"[^"]*"|\w+:[^,()]+|"[^"]*"|\w+)\s*`)
	matches := argRe.FindAllStringSubmatch(argStr, -1)

	for _, match := range matches {
		raw := strings.TrimSpace(match[1])
		var arg SyscallArg

		if strings.Contains(raw, ":") {
			// const:"baz", range:1-200
			parts := strings.SplitN(raw, ":", 2)
			arg.Type = strings.TrimSpace(parts[0])
			arg.Value = strings.Trim(parts[1], `" `)
		} else if strings.HasPrefix(raw, `"`) && strings.HasSuffix(raw, `"`) {
			// "baz"
			arg.Type = "const"
			arg.Value = strings.Trim(raw, `"`)
		} else {
			// raw değer: foo, flag1, vs.
			arg.Type = "raw"
			arg.Value = raw
		}

		args = append(args, arg)
	}

	return args
}
