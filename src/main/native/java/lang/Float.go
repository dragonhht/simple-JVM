package lang

import (
	"main/native"
	"main/rtda"
	"math"
)

func init() {
	native.Register("java/lang/Float", "floatToRawIntBits",
		"(F)I", floatToRawIntBits)
	native.Register("java/lang/Float", "intBitsToFloat",
		"(I)F", intBitsToFloat)
}

func floatToRawIntBits(frame *rtda.Frame) {
	value := frame.LocalVars().GetFloat(0)
	bits := math.Float32bits(value)
	frame.OperandStack().PushInt(int32(bits))
}

func intBitsToFloat(frame *rtda.Frame) {
	bits := frame.LocalVars().GetInt(0)
	value := math.Float32frombits(uint32(bits)) // todo
	frame.OperandStack().PushFloat(value)
}
