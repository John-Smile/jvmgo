package stores

import "jvmgo/ch11/instructions/base"
import "jvmgo/ch11/rtda"
import "jvmgo/ch11/rtda/heap"

type AASTORE struct {
	base.NoOperandsInstruction
}
type BASTORE struct {
	base.NoOperandsInstruction
}
type CASTORE struct {
	base.NoOperandsInstruction
}
type DASTORE struct {
	base.NoOperandsInstruction
}
type FASTORE struct {
	base.NoOperandsInstruction
}
type IASTORE struct {
	base.NoOperandsInstruction
}
type LASTORE struct {
	base.NoOperandsInstruction
}
type SASTORE struct {
	base.NoOperandsInstruction
}

func (self *IASTORE) Execute(frame *rtda.Frame)  {
	stack := frame.OperandStack()
	val := stack.PopInt()
	index := stack.PopInt()
	arrRef := stack.PopRef()

	checkNotNil(arrRef)
	ints := arrRef.Ints()
	checkIndex(len(ints), index)
	ints[index] = int32(val)
}
func (self *CASTORE) Execute(frame *rtda.Frame)  {
	stack := frame.OperandStack()
	val := stack.PopInt()
	index := stack.PopInt()
	arrRef := stack.PopRef()

	checkNotNil(arrRef)
	chars := arrRef.Chars()
	checkIndex(len(chars), index)
	chars[index] = uint16(val)
}
func (self *AASTORE) Execute(frame *rtda.Frame)  {
	stack := frame.OperandStack()
	ref := stack.PopRef()
	index := stack.PopInt()
	arrRef := stack.PopRef()

	checkNotNil(arrRef)
	refs := arrRef.Refs()
	checkIndex(len(refs), index)
	refs[index] = ref
}
func checkNotNil(ref *heap.Object)  {
	if ref == nil {
		panic("java.lang.NullPointerException")
	}
}
func checkIndex(arrLen int, index int32)  {
	if index < 0 || index >= int32(arrLen) {
		panic("ArrayIndexOutOfBoundException")
	}
}
