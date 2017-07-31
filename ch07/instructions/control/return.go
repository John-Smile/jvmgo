package control

import "jvmgo/ch07/instructions/base"
import "jvmgo/ch07/rtda"

type RETURN struct {
	base.NoOperandsInstruction
}
type ARETURN struct {
	base.NoOperandsInstruction
}
type DRETURN struct {
	base.NoOperandsInstruction
}
type FRETURN struct {
	base.NoOperandsInstruction
}
type IRETURN struct {
	base.NoOperandsInstruction
}
type LRETURN struct {
	base.NoOperandsInstruction
}

func (self *RETURN) Execute(frame *rtda.Frame)  {
	frame.Thread().PopFrame()
}
func (self *IRETURN) Execute(frame *rtda.Frame)  {
	thread := frame.Thread()
	currentFrame := thread.PopFrame()
	invokerFrame := thread.TopFrame()
	retVal := currentFrame.OperandStack().PopInt()
	invokerFrame.OperandStack.PushInt(retVal)
}
