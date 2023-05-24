package iteration

import "strings"

const RepeatCount = 5

func Repeat(character string) string {
	output := make([]string, RepeatCount)
	for i := range output {
		output[i] = character
	}

	return strings.Join(output, "")
}
