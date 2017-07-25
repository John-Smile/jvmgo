package control

import "jvmgo/ch05/rtda"
import "jvmgo/ch05/instructions/base"

type GOTO struct {
	base.BranchInstruction
}

func (self *GOTO) Execute(frame *rtda.Frame)  {
	base.Branch(frame, self.Offset)
}
