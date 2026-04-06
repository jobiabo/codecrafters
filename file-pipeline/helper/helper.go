package helper

import (
	"strings"
)

func Upper(line string) string {
	return strings.ToUpper(line)
}

func Lower(line string) string {
	return strings.ToLower(line)
}

func TODO(line string) string {
	split := strings.Fields(line)

	for i := 0; i < len(split); i++ {
		if split[i] == "TODO:" || split[i] == "todo" {
			split[i] = "ACTION:"

		}
	}
	return strings.Join(split, " ")
}

func Reverse(line string) string {
	var modresult []string
	slit := strings.Fields(line)
	for i := len(slit) - 1; i >= 0; i-- {

		modresult = append(modresult, slit[i])
	}

	result := strings.Join(modresult, " ")
	return result
}
