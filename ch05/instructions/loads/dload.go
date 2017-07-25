package loads

import (
	"jvmgo/ch05/instructions/base"
	"jvmgo/ch05/rtda"
)

type DLOAD struct {
	base.Index8Instruction
}
type DLOAD_0 struct {
	base.NoOperandsInstruction
}
type DLOAD_1 struct {
	base.NoOperandsInstruction
}
type DLOAD_2 struct {
	base.NoOperandsInstruction
}
type DLOAD_3 struct {
	base.NoOperandsInstruction
}

func _dload(frame *rtda.Frame, index uint)  {
	val := frame.LocalVars().GetDouble(index)
	frame.OperandStack().PushDouble(val)
}
func (self *DLOAD) Execute(frame *rtda.Frame)  {
	_lload(frame, uint(self.Index))
}
func (self *DLOAD_1) Execute(frame *rtda.Frame)  {
	_lload(frame, 1)
}

