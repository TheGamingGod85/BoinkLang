// parser/parser_tracing.go

package parser

import (
	"fmt"
	"strings"
)

var traceLevel int = 0

const traceIdentPlaceholder string = "\t"

// identLevel returns the indentation level based on the trace depth.
func identLevel() string {
	return strings.Repeat(traceIdentPlaceholder, traceLevel-1)
}

// tracePrint prints a formatted trace message with indentation.
func tracePrint(fs string) {
	fmt.Printf("%s%s\n", identLevel(), fs)
}

// incIdent increases the trace level.
func incIdent() { traceLevel = traceLevel + 1 }

// decIdent decreases the trace level.
func decIdent() { traceLevel = traceLevel - 1 }

// trace marks the beginning of a traced function and increases indentation.
func trace(msg string) string {
	incIdent()
	tracePrint("BEGIN " + msg)
	return msg
}

// untrace marks the end of a traced function and decreases indentation.
func untrace(msg string) {
	tracePrint("END " + msg)
	decIdent()
}
