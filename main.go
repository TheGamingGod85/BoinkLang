// main.go

package main

import (
	"BoinkLang/repl"
	"bufio"
	"fmt"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! This is the BoinkLang programming language!\n", user.Username)
	fmt.Println("Choose mode: [rlpl] Lexer Mode | [rppl] Parser Mode | [repl] Evaluator Mode")

	// Read mode selection
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter mode: ")
	scanner.Scan()
	mode := scanner.Text()

	switch mode {
	case "rlpl":
		repl.StartLexerMode(os.Stdin, os.Stdout)
	case "rppl":
		repl.StartParserMode(os.Stdin, os.Stdout)
	case "repl":
		repl.StartEvaluatorMode(os.Stdin, os.Stdout)
	default:
		fmt.Println("Invalid mode! Exiting.")
		os.Exit(1)
	}
}
