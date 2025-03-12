// repl/repl.go

package repl

import (
	"BoinkLang/evaluator"
	"BoinkLang/lexer"
	"BoinkLang/parser"
	"BoinkLang/token"
	"bufio"
	"fmt"
	"io"
)

// PROMPT defines the REPL prompt symbol.
const PROMPT = ">> "

// StartParserMode initializes the REPL in parser mode.
func StartParserMode(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Fprintf(out, PROMPT)
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

// StartLexerMode initializes the REPL in lexer mode.
func StartLexerMode(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Fprintf(out, PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Fprintf(out, "%+v\n", tok)
		}
	}
}

// StartEvaluatorMode initializes the REPL in evaluator mode.
func StartEvaluatorMode(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Fprintf(out, PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)
		p := parser.New(l)
		program := p.ParseProgram()

		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
			continue
		}
		evaluated := evaluator.Eval(program)
		if evaluated != nil {
			io.WriteString(out, evaluated.Inspect())
			io.WriteString(out, "\n")
		}
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
