package asm

type ASMRunner struct {
	Acc    int
	Offset int
	Lines  []string
}

// This should get ASMRunner as state.
var InstructionExecuter func() error

var InstructionMap = map[string]InstInstructionExecuter{
	"jmp": Jmp,
}

func (a *ASMRunner) Jmp() error {
	return nil
}

func NewASM(lines []string) (*ASMRunner, error) {
	return &ASMRunner{}, nil
}
