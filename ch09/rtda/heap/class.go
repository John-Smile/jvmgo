package heap

import "jvmgo/ch09/classfile"
import (
	"strings"
)


type Class struct {
	accessFlags          uint16
	name                 string
	superClassName       string
	interfaceNames       [] string
	constantPool         *ConstantPool
	fields               [] *Field
	methods              [] *Method
	loader               *ClassLoader
	superClass           *Class
	interfaces           [] *Class
	instanceSlotCount    uint
	staticSlotCount      uint
	staticVars           Slots
	initStarted          bool
	jClass               *Object
}

func newClass(cf *classfile.ClassFile) *Class {
	class := &Class{}
	class.accessFlags = cf.AccessFlags()
	class.name = cf.ClassName()
	class.superClassName = cf.SuperClassName()
	class.interfaceNames = cf.InterfaceNames()
	class.constantPool = newConstantPool(class, cf.ConstantPool())
	class.fields = newFields(class, cf.Fields())
	class.methods = newMethods(class, cf.Methods())
	return class
}
func (self *Class) IsPublic() bool  {
	return 0 != self.accessFlags & ACC_PUBLIC
}
func (self *Class) IsSuper() bool  {
	return 0 != self.accessFlags & ACC_SUPER
}
func (self *Class) IsInterface() bool  {
	return 0 != self.accessFlags & ACC_INTERFACE
}
func (self *Class) IsAbstract() bool  {
	return 0 != self.accessFlags & ACC_ABSTRACT
}
func (self *Class) IsAccessibleTo(other *Class) bool {
	return self.IsPublic() || self.GetPackageName() == other.GetPackageName()
}
func (self *Class) GetPackageName() string  {
	if i := strings.LastIndex(self.name, "/"); i >= 0 {
		return self.name[:i]
	}
	return ""
}
func (self *Class) NewObject() *Object {
	return newObject(self)
}
func newObject(class *Class) *Object  {
	return &Object{
		class: class,
		data: newSlots(class.instanceSlotCount),
	}
}

func (self *Class) Name() string  {
	return self.name
}
func (self *Class) SuperClass() *Class  {
	return self.superClass
}
func (self *Class) ConstantPool() *ConstantPool  {
	return self.constantPool
}
func (self *Class) Loader() *ClassLoader  {
	return self.loader
}
func (self *Class) JClass() *Object  {
	return self.jClass
}

func (self *Class) getStaticMethod(name, descriptor string) *Method  {
	for _, method := range self.methods {
		//fmt.Printf("IsStatic: %s, name: %s, descriptor: %s\n", method.IsStatic(), method.name, method.descriptor)
		//fmt.Printf("IsStatic: %s, name: %s, descriptor: %s\n", method.IsStatic(), method.name == name, method.descriptor == descriptor)
		if method.IsStatic() &&
		   method.name == name && method.descriptor == descriptor {
			return method
		}
	}
	return nil
}
func (self *Class) GetMainMethod() *Method  {
	return self.getStaticMethod("main", "([Ljava/lang/String;)V")
}
func (self *Class) StaticVars() Slots  {
	return self.staticVars
}
func (self *Class) InitStarted() bool  {
	return self.initStarted
}
func (self *Class) StartInit()  {
	self.initStarted = true
}
func (self *Class) GetClinitMethod() *Method  {
	return self.getStaticMethod("<clinit>", "()V")
}
func (self *Class) ArrayClass() *Class  {
	arrayClassName := getArrayClassName(self.name)
	return self.loader.LoadClass(arrayClassName)
}
func (self *Class) ComponentClass() *Class  {
	componentClassName := getComponentClassName(self.name)
	return self.loader.LoadClass(componentClassName)
}
func getComponentClassName(className string) string  {
	if className[0] == '[' {
		componentTypeDescriptor := className[1:]
		return toClassName(componentTypeDescriptor)
	}
	panic("Not array: " + className)
}
func (self *Class) isJlObject() bool  {
	return self.name == "java/lang/Object"
}
func (self *Class) isJlCloneable() bool {
	return self.name == "java/lang/Cloneable"
}
func (self *Class) isJioSerializable() bool  {
	return self.name == "java/io/Serializable"
}
func (self *Class) getField(name, descriptor string, isStatic bool) *Field  {
	for c := self; c != nil; c = c.superClass {
		for _, field := range c.fields {
			if field.IsStatic() == isStatic &&
			   field.name == name && field.descriptor == descriptor {
				return field
			}
		}
	}
	return nil
}
func (self *Class) JavaName() string  {
	return strings.Replace(self.name, "/", ".", -1)
}
func (self *Class) IsPrimitive() bool  {
	_, ok := primitiveTypes[self.name]
	return ok
}
func (self *Class) GetRefVar(fieldName, fieldDescriptor string) *Object  {
	field := self.getField(fieldName, fieldDescriptor, true)
	return self.staticVars.GetRef(field.slotId)
}
func (self *Class) SetRefVar(fieldName, fieldDescriptor string, ref *Object)  {
	field := self.getField(fieldName, fieldDescriptor, true)
	self.staticVars.SetRef(field.slotId, ref)
}
func (self *Class) GetInstanceMethod(name, descriptor string) *Method  {
	return self.getMethod(name, descriptor, false)
}
func (self *Class) getMethod(name, descriptor string, isStatic bool) *Method {
	for c := self; c != nil; c = c.superClass {
		for _, method := range c.methods {
			if method.IsStatic() == isStatic &&
			   method.name == name &&
			method.descriptor == descriptor {
				return method
			}
		}
	}
	return nil
}