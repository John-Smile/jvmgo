package control

import "jvmgo/ch08/rtda"
import "jvmgo/ch08/instructions/base"

type GOTO struct {
	base.BranchInstruction
}

func (self *GOTO) Execute(frame *rtda.Frame)  {
	base.Branch(frame, self.Offset)
}
