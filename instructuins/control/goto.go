package control

import (
	"github.com/kuangcp/simple-jvm/instructuins/base"
	"github.com/kuangcp/simple-jvm/rtda"
)

type GOTO struct {
	base.BranchInstruction
}

func (self *GOTO) Execute(frame *rtda.Frame) {
	base.Branch(frame, self.Offset)
}
