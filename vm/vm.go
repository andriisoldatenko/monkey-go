package vm

import (
	"github.com/andriisoldatenko/monkey-go/code"
	"github.com/andriisoldatenko/monkey-go/compiler"
	"github.com/andriisoldatenko/monkey-go/object"
)

const StackSize = 2048

type VM struct {
	constants    []object.Object
	instructions code.Instructions

	stack []object.Object

	sp int // Always points to the next value. Top of stack is stack[sp-1]
}

func New(bytecode *compiler.Bytecode) *VM {
	return &VM{
		instructions: bytecode.Instructions,
		constants:    bytecode.Constants,

		stack: make([]object.Object, StackSize),
		sp:    0,
	}
}
