package parsers

import (
	"regexp"
	"strings"
)

type FlagDef struct {
	FlagName string
	Values   []string
}

func ParseFlags(content []byte) []FlagDef {
	var results []FlagDef

	// Satırda "#flags flag=AAA,BBB,CCC,DD" formatını yakala
	re := regexp.MustCompile(`#Flags\s+(\w+)\s*=\s*([A-Za-z0-9_,]+)`)
	matches := re.FindAllStringSubmatch(string(content), -1)

	for _, match := range matches {
		flagName := match[1]
		values := strings.Split(match[2], ",")

		results = append(results, FlagDef{
			FlagName: flagName,
			Values:   values,
		})
	}

	return results
}
