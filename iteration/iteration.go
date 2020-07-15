package iteration

import "strings"

func Repeat(unit string, time int) (repeated string) {
	for i := 0; i < time; i++ {
		repeated += unit
	}
	return
}

func RepeatWithStrings(unit string, time int) string {
	return strings.Repeat(unit, time)
}
