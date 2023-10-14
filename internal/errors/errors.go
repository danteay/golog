package errors

import (
	"runtime/debug"
	"strings"
)

func GetStackTrace() []string {
	stack := strings.ReplaceAll(string(debug.Stack()), "\t", "")
	return strings.Split(stack, "\n")
}
