package boogie

/*
Bytecode defines the instruction set of our virtual machine.
*/
type Bytecode int

const (
	// IADD is integer add.
	IADD Bytecode = 0 + iota
	// HALT the VM.
	HALT
)

/*
VM is a wrapper for all data and methods needed to make our virtual machine work.
*/
type VM struct {
	stack []int
	code  []Bytecode
	data  []int
	ip    int
	sp    int
	fp    int
}

/*
NewVM is a constructor that returns a reference to a new instance of VM.
Takes is the code, and a starting point for the instruction pointer.
*/
func NewVM(code []Bytecode, main int, dsize int) *VM {
	return &VM{
		code:  code,
		ip:    main,
		data:  make([]int, dsize),
		stack: make([]int, 100),
		sp:    -1,
	}
}

func (vm *VM) cpu() {
	// Fetch.
	opcode := vm.code[vm.ip]
	vm.ip++

	switch opcode {
	case HALT:
		return
	}
}
