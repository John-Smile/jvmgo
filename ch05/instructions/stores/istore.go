package stores

import "jvmgo/ch05/instructions/base"
import "jvmgo/ch05/rtda"

type ISTORE struct {
	base.Index8Instruction
}
type ISTORE_0 struct {
	base.NoOperandsInstruction
}
type ISTORE_1 struct {
	base.NoOperandsInstruction
}
type ISTORE_2 struct {
	base.NoOperandsInstruction
}
type ISTORE_3 struct {
	base.NoOperandsInstruction
}

func _istore(frame *rtda.Frame, index uint)  {
	val := frame.OperandStack().PopInt()
	frame.LocalVars().SetInt(index, val)
}
func (self *ISTORE) Execute(frame *rtda.Frame)  {
	_lstore(frame, uint(self.Index))
}
func (self *ISTORE_2) Execute(frame *rtda.Frame)  {
	_lstore(frame, 2)
}