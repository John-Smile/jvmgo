package references

import "jvmgo/ch11/instructions/base"
import "jvmgo/ch11/rtda"
import "jvmgo/ch11/rtda/heap"

type PUT_FIELD struct {
	base.Index16Instruction
}

func (self *PUT_FIELD) Execute(frame *rtda.Frame)  {
	currentMethod := frame.Method()
	currentClass := currentMethod.Class()
	cp := currentClass.ConstantPool()
	fieldRef := cp.GetConstant(self.Index).(*heap.FieldRef)
	field := fieldRef.ResolvedField()

	if field.IsStatic() {
		panic("java.lang.IncompatibleClassChangeError")
	}
	if field.IsFinal() {
		if currentClass != field.Class() || currentMethod.Name() != "<init>" {
			panic("java.lang.IllegalAccessError")
		}
	}

	descriptor := field.Descriptor()
	slotId := field.SlotId()
	stack := frame.OperandStack()
	switch descriptor[0] {
	case 'Z', 'B', 'C', 'S', 'I':
		val := stack.PopInt()
		ref := popFieldRef(stack)
		ref.Fields().SetInt(slotId, val)
	case 'F':
		val := stack.PopFloat()
		ref := popFieldRef(stack)
		ref.Fields().SetFloat(slotId, val)
	case 'J':
		val := stack.PopLong()
		ref := popFieldRef(stack)
		ref.Fields().SetLong(slotId, val)
	case 'D':
		val := stack.PopDouble()
		ref := popFieldRef(stack)
		ref.Fields().SetDouble(slotId, val)
	case 'L', '[':
		val := stack.PopRef()
		ref := popFieldRef(stack)
		ref.Fields().SetRef(slotId, val)
	}
}
func popFieldRef(stack *rtda.OperandStack) *heap.Object {
	ref := stack.PopRef()
	if ref == nil {
		panic("java.lang.NullPointerException")
	}
	return ref
}
