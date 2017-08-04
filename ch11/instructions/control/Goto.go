package control

import "jvmgo/ch11/rtda"
import "jvmgo/ch11/instructions/base"

type GOTO struct {
	base.BranchInstruction
}

func (self *GOTO) Execute(frame *rtda.Frame)  {
	base.Branch(frame, self.Offset)
}
