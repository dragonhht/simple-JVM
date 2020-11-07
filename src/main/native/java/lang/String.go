package lang

import (
	"main/native"
	"main/rtda"
	"main/rtda/heap"
)

func init() {
	native.Register("java/lang/String", "intern",
		"()Ljava/lang/String;", intern)
}

func intern(frame *rtda.Frame) {
	this := frame.LocalVars().GetThis()
	interned := heap.InternString(this)
	frame.OperandStack().PushRef(interned)
}
