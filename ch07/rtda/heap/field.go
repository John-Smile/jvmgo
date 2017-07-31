package heap

import "jvmgo/ch07/classfile"

type Field struct {
	ClassMember
	constValueIndex uint16
	slotId          uint
}

func newFields(class *Class, cfFields [] * classfile.MemberInfo) [] *Field  {
	fields := make([] *Field, len(cfFields))
	for i, cfField := range cfFields {
		fields[i] = &Field{}
		fields[i].class = class
		fields[i].copyMemberInfo(cfField)
		fields[i].copyAttributes(cfField)
	}
	return fields
}
func (self *Field) isLongOrDouble() bool  {
	return self.descriptor == "J" || self.descriptor == "D"
}
func (self *Field) copyAttributes(cfField *classfile.MemberInfo)  {
	if valAttr := cfField.ConstantValueAttribute(); valAttr != nil {
		self.constValueIndex = valAttr.ConstantValueIndex()
	}
}
func (self *Field) ConstValueIndex() uint16 {
	return self.constValueIndex
}
func (self *Field) SlotId() uint  {
	return self.slotId
}
