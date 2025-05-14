package parsers

import "regexp"

type MacroCall struct {
	MacroName string
	FuncName  string
	Args      []string
	OutputVar string
}

func ParserMacros(context []byte) []MacroCall {

	var macros []MacroCall

	macroRe := regexp.MustCompile(`\$(\w+)#(\w+)\s*\((.*?)\)(?:\s*->\s*(\w+))?`)
	matches := macroRe.Copy().FindAllStringSubmatch(string(context), -1)

	for _, match := range matches {
		args := splitArgs(match[3])
		output := match[4]

		macros = append(macros, MacroCall{
			MacroName: match[1],
			FuncName:  match[2],
			Args:      args,
			OutputVar: output,
		})

	}
	return macros

}
func splitArgs(argStr string) []string {
	argRe := regexp.MustCompile(`"[^"]*"|\S+`)
	return argRe.FindAllString(argStr, -1)
}
