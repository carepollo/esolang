package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/carepollo/sexlang/evaluator"
	"github.com/carepollo/sexlang/lexer"
	"github.com/carepollo/sexlang/parser"
)

const PROMPT = ">> "

// start the interpreter in interactive mode (read, eval, print, loop), basically console mode
func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	for {
		fmt.Print(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}
		line := scanner.Text()
		lexer := lexer.New(line)
		parser := parser.New(lexer)

		program := parser.ParseProgram()
		if len(parser.Errors()) != 0 {
			printParserErrors(out, parser.Errors())
			continue
		}

		evaluated := evaluator.Eval(program)
		if evaluated != nil {
			io.WriteString(out, evaluated.Inspect())
			io.WriteString(out, "\n")
		}
	}
}

// helper to show internal record of errors
func printParserErrors(out io.Writer, errors []string) {
	io.WriteString(out, "Error raised")
	for _, message := range errors {
		io.WriteString(out, "\t"+message+"\n")
	}
}
