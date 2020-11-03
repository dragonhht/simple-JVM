package rtda

import "github.com/kuangcp/simple-jvm/src/main/rtda/heap"

type Frame struct {
	lower *Frame
	localVars LocalVars
	operandStack *OperandStack
	thread *Thread
	nextPC int
	method *heap.Method
}

func (self *Frame) LocalVars() LocalVars {
	return self.localVars
}

func (self *Frame) Method() *heap.Method {
	return self.method
}

func (self *Frame) Lower() *Frame {
	return self.lower
}

func (self *Frame) Thread() *Thread {
	return self.thread
}

func (self *Frame) NextPC() int {
	return self.nextPC
}

func (self *Frame) SetNextPC(pc int)  {
	self.nextPC = pc
}

func (self *Frame) OperandStack() *OperandStack {
	return self.operandStack
}

func (self *Frame) RevertNextPC() {
	self.nextPC = self.thread.pc
}

func newFrame(thread *Thread, method *heap.Method) *Frame {
	return &Frame{
		thread: thread,
		method: method,
		localVars: newLocalVars(method.MaxLocals()),
		operandStack: newOperandStack(method.MaxStack()),
	}
}
