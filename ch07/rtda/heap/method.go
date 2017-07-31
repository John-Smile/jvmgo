package heap

import "jvmgo/ch07/classfile"

type Method struct {
	ClassMember
	maxStack        uint
	maxLocals       uint
	code            [] byte
}

func newMethods(class *Class, cfMethods [] *classfile.MemberInfo) [] *Method {
	methods := make([] *Method, len(cfMethods))
	for i, cfMethod := range cfMethods {
		methods[i] = &Method{}
		methods[i].class = class
		methods[i].copyMemberInfo(cfMethod)
		methods[i].copyAttributes(cfMethod)
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
func (self *Method) MaxStack() uint  {
	return self.maxStack
}
func (self *Method) MaxLocals() uint  {
	return self.maxLocals
}
func (self *Method) Code() [] byte  {
	return self.code
}
