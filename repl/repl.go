// repl/repl.go

package repl

import (
	"BoinkLang/lexer"
	"BoinkLang/parser"
	"bufio"
	"fmt"
	"io"
)

// PROMPT defines the REPL prompt symbol.
const PROMPT = ">> "

// Start initializes the REPL, reads input, tokenizes it, parses it, and prints the output.
func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Fprintf(out, PROMPT) // Display the prompt
		scanned := scanner.Scan()
		if !scanned {
			return // Exit on input end
		}

		line := scanner.Text()
		l := lexer.New(line)
		p := parser.New(l)

		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors()) // Handle parser errors
			continue
		}

		io.WriteString(out, program.String()) // Print parsed program
		io.WriteString(out, "\n")
	}
}

// printParserErrors prints parser errors to the output.
func printParserErrors(out io.Writer, errors []string) {
	io.WriteString(out, "Beep boop! Something went brrr.\n")
	io.WriteString(out, "parser errors:\n")
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n") // Print each error message
	}
}
