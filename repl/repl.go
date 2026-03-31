package repl

import (
	"Interpreter/lexer"
	"Interpreter/token"
	"bufio"
	"fmt"
	"io"
	"os"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}
		line := scanner.Text()
		l := lexer.New(line)
		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Println("%+v\\n", tok)

		}

	}

}
