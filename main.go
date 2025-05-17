package main

import (
	"fmt"
	"kkf/compile"
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

func test() {
	data := []byte(`
$open@first_open(const:"./file", range:1-200, flags:flag1, const: "baz") -> openfd
$open@first_open(const:"./file2", range:1-200, flags:flag1, const: "222") -> openfd
`)

	macros := parsers.ParserMacros(data)

	fmt.Printf("%s", macros[1].Args)
}
func main() {

	/*content, err := ioutil.ReadFile("./formats/ornek.k")
	if err != nil {
		fmt.Println("Dosya okunamadı:", err)
		return
	}
	var flags []parsers.FlagDef
	flags = parsers.ParseFlags(content)

	//for _, f := range flags {
	//	fmt.Println("Flag:", f.FlagName)
	//	fmt.Println("Values:", f.Values)
	//}

	var formatDefines []parsers.DefineEntry
	formatDefines = parsers.ParseDefine(content)
	//for _, entry := range formatDefines {
	//	fmt.Printf("Key: %-5s Type: %-6s Value: %s\n", entry.Key, entry.Type, entry.Value)
	//}

	//fmt.Println("-------------------------------")
	var formatStructs []parsers.StructDef
	formatStructs = parsers.ParseStructs(content)
	//for _, s := range formatStructs {
	//	fmt.Println("Struct:", s.Name)
	//	for _, f := range s.Fields {
	//		fmt.Printf("  - Type: %-10s Name: %-10s Meta: %s\n", f.Type, f.Name, f.Meta)
	//	}
	//}
	//fmt.Println("-------------------------------")
	var formatMacros []parsers.MacroCall
	formatMacros = parsers.ParserMacros(content)

	converter.ConvertToCLine(flags, formatDefines, formatStructs, formatMacros)

	return
	*/

	compile.GetOs()
}
