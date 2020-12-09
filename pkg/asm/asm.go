package asm

import (
	"fmt"
	"strconv"
	"strings"
)

// Runner runs a list of asm instructions.
type Runner struct {
	Accumulator    int
	Offset         int
	Lines          []string
	InstructionMap map[string]InstructionExecuter
}

// InstructionExecuter will execute an instruction.
// Perhaps this should return success or failure.
type InstructionExecuter func(x int)

// Jmp performs a jump operation.
func (r *Runner) Jmp(x int) {
	r.Offset += x
}

// Acc performs a jump operation.
func (r *Runner) Acc(x int) {
	r.Accumulator += x
	r.Offset++
}

// Nop does nothing.
func (r *Runner) Nop(x int) {
	r.Offset++
}

// NewASMRunner creates a runner and sets up all known instructions
// and their respective executioners.
func NewASMRunner(lines []string) (*Runner, error) {
	a := &Runner{
		Lines: lines,
	}
	a.InstructionMap = map[string]InstructionExecuter{
		"jmp": a.Jmp,
		"nop": a.Nop,
		"acc": a.Acc,
	}
	return a, nil
}

// Run runs the code to its completion.
func (r *Runner) Run() (bool, error) {
	seen := make(map[int]struct{})
	for {
		if _, ok := seen[r.Offset]; ok {
			return false, fmt.Errorf("Already seen. Possible endless loop? Offset: %d, Acc: %d", r.Offset, r.Accumulator)
		}
		seen[r.Offset] = struct{}{}
		instruction := strings.Split(r.Lines[r.Offset], " ")
		op := instruction[0]
		inst, _ := strconv.Atoi(instruction[1])

		v, ok := r.InstructionMap[op]
		if !ok {
			return false, fmt.Errorf("Unknown instruction %s at offset %d", op, r.Offset)
		}

		// Run the instruction.
		v(inst)

		if r.Offset >= len(r.Lines) {
			fmt.Println("Got the answer: ", r.Accumulator)
			return true, nil
		}
	}
}
