package references

import "jvmgo/ch08/instructions/base"
import "jvmgo/ch08/rtda"

type ARRAY_LENGTH struct {
	base.NoOperandsInstruction
}

func (self *ARRAY_LENGTH) Execute(frame *rtda.Frame)  {
	stack := frame.OperandStack()
	arrRef := stack.PopRef()
	if arrRef == nil {
		panic("java.lang.NulPointerException")
	}
	arrLen := arrRef.ArrayLength()
	stack.PushInt(arrLen)
}