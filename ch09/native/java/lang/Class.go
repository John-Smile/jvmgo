package lang

import "jvmgo/ch09/native"
import "jvmgo/ch09/rtda"
import (
	"jvmgo/ch09/rtda/heap"
	"fmt"
)

const jlClass = "java/lang/Class"

func init()  {
	native.Register(jlClass, "getPrimitiveClass", "(Ljava/lang/String;)Ljava/lang/Class;", getPrimitiveClas)
	native.Register(jlClass, "getName0", "()Ljava/lang/String;", getName0)
	native.Register(jlClass, "desiredAssertionStatus0", "(Ljava/lang/Class;)Z", desiredAssertionStatus0)
}
func getPrimitiveClas(frame *rtda.Frame)  {
	nameObj := frame.LocalVars().GetRef(0)
	name := heap.GoString(nameObj)
	loader := frame.Method().Class().Loader()
	class := loader.LoadClass(name).JClass()
	frame.OperandStack().PushRef(class)
}
func getName0(frame *rtda.Frame)  {
	fmt.Printf("getName0 frame.OperandStack:%v", frame.OperandStack())
	this := frame.LocalVars().GetThis()
	class := this.Extra().(*heap.Class)
	name := class.JavaName()
	nameObj := heap.JString(class.Loader(), name)
	frame.OperandStack().PushRef(nameObj)
}
func desiredAssertionStatus0(frame *rtda.Frame)  {
	frame.OperandStack().PushBoolean(false)
}