package parsers

import (
	"regexp"
	"strings"
)

type Field struct {
	Type string
	Name string
	Meta string
}

type StructDef struct {
	Name   string
	Fields []Field
}

func ParseStructs(source []byte) []StructDef {
	structRe := regexp.MustCompile(`struct\s+([A-Za-z_][A-Za-z0-9_]*)\s*\{([^}]*)\}`)
	metaRe := regexp.MustCompile(`<\s*.*?\s*>`)
	matches := structRe.FindAllStringSubmatch(string(source), -1)
	var structs []StructDef

	for _, match := range matches {
		structName := strings.TrimSpace(match[1])
		body := match[2]
		lines := strings.Split(body, "\n")
		var fields []Field

		for _, line := range lines {
			line = strings.TrimSpace(line)
			if line == "" {
				continue
			}

			meta := metaRe.FindString(line)
			cleanLine := metaRe.ReplaceAllString(line, "")
			parts := strings.Fields(cleanLine)

			if len(parts) >= 2 {
				field := Field{
					Type: parts[0],
					Name: parts[1],
					Meta: strings.TrimSpace(meta),
				}
				fields = append(fields, field)
			}
		}

		structs = append(structs, StructDef{
			Name:   structName,
			Fields: fields,
		})
	}

	return structs
}
