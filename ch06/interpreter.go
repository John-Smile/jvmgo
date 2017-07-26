package main

import "fmt"
import "jvmgo/ch06/classfile"
import "jvmgo/ch06/instructions"
import "jvmgo/ch06/instructions/base"
import "jvmgo/ch06/rtda"

func interpret(methodInfo *classfile.MemberInfo)  {
	codeAttr := methodInfo.CodeAttribute()
	maxLocals := codeAttr.MaxLocals()
	maxStack := codeAttr.MaxStack()
	bytecode := codeAttr.Code()
	thread := rtda.NewThread()
	frame := thread.NewFrame(maxLocals, maxStack)
	fmt.Printf("frame created pc: %d\n", frame.NextPC())
	thread.PushFrame(frame)
	defer catchErr(frame)
	loop(thread, bytecode)
}

func catchErr(frame *rtda.Frame)  {
	if r := recover(); r != nil {
		fmt.Printf("LocalVars:%v\n", frame.LocalVars())
		fmt.Printf("OperandStack:%v\n", frame.OperandStack())
		panic(r)
	}
}

func loop(thread *rtda.Thread, byteCode []byte)  {
	frame := thread.PopFrame()
	fmt.Printf("frame poped pc: %d\n", frame.NextPC())
	reader := &base.BytecodeReader{}
	for {
		pc := frame.NextPC()
		fmt.Printf("pc: %d\n", pc)
		thread.SetPC(pc)
		reader.Reset(byteCode, pc)
		opcode := reader.ReadUint8()
		inst := instructions.NewInstruction(opcode)
		inst.FetchOperands(reader)
		frame.SetNextPC(reader.PC())
		fmt.Printf("pc:%2d inst:%T %v\n", pc, inst, inst)
		inst.Execute(frame)
	}
}
