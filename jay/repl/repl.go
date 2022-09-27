package repl

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"jay/jay/lexer"
	"jay/jay/token"
)

const prompt = "->> "

func Start(ctx context.Context, in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		select {
		case <-ctx.Done():
			return
		default:
		}

		fmt.Fprint(out, prompt)

		if scanned := scanner.Scan(); !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)

		for t := l.NextToken(); t.Type != token.EOF; t = l.NextToken() {
			fmt.Fprintf(out, "%+v\n", t)
		}
	}
}
