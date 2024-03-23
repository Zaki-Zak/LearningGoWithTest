package numerals

import "strings"

func ConvertRoman(arabic int) string {
	var result strings.Builder
	for i := 0; i < arabic; i++ {
		result.WriteString("I")
	}
	return result.String()
}
