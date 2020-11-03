package comparisons

import (
	"simple-jvm/instructuins/base"
	"simple-jvm/rtda"
)

/*
	double类型变量比较
*/
type DCMPG struct {
	base.NoOperandsInstruction
}

func (self *DCMPG) Execute(frame *rtda.Frame) {
	_dcmp(frame, true)
}

type DCMPL struct {
	base.NoOperandsInstruction
}

func (self *DCMPL) Execute(frame *rtda.Frame) {
	_dcmp(frame, true)
}

func _dcmp(frame *rtda.Frame, gFlag bool) {
	stack := frame.OperandStack()
	v2 := stack.PopDouble()
	v1 := stack.PopDouble()
	if v1 > v2 {
		stack.PushInt(1)
	} else if v1 == v2 {
		stack.PushInt(0)
	} else if v1 < v2 {
		stack.PushInt(-1)
	} else if gFlag {
		// 当存在NaN时，比较结果为1
		stack.PushInt(1)
	} else {
		// 当存在NaN时，比较结果为-1
		stack.PushInt(-1)
	}
}