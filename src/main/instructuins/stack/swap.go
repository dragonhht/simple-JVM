package stack

import (
	"main/instructuins/base"
	"main/rtda"
)

type SWAP struct {
	base.NoOperandsInstruction
}

/*
	交换栈顶两个变量
*/
func (self *SWAP) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	slot1 := stack.PopSlot()
	slot2 := stack.PopSlot()
	stack.PushSlot(slot1)
	stack.PushSlot(slot2)
}
