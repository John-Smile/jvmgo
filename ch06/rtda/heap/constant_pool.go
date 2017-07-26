package heap

import "fmt"
import "jvmgo/ch06/classfile"
type Constant interface {

}
type ConstantPool struct {
	class            *Class
	consts           [] Constant
}

func newConstantPool(class *Class, cfCp classfile.ConstantPool)  *ConstantPool {
	cpCount := len(cfCp)
	consts := make([] Constant, cpCount)
	rtCp := &ConstantPool{class, consts}
	for i := 1; i < cpCount; i++ {
		cpInfo := cfCp[i]
		switch cpInfo.(type) {

		}
	}
	return rtCp
}
func (self *ConstantPool) GetConstant(index uint) Constant  {
	if c := self.consts[index]; c != nil {
		return c
	}
	panic(fmt.Sprintf("No constants at index %d", index))
}
