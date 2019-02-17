package rtda

type Frame struct {
	lower *Frame
	localVars LocalVars
	operandStack *OperandStack
}

func (self Frame) LocalVars() LocalVars {
	return self.localVars
}

func (self Frame) Lower() *Frame {
	return self.lower
}

func (self Frame) OperandStack() *OperandStack {
	return self.operandStack
}

func newFrame(maxLocals, maxStack uint) *Frame {
	return &Frame{
		localVars: newLocalVars(maxLocals),
		operandStack: newOperandStack(maxStack),
	}
}
