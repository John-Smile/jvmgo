package lang

import "jvmgo/ch11/native"
import "jvmgo/ch11/rtda"
import (
	"jvmgo/ch11/rtda/heap"
	"fmt"
)

func init()  {
	native.Register("java/lang/Throwable", "fillInStackTrace", "(I)Ljava/lang/Throwable;", fillInStackTrace)
}
func fillInStackTrace(frame *rtda.Frame)  {
	this := frame.LocalVars().GetThis()
	frame.OperandStack().PushRef(this)

	stes := createStackTraceElements(this, frame.Thread())
	this.SetExtra(stes)
}

type StackTraceElement struct {
	fileName            string
	className           string
	methodName          string
	lineNumber          int
}

func (self *StackTraceElement) String() string  {
	return fmt.Sprintf("%s.%s(%s:%d)",
	    self.className, self.methodName, self.fileName, self.lineNumber)
}

func createStackTraceElements(tObj *heap.Object, thread *rtda.Thread) [] *StackTraceElement {
	skip := distanceToObject(tObj.Class()) + 2
	frames := thread.GetFrames()[skip:]
	stes := make([] *StackTraceElement, len(frames))
	for i, frame := range frames {
		stes[i] = createStackTraceElement(frame)
	}
	return stes
}
func createStackTraceElement(frame *rtda.Frame) *StackTraceElement {
	method := frame.Method()
	class := method.Class()
	return &StackTraceElement {
		fileName:              class.SourceFile(),
		className:             class.JavaName(),
		methodName:            method.Name(),
		lineNumber:            method.GetLineNumber(frame.NextPC() - 1),
	}

}
func distanceToObject(class *heap.Class) int {
	distance := 0
	for c := class.SuperClass(); c != nil; c = c.SuperClass() {
		distance++
	}
	return distance
}