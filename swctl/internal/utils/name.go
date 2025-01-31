package utils

import "strings"

func ConvertToSnakeCase(name string) string {
	var result strings.Builder
	for i, char := range name {
		if i > 0 {
			// Handle camelCase to snake_case
			if isUpper := char >= 'A' && char <= 'Z'; isUpper {
				// Don't add underscore if previous char was uppercase
				// (handles acronyms like "URL" or "HTTP")
				prevIsUpper := name[i-1] >= 'A' && name[i-1] <= 'Z'
				// Don't add underscore if next char is uppercase
				nextIsUpper := i+1 < len(name) && name[i+1] >= 'A' && name[i+1] <= 'Z'

				if !prevIsUpper || (!nextIsUpper && i+1 < len(name)) {
					result.WriteRune('_')
				}
			}
		}

		// Convert to lowercase and handle special characters
		switch {
		case char >= 'A' && char <= 'Z':
			result.WriteRune(char + ('a' - 'A'))
		case char >= 'a' && char <= 'z':
			result.WriteRune(char)
		case char >= '0' && char <= '9':
			result.WriteRune(char)
		case char == ' ' || char == '-' || char == '.':
			// Replace spaces, hyphens, and dots with underscore
			if result.Len() > 0 && result.String()[result.Len()-1] != '_' {
				result.WriteRune('_')
			}
		}
	}

	// Trim leading/trailing underscores and collapse multiple underscores
	converted := result.String()
	converted = strings.Trim(converted, "_")
	return strings.Join(strings.FieldsFunc(converted, func(r rune) bool {
		return r == '_'
	}), "_")
}
