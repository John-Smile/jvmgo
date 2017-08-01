package rtda

import (
	"math"
	"jvmgo/ch08/rtda/heap"
)

type OperandStack struct {
	size   uint
	slots  [] Slot
}

func newOperandStack(maxStack uint) *OperandStack  {
	if maxStack > 0 {
		return &OperandStack{
			slots: make([] Slot, maxStack),
		}
	}
	return nil
}
func (self *OperandStack) PushInt(val int32)  {
	//fmt.Printf("before exec OperandStack length: %d, current Index: %d\n", len(self.slots), self.size)
	self.slots[self.size].num = val
	self.size++
	//fmt.Printf("after exec OperandStack length: %d, current Index: %d\n", len(self.slots), self.size)
}
func (self *OperandStack) PopInt() int32 {
	//fmt.Printf("before exec OperandStack length: %d, current Index: %d\n", len(self.slots), self.size)
	self.size--
	//fmt.Printf("after exec OperandStack length: %d, current Index: %d\n", len(self.slots), self.size)
	return self.slots[self.size].num
}
func (self *OperandStack) PushFloat(val float32)  {
	//fmt.Printf("before exec OperandStack length: %d, current Index: %d\n", len(self.slots), self.size)
	bits := math.Float32bits(val)
	self.slots[self.size].num = int32(bits)
	self.size++
	//fmt.Printf("after exec OperandStack length: %d, current Index: %d\n", len(self.slots), self.size)
}
func (self *OperandStack) PopFloat() float32 {
	//fmt.Printf("before exec OperandStack length: %d, current Index: %d\n", len(self.slots), self.size)
	self.size--
	bits := uint32(self.slots[self.size].num)
	//fmt.Printf("after exec OperandStack length: %d, current Index: %d\n", len(self.slots), self.size)
	return math.Float32frombits(bits)
}
func (self *OperandStack) PushLong(val int64)  {
	//fmt.Printf("before exec OperandStack length: %d, current Index: %d\n", len(self.slots), self.size)
	self.slots[self.size].num = int32(val)
	self.slots[self.size + 1].num = int32(val >> 32)
	self.size += 2
	//fmt.Printf("after exec OperandStack length: %d, current Index: %d\n", len(self.slots), self.size)
}
func (self *OperandStack) PopLong() int64  {
	//fmt.Printf("before exec OperandStack length: %d, current Index: %d\n", len(self.slots), self.size)
	self.size -= 2
	low := uint32(self.slots[self.size].num)
	high := uint32(self.slots[self.size + 1].num)
	//fmt.Printf("after exec OperandStack length: %d, current Index: %d\n", len(self.slots), self.size)
	return int64(high) << 32 | int64(low)
}
func (self *OperandStack) PushDouble(val float64)  {
	bits := math.Float64bits(val)
	self.PushLong(int64(bits))
}
func (self *OperandStack) PopDouble() float64  {
	bits := uint64(self.PopLong())
	return math.Float64frombits(bits)
}
func (self *OperandStack) PushRef(ref *heap.Object) {
	//fmt.Printf("before exec OperandStack length: %d, current Index: %d\n", len(self.slots), self.size)
	self.slots[self.size].ref = ref
	self.size++
	//fmt.Printf("after exec OperandStack length: %d, current Index: %d\n", len(self.slots), self.size)
}
func (self *OperandStack) PopRef() *heap.Object  {
	//fmt.Printf("before exec OperandStack length: %d, current Index: %d\n", len(self.slots), self.size)
	self.size--
	ref := self.slots[self.size].ref
	self.slots[self.size].ref = nil
	//fmt.Printf("after exec OperandStack length: %d, current Index: %d\n", len(self.slots), self.size)
	return ref
}
func (self *OperandStack) PushSlot(slot Slot)  {
	//fmt.Printf("before exec OperandStack length: %d, current Index: %d\n", len(self.slots), self.size)
	self.slots[self.size] = slot
	self.size++
	//fmt.Printf("after exec OperandStack length: %d, current Index: %d\n", len(self.slots), self.size)
}
func (self *OperandStack) PopSlot() Slot  {
	//fmt.Printf("before exec OperandStack length: %d, current Index: %d\n", len(self.slots), self.size)
	self.size--
	//fmt.Printf("after exec OperandStack length: %d, current Index: %d\n", len(self.slots), self.size)
	return self.slots[self.size]
}
func (self *OperandStack) GetRefFromTop(n uint) *heap.Object  {
	return self.slots[self.size - 1 - n].ref
}