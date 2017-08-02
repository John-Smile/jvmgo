package heap

import "jvmgo/ch08/classfile"
import "jvmgo/ch08/classpath"

type ClassLoader struct {
	cp                    *classpath.ClassPath
	verboseFlag           bool
	classMap              map[string]*Class
}
func NewClassLoader(cp *classpath.ClassPath, verboseFlag bool) *ClassLoader {
	return &ClassLoader {
		cp:              cp,
		verboseFlag:     verboseFlag,
		classMap:        make(map[string]*Class),
	}
}
func (self *ClassLoader) LoadClass(name string) *Class  {
	if class, ok := self.classMap[name]; ok {
		return class
	}
	if name[0] == '[' {
		return self.loadArrayClass(name)
	}
	return self.loadNonArrayClass(name)
}
func (self *ClassLoader) loadArrayClass(name string) *Class {
	class := &Class{
		accessFlags:            ACC_PUBLIC,
		name:                   name,
		loader:                 self,
		initStarted:            true,
		superClass:             self.LoadClass("java/lang/Object"),
		interfaces:             [] *Class{
			self.LoadClass("java/lang/Cloneable"),
			self.LoadClass("java/io/Serializable"),
		},
	}
	self.classMap[name] = class
	return class
}
func (self *ClassLoader) loadNonArrayClass(name string) *Class  {
	data, entry := self.readClass(name)
	class := self.defineClass(data)
	link(class)
	//fmt.Printf("[Loaded %s from %s]\n", name, entry)
	entry.String()
	return class
}
func (self *ClassLoader) readClass(name string) ([] byte, classpath.Entry)  {
	data, entry, err := self.cp.ReadClass(name)
	if err != nil {
		panic("java.lang.ClassNotFoundException: " + name)
	}
	return data, entry
}
func (self *ClassLoader) defineClass(data [] byte) *Class  {
	class := parseClass(data)
	class.loader = self
	resolveSuperClass(class)
	resolveInterfaces(class)
	self.classMap[class.name] = class
	return class
}
func parseClass(data [] byte) *Class  {
	cf, err := classfile.Parse(data)
	if err != nil {
		panic("java.lang.ClassFormatError")
	}
	return newClass(cf)
}
func resolveSuperClass(class *Class)  {
	if class.name != "java/lang/Object" {
		class.superClass = class.loader.LoadClass(class.superClassName)
	}
}
func resolveInterfaces(class *Class)  {
	interfaceCount := len(class.interfaceNames)
	if interfaceCount > 0 {
		class.interfaces = make([] *Class, interfaceCount)
		for i, interfaceName := range class.interfaceNames {
			class.interfaces[i] = class.loader.LoadClass(interfaceName)
		}
	}
}
func link(class *Class)  {
	verify(class)
	prepare(class)
}
func verify(class *Class)  {
	// todo
}
func prepare(class *Class)  {
	calcInstanceFieldSlotIds(class)
	calcStaticFieldSlotIds(class)
	allocAndInitStaticVars(class)
}
func calcInstanceFieldSlotIds(class *Class)  {
	slotId := uint(0)
	if class.superClass != nil {
		slotId = class.superClass.instanceSlotCount
	}
	for _, field := range class.fields {
		if !field.IsStatic() {
			field.slotId = slotId
			slotId++
			if field.isLongOrDouble() {
				slotId++
			}
		}
	}
	class.instanceSlotCount = slotId
}
func calcStaticFieldSlotIds(class *Class)  {
	slotId := uint(0)
	for _, field := range class.fields {
		if field.IsStatic() {
			field.slotId = slotId
			slotId++
			if field.isLongOrDouble() {
				slotId++
			}
		}
	}
	class.staticSlotCount = slotId
}
func allocAndInitStaticVars(class *Class)  {
	class.staticVars = newSlots(class.staticSlotCount)
	for _, field := range class.fields {
		if field.IsStatic() && field.IsFinal() {
			initStaticFinalVar(class, field)
		}
	}
}
func initStaticFinalVar(class *Class, field *Field)  {
	vars := class.staticVars
	cp := class.constantPool
	cpIndex := field.ConstValueIndex()
	slotId := field.SlotId()
	if cpIndex > 0 {
		switch field.Descriptor() {
		case "Z", "B", "C", "S", "I":
			val := cp.GetConstant(cpIndex).(int32)
			vars.SetInt(slotId, val)
		case "J":
			val := cp.GetConstant(cpIndex).(int64)
			vars.SetLong(slotId, val)
		case "F":
			val := cp.GetConstant(cpIndex).(float32)
			vars.SetFloat(slotId, val)
		case "D":
			val := cp.GetConstant(cpIndex).(float64)
			vars.SetDouble(slotId, val)
		case "Ljava/lang/String;":
			goStr := cp.GetConstant(cpIndex).(string)
			jStr := JString(class.Loader(), goStr)
			vars.SetRef(slotId, jStr)
		}
	}
}
