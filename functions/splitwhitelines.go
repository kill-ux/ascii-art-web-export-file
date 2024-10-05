package ascii

// splite with lines
func SplitWhitelines(s string) []string {
	var result []string
	var currentLine string

	for i := 0; i < len(s); i++ {
		if s[i] == 13 {
			continue
		}
	
		if s[i] == 10 {
			if len(currentLine) > 0 {
				result = append(result, currentLine)
			}
			result = append(result, "\n")
			currentLine = ""
		} else {
			currentLine += string(s[i])
		}
	}

	// Append the last line if there is any
	if len(currentLine) > 0 {
		result = append(result, currentLine)
	}
	return result
}
