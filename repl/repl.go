package repl

import (
	"BoinkLang/evaluator"
	"BoinkLang/lexer"
	"BoinkLang/object"
	"BoinkLang/parser"
	"BoinkLang/token"
	"bufio"
	"fmt"
	"io"
)

// PROMPT defines the REPL prompt symbol.
const PROMPT = ">> "

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

// StartEvaluatorMode initializes the REPL in evaluator mode.
func StartEvaluatorMode(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	env := object.NewEnvironment()

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
		evaluated := evaluator.Eval(program, env)

		// Ignore NULL values in REPL output
		if evaluated != nil && evaluated.Type() != object.NULL_OBJ {
			io.WriteString(out, evaluated.Inspect())
			io.WriteString(out, "\n")
		}
	}
}

// RunFile executes BoinkLang code from a file
func RunFile(input string) {
	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()

	if len(p.Errors()) > 0 {
		fmt.Println("Parsing errors:")
		for _, err := range p.Errors() {
			fmt.Println("\t" + err)
		}
		return
	}

	// Evaluate the program
	env := object.NewEnvironment()
	result := evaluator.Eval(program, env)
	if result != nil {
		fmt.Println(result.Inspect())
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
