package references

import "jvmgo/ch09/instructions/base"
import "jvmgo/ch09/rtda"
import "jvmgo/ch09/rtda/heap"

type INVOKE_INTERFACE struct {
	index                      uint16
}

func (self *INVOKE_INTERFACE) FetchOperands(reader *base.BytecodeReader)  {
	self.index = reader.ReadUint16()
	reader.ReadUint8()
	reader.ReadUint8()
}
func (self *INVOKE_INTERFACE) Execute(frame *rtda.Frame)  {
	cp := frame.Method().Class().ConstantPool()
	methodRef := cp.GetConstant(self.index).(*heap.InterfaceMethodRef)
	resovledMethod := methodRef.ResolvedInterfaceMethod()
	if resovledMethod.IsStatic() || resovledMethod.IsPrivate() {
		panic("java.lang.IncompatibleClassChangeError")
	}

	ref := frame.OperandStack().GetRefFromTop(resovledMethod.ArgSlotCount() - 1)
	if ref == nil {
		panic("java.lang.NullPointerException")
	}
	if !ref.Class().IsImplements(methodRef.ResolvedClass()) {
		panic("java.lang.IncompatibleClassChangeError")
	}

	methodToBoInvoked := heap.LookupMethodInClass(ref.Class(), methodRef.Name(), methodRef.Descriptor())
	if methodToBoInvoked == nil || methodToBoInvoked.IsAbstract() {
		panic("java.lang.AbstractMethodError")
	}
	if !methodToBoInvoked.IsPublic() {
		panic("java.lang.IllegalAccessError")
	}
	base.InvokeMethod(frame, methodToBoInvoked)
}
