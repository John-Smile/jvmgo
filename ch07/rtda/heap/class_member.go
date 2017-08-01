package heap

import "jvmgo/ch07/classfile"

type ClassMember struct {
	accessFlags        uint16
	name               string
	descriptor         string
	class              *Class
}

func (self *ClassMember) copyMemberInfo(memberInfo *classfile.MemberInfo)  {
	self.accessFlags = memberInfo.AccessFlags()
	self.name = memberInfo.Name()
	self.descriptor = memberInfo.Descriptor()
}
func (self *ClassMember) isAccessibleTo(d *Class) bool {
	if self.IsPublic(){
		return true
	}
	c := self.class
	if self.IsProtected() {
		return d == c || d.IsSubClassOf(c) || c.GetPackageName() == d.GetPackageName()
	}
	if !self.IsPrivate() {
		return c.GetPackageName() == d.GetPackageName()
	}
	return d == c
}
func (self *ClassMember) IsPublic() bool  {
	return 0 != self.accessFlags & ACC_PUBLIC
}
func (self *ClassMember) IsPrivate() bool  {
	return 0 != self.accessFlags & ACC_PRIVATE
}
func (self *ClassMember) IsProtected() bool  {
	return 0 != self.accessFlags & ACC_PROTECTED
}
func (self *ClassMember) IsStatic() bool  {
	return 0 != self.accessFlags & ACC_STATIC
}
func (self *ClassMember) IsFinal() bool  {
	return 0 != self.accessFlags & ACC_FINAl
}
func (self *ClassMember) IsAbstract() bool  {
	return 0 != self.accessFlags & ACC_ABSTRACT
}

func (self *ClassMember) Descriptor() string  {
	return self.descriptor
}
func (self *ClassMember) Class() *Class  {
	return self.class
}
func (self *ClassMember) Name() string  {
	return self.name
}