package lang

import "jvmgo/ch10/native"
import "jvmgo/ch10/rtda"
import "jvmgo/ch10/rtda/heap"

func init()  {
	native.Register("java/lang/String", "intern","()Ljava/lang/String;", intern)
}
func intern(frame *rtda.Frame)  {
	this := frame.LocalVars().GetThis()
	interned := heap.InternString(this)
	frame.OperandStack().PushRef(interned)
}
