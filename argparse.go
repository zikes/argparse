package argparse

// Parse accepts a string input and breaks down the string into arguments based
// on spacing and quotation.
func Parse(input string) []string {
	chars := []rune(input)
	output := []string{}
	current := ""
	inquote := false
	escapeNext := false

	for _, letter := range chars {
		if letter == ' ' && !inquote && !escapeNext {
			if len(current) > 0 {
				output = append(output, current)
				current = ""
			}
			continue
		}
		if letter == '"' && !escapeNext {
			if inquote {
				output = append(output, current)
				current = ""
				inquote = false
				continue
			} else {
				inquote = true
				continue
			}
		}
		if letter == '\\' && !escapeNext {
			escapeNext = true
			continue
		}
		current += string(letter)
		escapeNext = false
	}
	if len(current) > 0 {
		output = append(output, current)
	}
	return output
}
