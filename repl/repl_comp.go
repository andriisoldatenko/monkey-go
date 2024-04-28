package repl

import (
	"bufio"
	"fmt"
	"github.com/andriisoldatenko/monkey-go/compiler"
	"github.com/andriisoldatenko/monkey-go/lexer"
	"github.com/andriisoldatenko/monkey-go/parser"
	"github.com/andriisoldatenko/monkey-go/vm"
	"io"
)

func StartComp(in io.Reader, out io.Writer) {
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
		comp := compiler.New()

		err := comp.Compile(program)

		if err != nil {
			fmt.Fprintf(out, "Whoops! Compilation failed:\n %s\n", err)
			continue
		}

		machine := vm.New(comp.Bytecode())

		err = machine.Run()

		if err != nil {
			fmt.Fprintf(out, "Whoops! Executing bytecode failed:\n %s\n", err)
			continue
		}

		lastPopped := machine.LastPoppedStackElem()
		io.WriteString(out, lastPopped.Inspect())
		io.WriteString(out, "\n")
	}
}
