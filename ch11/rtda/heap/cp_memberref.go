package heap

import "jvmgo/ch11/classfile"

type MemberRef struct {
	SymRef
	name               string
	descriptor         string
}

func (self *MemberRef) copyMemberRefInfo(refInfo *classfile.ConstantMemberrefInfo)  {
	self.className = refInfo.ClassName()
	self.name, self.descriptor = refInfo.NameAndDescriptor()
}
func (self *MethodRef) Name() string  {
	return self.name
}
func (self *MethodRef) Descriptor() string  {
	return self.descriptor
}
