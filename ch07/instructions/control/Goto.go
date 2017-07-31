package control

import "jvmgo/ch07/rtda"
import "jvmgo/ch07/instructions/base"

type GOTO struct {
	base.BranchInstruction
}

func (self *GOTO) Execute(frame *rtda.Frame)  {
	base.Branch(frame, self.Offset)
}
