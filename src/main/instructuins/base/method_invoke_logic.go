package base

import (
	"fmt"
	"github.com/kuangcp/simple-jvm/src/main/rtda"
	"github.com/kuangcp/simple-jvm/src/main/rtda/heap"
)

func InvokeMethod(invokeFrame *rtda.Frame, method *heap.Method) {
	thread := invokeFrame.Thread()
	newFrame := thread.NewFrame(method)
	thread.PushFrame(newFrame)

	argSlot := int(method.ArgSlotCount())
	if argSlot > 0 {
		for i := argSlot - 1; i >= 0; i-- {
			slot := invokeFrame.OperandStack().PopSlot()
			newFrame.LocalVars().SetSlot(uint(i), slot)
		}
	}

	// hack!
	if method.IsNative() {
		if method.Name() == "registerNatives" {
			thread.PopFrame()
		} else {
			panic(fmt.Sprintf("native method: %v.%v%v\n", method.Class().Name() ,method.Name(), method.Descriptor()))
		}
	}
}
