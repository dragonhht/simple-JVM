package references

import (
	"main/instructuins/base"
	"main/rtda"
	"main/rtda/heap"
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
