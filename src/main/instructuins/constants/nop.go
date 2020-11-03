package constants

import (
	"main/instructuins/base"
	"main/rtda"
)

type NOP struct {
	base.NoOperandsInstruction
}

func (self *NOP) Execute(frame *rtda.Frame) {
	// 不做操作
}
