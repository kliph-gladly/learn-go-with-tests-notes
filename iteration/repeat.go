package iteration

import "strings"

const REPEAT_COUNT = 5

func Repeat(character string) string {
	output := make([]string, REPEAT_COUNT)
	for i := range output {
		output[i] = character
	}
	return strings.Join(output, "")
}
