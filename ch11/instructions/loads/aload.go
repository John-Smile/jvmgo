package loads

import (
	"jvmgo/ch11/instructions/base"
	"jvmgo/ch11/rtda"
	"fmt"
)

type ALOAD struct {
	base.Index8Instruction
}
type ALOAD_0 struct {
	base.NoOperandsInstruction
}
type ALOAD_1 struct {
	base.NoOperandsInstruction
}
type ALOAD_2 struct {
	base.NoOperandsInstruction
}
type ALOAD_3 struct {
	base.NoOperandsInstruction
}

func _aload(frame *rtda.Frame, index uint)  {
	val := frame.LocalVars().GetRef(index)
	frame.OperandStack().PushRef(val)
}
func (self *ALOAD) Execute(frame *rtda.Frame)  {
	_aload(frame, uint(self.Index))
}
func (self *ALOAD_0) Execute(frame *rtda.Frame)  {
	_aload(frame, 0)
}
func (self *ALOAD_1) Execute(frame *rtda.Frame)  {
	fmt.Printf("BEFORE ALOAD_1 OperandStack: %v, resolvedMethod\n", frame.OperandStack())
	_aload(frame, 1)
	fmt.Printf("AFTER ALOAD_1 OperandStack: %v, resolvedMethod\n", frame.OperandStack())
}
func (self *ALOAD_2) Execute(frame *rtda.Frame)  {
	_aload(frame, 2)
}
func (self *ALOAD_3) Execute(frame *rtda.Frame)  {
	_aload(frame, 3)
}

