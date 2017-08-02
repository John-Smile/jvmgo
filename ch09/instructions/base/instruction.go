package base

import "jvmgo/ch09/rtda"

type Instruction interface {
	FetchOperands(reader *BytecodeReader)
	Execute(frame *rtda.Frame)
}

type NoOperandsInstruction struct {
}

func (self *NoOperandsInstruction) FetchOperands(reader *BytecodeReader)  {
	// do nothing
}

type BranchInstruction struct {
	Offset            int
}

func (self *BranchInstruction) FetchOperands(reader *BytecodeReader)  {
	self.Offset = int(reader.ReadInt16())
}

type Index8Instruction struct {
	Index              uint
}

func (self *Index8Instruction) FetchOperands(reader *BytecodeReader)  {
	self.Index = uint(reader.ReadUint8())
}

type Index16Instruction struct {
	Index       uint16
}

func (self *Index16Instruction) FetchOperands(reader *BytecodeReader)  {
	self.Index = reader.ReadUint16()
}