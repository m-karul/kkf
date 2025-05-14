package parsers

import (
	"regexp"
)

type DefineEntry struct {
	Key   string
	Type  string
	Value string
}

func ParseDefine(source []byte) []DefineEntry {
	defineRe := regexp.MustCompile(`#define\s+([a-zA-Z_][a-zA-Z0-9_]*)\[(\w+)\]\s*=\s*(?:"([^"]*)"|(\d+))`)
	matches := defineRe.FindAllStringSubmatch(string(source), -1)

	var entries []DefineEntry

	for _, match := range matches {
		value := match[3]
		if value == "" {
			value = match[4]
		}

		entries = append(entries, DefineEntry{
			Key:   match[1],
			Type:  match[2],
			Value: value,
		})
	}

	return entries
}
