package main

import "fmt"
import "jvmgo/ch11/instructions"
import "jvmgo/ch11/instructions/base"
import (
	"jvmgo/ch11/rtda"
	"jvmgo/ch11/rtda/heap"
)

func interpret(thread *rtda.Thread, logInst bool)  {
	//thread := rtda.NewThread()
	//frame := thread.NewFrame(method)
	//thread.PushFrame(frame)
	//jArgs := createArgsArray(method.Class().Loader(), args)
	//frame.LocalVars().SetRef(0, jArgs)
	defer catchErr(thread)
	loop(thread, logInst)
}

func createArgsArray(loader *heap.ClassLoader, args [] string) *heap.Object  {
	stringClass := loader.LoadClass("java/lang/String")
	argsArr := stringClass.ArrayClass().NewArray(uint(len(args)))
	jArgs := argsArr.Refs()
	for i, arg := range args {
		jArgs[i] = heap.JString(loader, arg)
	}
	return argsArr
}

func catchErr(thread *rtda.Thread)  {
	if r := recover(); r != nil {
		logFrames(thread)
		//fmt.Printf("LocalVars:%v\n", frame.LocalVars())
		//fmt.Printf("OperandStack:%v\n", frame.OperandStack())
		panic(r)
	}
}
func logFrames(thread *rtda.Thread)  {
	for !thread.IsStackEmpty() {
		frame := thread.PopFrame()
		method := frame.Method()
		className := method.Class().Name()
		fmt.Printf(">> pc:%4d %v.%v%v \n", frame.NextPC(), className, method.Name(), method.Descriptor())
	}
}

func loop(thread *rtda.Thread, logInst bool)  {
	reader := &base.BytecodeReader{}
	for {
		frame := thread.CurrnetFrame()
		pc := frame.NextPC()
		thread.SetPC(pc)
		reader.Reset(frame.Method().Code(), pc)
		opcode := reader.ReadUint8()
		inst := instructions.NewInstruction(opcode)
		inst.FetchOperands(reader)
		frame.SetNextPC(reader.PC())
		//if (logInst) {
			logInstruction(frame, inst)
		//}
		inst.Execute(frame)
		if thread.IsStackEmpty() {
			break
		}
	}
}
func logInstruction(frame *rtda.Frame, inst base.Instruction)  {
	method := frame.Method()
	className := method.Class().Name()
	methodName := method.Name()
	pc := frame.Thread().PC()
	fmt.Printf("%v.%v() #%2d %T %v\n", className, methodName, pc, inst, inst)
}
