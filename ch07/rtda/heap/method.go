package heap

import "jvmgo/ch07/classfile"

type Method struct {
	ClassMember
	maxStack        uint
	maxLocals       uint
	code            [] byte
	argSlotCount    uint
}

func newMethods(class *Class, cfMethods [] *classfile.MemberInfo) [] *Method {
	methods := make([] *Method, len(cfMethods))
	for i, cfMethod := range cfMethods {
		methods[i] = &Method{}
		methods[i].class = class
		methods[i].copyMemberInfo(cfMethod)
		methods[i].copyAttributes(cfMethod)
		methods[i].calcArgSlotCount()
	}
	return methods
}
func (self *Method) copyAttributes(cfMethod *classfile.MemberInfo)  {
	if codeAttr := cfMethod.CodeAttribute(); codeAttr != nil {
		self.maxStack = uint(codeAttr.MaxStack())
		self.maxLocals = uint(codeAttr.MaxLocals())
		self.code = codeAttr.Code()
	}
}
func (self *Method) calcArgSlotCount()  {
	parsedDescriptor := parseMethodDescriptor(self.descriptor)
	for _, paramType := range parsedDescriptor.parameterTypes {
		self.argSlotCount++
		if paramType == "J" || paramType == "D" {
			self.argSlotCount++
		}
	}
	if !self.IsStatic() {
		self.argSlotCount++
	}
}
func (self *Method) MaxStack() uint  {
	return self.maxStack
}
func (self *Method) MaxLocals() uint  {
	return self.maxLocals
}
func (self *Method) Code() [] byte  {
	return self.code
}
func (self *Method) ArgSlotCount() uint  {
	return self.argSlotCount
}
