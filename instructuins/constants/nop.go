package constants

import (
	"simple-jvm/instructuins/base"
	"simple-jvm/rtda"
)

type NOP struct {
	base.NoOperandsInstruction
}

func (self *NOP) Execute(frame *rtda.Frame) {
	// 不做操作
}
