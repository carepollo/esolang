package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/carepollo/sexlang/compiler"
	"github.com/carepollo/sexlang/lexer"
	"github.com/carepollo/sexlang/parser"
	"github.com/carepollo/sexlang/vm"
)

const PROMPT = "--> "

// start the interpreter in interactive mode (read, eval, print, loop), basically console mode
func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	// env := object.NewEnvironment()

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

		comp := compiler.New()
		err := comp.Compile(program)

		if err != nil {
			fmt.Fprintf(out, "Compilation failed:\n %s\n", err)
			continue
		}

		machine := vm.New(comp.Bytecode())
		err = machine.Run()
		if err != nil {
			fmt.Fprintf(out, "executing bytecode failed:\n %s\n", err)
			continue
		}

		stackTop := machine.StackTop()
		io.WriteString(out, stackTop.Inspect())
		io.WriteString(out, "\n")
	}
}

// helper to show internal record of errors
func printParserErrors(out io.Writer, errors []string) {
	io.WriteString(out, "Error raised")
	for _, message := range errors {
		io.WriteString(out, "\t"+message+"\n")
	}
}
