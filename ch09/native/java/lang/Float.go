package lang

import "math"
import "jvmgo/ch09/native"
import "jvmgo/ch09/rtda"

func init()  {
	native.Register("java/lang/Float", "floatToRawIntBits", "(F)I", floatToRawIntBits)
}
func floatToRawIntBits(frame *rtda.Frame)  {
	value := frame.LocalVars().GetFloat(0)
	bits := math.Float32bits(value)
	frame.OperandStack().PushInt(int32(bits))
}
