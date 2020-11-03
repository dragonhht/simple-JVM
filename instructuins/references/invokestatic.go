package references

import (
	"github.com/kuangcp/simple-jvm/instructuins/base"
	"github.com/kuangcp/simple-jvm/rtda"
	"github.com/kuangcp/simple-jvm/rtda/heap"
)

type INVOKE_STATIC struct {
	base.Index16Instruction
}

func (self *INVOKE_STATIC) Execute(frame *rtda.Frame) {
	cp := frame.Method().Class().ConstantPool()
	methodRef := cp.GetConstant(self.Index).(*heap.MethodRef)
	method := methodRef.ResolvedMethod()
	class := method.Class()
	if !class.InitStarted() {
		frame.RevertNextPC()
		base.InitClass(frame.Thread(), class)
		return
	}
	if !method.IsStatic() {
		panic("java.lang.IncompatibleClassChangeError")
	}
	base.InvokeMethod(frame, method)
}
