package constants

import (
	"github.com/kuangcp/simple-jvm/instructuins/base"
	"github.com/kuangcp/simple-jvm/rtda"
)

type NOP struct {
	base.NoOperandsInstruction
}

func (self *NOP) Execute(frame *rtda.Frame) {
	// 不做操作
}
