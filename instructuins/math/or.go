package math

/*
	按位或
 */

import (
	"github.com/kuangcp/simple-jvm/instructuins/base"
	"github.com/kuangcp/simple-jvm/rtda"
)

type IOR struct {
	base.NoOperandsInstruction
}

func (self *IOR) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	result := v1 | v2
	stack.PushInt(result)
}

type LOR struct {
	base.NoOperandsInstruction
}

func (self *LOR) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	result := v1 | v2
	stack.PushLong(result)
}