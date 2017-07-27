package heap

import "fmt"
import "jvmgo/ch06/classfile"
import "jvmgo/ch06/classpath"

type ClassLoader struct {
	cp                    *classpath.ClassPath
	classMap              map[string]*Class
}
func NewClassLoader(cp *classpath.ClassPath) *ClassLoader {
	return &ClassLoader {
		cp:              cp,
		classMap:        make(map[string]*Class),
	}
}
func (self *ClassLoader) LoadClass(name string) *Class  {
	if class, ok := self.classMap[name]; ok {
		return class
	}
	return self.loadNonArrayClass(name)
}
func (self *ClassLoader) loadNonArrayClass(name string) *Class  {
	data, entry := self.readClass(name)
	class := self.defineClass(data)
	link(class)
	fmt.Printf("[Loaded %s from %s]\n", name, entry)
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
func verify(class Class)  {
	// todo
}
