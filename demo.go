package demo

import (
	"fmt"
	"strconv"
)

func intStringFmt(i int) string {
	return fmt.Sprintf("%d", i)
}

func intStringItoa(i int) string {
	return strconv.Itoa(i)
}

func intStringFormatInt(i int64) string {
	return strconv.FormatInt(i, 10)
}
