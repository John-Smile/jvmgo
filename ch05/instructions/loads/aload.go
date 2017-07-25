package loads

import (
	"jvmgo/ch05/instructions/base"
	"jvmgo/ch05/rtda"
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
	_lload(frame, uint(self.Index))
}
func (self *ALOAD_1) Execute(frame *rtda.Frame)  {
	_lload(frame, 1)
}

