package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/carepollo/sexlang/lexer"
	"github.com/carepollo/sexlang/token"
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

		for tok := lexer.NextToken(); tok.Type != token.EOF; tok = lexer.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}
