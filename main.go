package main

import (
	"fmt"
	"io/ioutil"
	"kkf/parsers"
)

var allowedTypes = map[string]bool{
	"bool":   true,
	"char":   true,
	"uchar":  true,
	"int32":  true,
	"int64":  true,
	"uint32": true,
	"ul64":   true,
	"float":  true,
	"double": true,
}

func main() {
	content, err := ioutil.ReadFile("./formats/ornek.k")
	if err != nil {
		fmt.Println("Dosya okunamadı:", err)
		return
	}

	defines := parsers.ParseDefine(content)
	for _, entry := range defines {
		fmt.Printf("Key: %-5s Type: %-6s Value: %s\n", entry.Key, entry.Type, entry.Value)
	}

	fmt.Println("-------------------------------")
	structs := parsers.ParseStructs(content)
	for _, s := range structs {
		fmt.Println("Struct:", s.Name)
		for _, f := range s.Fields {
			fmt.Printf("  - Type: %-10s Name: %-10s Meta: %s\n", f.Type, f.Name, f.Meta)
		}
	}
	fmt.Println("-------------------------------")

	macroCalls := parsers.ParserMacros(content)
	for _, m := range macroCalls {
		fmt.Println("---------------")
		fmt.Printf("Macro:   %s\n", m.MacroName)
		fmt.Printf("Func:    %s\n", m.FuncName)
		fmt.Printf("Args:    %v\n", m.Args)
		fmt.Printf("Output:  %s\n", m.OutputVar)
	}

	return

}
