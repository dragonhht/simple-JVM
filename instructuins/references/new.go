package references

import (
	"github.com/kuangcp/simple-jvm/instructuins/base"
	"github.com/kuangcp/simple-jvm/rtda"
	"github.com/kuangcp/simple-jvm/rtda/heap"
)

/*
	new指令
 */

type NEW struct {
	base.Index16Instruction
}

func (self *NEW) Execute(frame *rtda.Frame) {
	cp := frame.Method().Class().ConstantPool()
	classRef := cp.GetConstant(self.Index).(*heap.ClassRef)
	class := classRef.ResolvedClass()
	if !class.InitStarted() {
		frame.RevertNextPC()
		base.InitClass(frame.Thread(), class)
		return
	}

	if class.IsInterface() || class.IsAbstract() {
		panic("java.lang.InstantiationError")
	}
	ref := class.NewObject()
	frame.OperandStack().PushRef(ref)
}
