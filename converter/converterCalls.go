package converter

import (
	"fmt"
	"kkf/parsers"
	"math"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func findFlagIndex(formatFlags []parsers.FlagDef, findFlag string) int {
	for i, flag := range formatFlags {
		if flag.FlagName == findFlag {
			return i
		}
	}
	return -1
}
func getFlags(formatFlags []parsers.FlagDef, findFlag string) string {

	index := findFlagIndex(formatFlags, findFlag)
	rand.Seed(time.Now().UnixNano())

	return formatFlags[index].Values[rand.Intn(len(formatFlags[index].Values))]
}
func randInt64InRange(min, max int64) int64 {
	return rand.Int63n(max-min+1) + min
}

// randUint64InRange returns random uint64 in [min, max]
func randUint64InRange(min, max uint64) uint64 {
	return min + uint64(rand.Int63n(int64(max-min+1)))
}

// RandomValue generates a random integer value based on type and value string.
// value: e.g. "1-200" for ranges or ignored for int types
// typ: e.g. "range", "int8", "uint8", "int16", "uint16", "int32", "uint32", "int64", "uint64"
func RandomIntFromRange(value string) (uint64, bool) {
	switch value {
	case "int8":
		return uint64(int64(randInt64InRange(math.MinInt8, math.MaxInt8))), true
	case "uint8":
		return uint64(randUint64InRange(0, math.MaxUint8)), true

	case "int16":
		return uint64(int64(randInt64InRange(math.MinInt16, math.MaxInt16))), true
	case "uint16":
		return uint64(randUint64InRange(0, math.MaxUint16)), true

	case "int32":
		return uint64(int64(randInt64InRange(math.MinInt32, math.MaxInt32))), true
	case "uint32":
		return uint64(randUint64InRange(0, math.MaxUint32)), true

	case "int64":
		return uint64(int64(randInt64InRange(math.MinInt64, math.MaxInt64))), true
	case "uint64":
		// uint64 max range, dikkat
		return randUint64InRange(0, math.MaxUint64), true

	default:
		bounds := strings.Split(value, "-")
		if len(bounds) != 2 {
			return 0, false
		}
		start, err1 := strconv.ParseUint(strings.TrimSpace(bounds[0]), 10, 64)
		end, err2 := strconv.ParseUint(strings.TrimSpace(bounds[1]), 10, 64)
		if err1 != nil || err2 != nil || start > end {
			return 0, false
		}
		return randUint64InRange(start, end), true
	}
}
func ConvertToCLine(formatFlags []parsers.FlagDef, formatDefines []parsers.DefineEntry, formatStructs []parsers.StructDef, formatMacros []parsers.MacroCall) {
	if len(formatMacros) == 0 {
		fmt.Println("Hiç syscall makrosu yok.")
		return
	}

	// Rastgelelik başlat
	rand.Seed(time.Now().UnixNano())

	// 1. macro'yu seç
	selected := formatMacros[0]

	var calform []string
	calform = append(calform, selected.OutputVar)
	calform = append(calform, "=")
	calform = append(calform, selected.MacroName)
	calform = append(calform, "(")

	var finalArgs []string

	for _, arg := range selected.Args {
		switch arg.Type {
		case "Const":
			finalArgs = append(finalArgs, fmt.Sprintf(`"%s"`, arg.Value))

		case "Flags":
			fl := getFlags(formatFlags, arg.Value)
			finalArgs = append(finalArgs, fl)
		case "Range":
			if val, ok := RandomIntFromRange(arg.Value); ok {
				finalArgs = append(finalArgs, fmt.Sprintf("%d", val))
			}

		default:
			// struct, define, macro vb. şimdilik dokunma
			finalArgs = append(finalArgs, arg.Value)
		}
	}

	calform = append(calform, strings.Join(finalArgs, ","))
	calform = append(calform, ")")
	fmt.Println(calform)

}
