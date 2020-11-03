package control

import (
	"github.com/kuangcp/simple-jvm/src/main/instructuins/base"
	"github.com/kuangcp/simple-jvm/src/main/rtda"
)

type RETURN struct {
	base.NoOperandsInstruction
}

func (self *RETURN) Execute(frame *rtda.Frame) {
	frame.Thread().PopFrame()
}

type ARETURN struct {
	base.NoOperandsInstruction
}

func (self *ARETURN) Execute(frame *rtda.Frame) {
	thread := frame.Thread()
	currentFrame := thread.PopFrame()
	invokeFrame := thread.TopFrame()
	retVal := currentFrame.OperandStack().PopRef()
	invokeFrame.OperandStack().PushRef(retVal)
}

type DRETURN struct {
	base.NoOperandsInstruction
}

func (self *DRETURN) Execute(frame *rtda.Frame) {
	thread := frame.Thread()
	currentFrame := thread.PopFrame()
	invokeFrame := thread.TopFrame()
	retVal := currentFrame.OperandStack().PopDouble()
	invokeFrame.OperandStack().PushDouble(retVal)
}

type FRETURN struct {
	base.NoOperandsInstruction
}

func (self *FRETURN) Execute(frame *rtda.Frame) {
	thread := frame.Thread()
	currentFrame := thread.PopFrame()
	invokeFrame := thread.TopFrame()
	retVal := currentFrame.OperandStack().PopFloat()
	invokeFrame.OperandStack().PushFloat(retVal)
}

type IRETURN struct {
	base.NoOperandsInstruction
}

func (self *IRETURN) Execute(frame *rtda.Frame) {
	thread := frame.Thread()
	currentFrame := thread.PopFrame()
	invokeFrame := thread.TopFrame()
	retVal := currentFrame.OperandStack().PopInt()
	invokeFrame.OperandStack().PushInt(retVal)
}

type LRETURN struct {
	base.NoOperandsInstruction
}

func (self *LRETURN) Execute(frame *rtda.Frame) {
	thread := frame.Thread()
	currentFrame := thread.PopFrame()
	invokeFrame := thread.TopFrame()
	retVal := currentFrame.OperandStack().PopLong()
	invokeFrame.OperandStack().PushLong(retVal)
}