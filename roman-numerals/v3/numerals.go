package numerals

import "strings"

func ConvertRoman(arabic int) string {
	var result strings.Builder
	for arabic > 0 {
		switch {
		case arabic > 8:
			result.WriteString("IX")
			arabic -= 9
		case arabic > 4:
			result.WriteString("V")
			arabic -= 5
		case arabic > 3:
			result.WriteString("IV")
			arabic -= 4
		default:
			result.WriteString("I")
			arabic -= 1
		}
	}
	return result.String()
}
