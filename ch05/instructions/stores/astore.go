package stores

import "jvmgo/ch05/instructions/base"
import "jvmgo/ch05/rtda"

type ASTORE struct {
	base.Index8Instruction
}
type ASTORE_0 struct {
	base.NoOperandsInstruction
}
type ASTORE_1 struct {
	base.NoOperandsInstruction
}
type ASTORE_2 struct {
	base.NoOperandsInstruction
}
type ASTORE_3 struct {
	base.NoOperandsInstruction
}

func _astore(frame *rtda.Frame, index uint)  {
	val := frame.OperandStack().PopRef()
	frame.LocalVars().SetRef(index, val)
}
func (self *ASTORE) Execute(frame *rtda.Frame)  {
	_lstore(frame, uint(self.Index))
}
func (self *ASTORE_2) Execute(frame *rtda.Frame)  {
	_lstore(frame, 2)
}