package utils

import (
	"strings"
	"unicode"
)

func ToSnakeCase(s string) string {
	var builder strings.Builder

	// Iterate over each character in the string
	for i, char := range s {
		if unicode.IsUpper(char) {
			// If the character is upper case and not the first character, add an underscore
			if i > 0 {
				builder.WriteRune('_')
			}
			// Convert the upper case character to lower case
			char = unicode.ToLower(char)
		}
		// Add the character to the string builder
		builder.WriteRune(char)
	}

	// Return the snake case string
	return strings.TrimSpace(builder.String())
}
