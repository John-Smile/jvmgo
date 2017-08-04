package lang

import "jvmgo/ch11/native"
import "jvmgo/ch11/rtda"
import "jvmgo/ch11/rtda/heap"

func init()  {
	native.Register("java/lang/String", "intern","()Ljava/lang/String;", intern)
}
func intern(frame *rtda.Frame)  {
	this := frame.LocalVars().GetThis()
	interned := heap.InternString(this)
	frame.OperandStack().PushRef(interned)
}
