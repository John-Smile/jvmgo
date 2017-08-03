package stack

import "jvmgo/ch10/rtda"
import "jvmgo/ch10/instructions/base"

type POP struct {
	base.NoOperandsInstruction
}
type POP2 struct {
	base.NoOperandsInstruction
}

func (self *POP) Execute(frame *rtda.Frame)  {
	stack := frame.OperandStack()
	stack.PopSlot()
}
func (self *POP2) Execute(frame *rtda.Frame)  {
	stack := frame.OperandStack()
	stack.PopSlot()
	stack.PopSlot()
}
