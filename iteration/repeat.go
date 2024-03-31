package iteration

import (
	"strings"
)

func Repeat(character string, count int) string {
	var strBuilder strings.Builder
	for i := 0; i < count; i++ {
		strBuilder.WriteString(character)
	}
	return strBuilder.String()
}
