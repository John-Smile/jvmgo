package heap

import "fmt"

func (self *Object) Clone() *Object  {
	return &Object{
		class: self.class,
		data: self.cloneData(),
	}
}
func (self *Object) cloneData() interface{} {
	switch self.data.(type) {
	case [] int8:
	   fmt.Printf("cloneData type [] int8\n")
		elements := self.data.([] int8)
		elements2 := make([] int8, len(elements))
		copy(elements2, elements)
		return elements2
	case [] int16:
		fmt.Printf("cloneData type [] int16\n")
		elements := self.data.([] int16)
		elements2 := make([] int16, len(elements))
		copy(elements2, elements)
		return elements2
	case [] uint16:
		fmt.Printf("cloneData type [] uint16\n")
		elements := self.data.([] uint16)
		elements2 := make([] uint16, len(elements))
		copy(elements2, elements)
		return elements2
	case [] int32:
		fmt.Printf("cloneData type [] int32\n")
		elements := self.data.([] int32)
		elements2 := make([] int32, len(elements))
		copy(elements2, elements)
		return elements2
	case [] int64:
		fmt.Printf("cloneData type [] int64\n")
		elements := self.data.([] int64)
		elements2 := make([] int64, len(elements))
		copy(elements2, elements)
		return elements2
	case [] float32:
		fmt.Printf("cloneData type [] float32\n")
		elements := self.data.([] float32)
		elements2 := make([] float32, len(elements))
		copy(elements2, elements)
		return elements2
	case [] float64:
		fmt.Printf("cloneData type [] float64\n")
		elements := self.data.([] float64)
		elements2 := make([] float64, len(elements))
		copy(elements2, elements)
		return elements2
	case [] *Object:
		fmt.Printf("cloneData type [] *Object\n")
		elements := self.data.([] *Object)
		elements2 := make([] *Object, len(elements))
		copy(elements2, elements)
		return elements2
	default:
		fmt.Printf("cloneData type [] Slots\n")
		slots := self.data.(Slots)
		slots2 := newSlots(uint(len(slots)))
		copy(slots2, slots)
		return slots
	}
}
