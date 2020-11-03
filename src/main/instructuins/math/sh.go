package math

/*
	位移指令
 */

import (
	"github.com/kuangcp/simple-jvm/src/main/instructuins/base"
	"github.com/kuangcp/simple-jvm/src/main/rtda"
)

/*
	int左位移
 */
 type ISHL struct {
 	base.NoOperandsInstruction
 }

func (self *ISHL) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	s := uint32(v2) & 0x1f
	result := v1 << s
	stack.PushInt(result)
}

 /*
 	int算术右位移
  */
type ISHR struct {
	base.NoOperandsInstruction
}

func (self *ISHR) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	s := uint32(v2) & 0x1f
	result := v1 >> s
	stack.PushInt(result)
}

/*
	int逻辑右位移
 */
type IUSHR struct {
	base.NoOperandsInstruction
}

func (self *IUSHR) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	s := uint32(v2) & 0x1f
	result := int32(uint32(v1) >> s)
	stack.PushInt(result)
}

/*
	long左位移
 */
type LSHL struct {
	base.NoOperandsInstruction
}

func (self *LSHL) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopLong()
	s := uint32(v2) & 0x3f
	result := v1 << s
	stack.PushLong(result)
}

/*
	long算术右位移
 */
type LSHR struct {
	base.NoOperandsInstruction
}

func (self *LSHR) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopLong()
	s := uint32(v2) & 0x3f
	result := v1 >> s
	stack.PushLong(result)
}

/*
	long逻辑右位移
 */
type LUSHR struct {
	base.NoOperandsInstruction
}

func (self *LUSHR) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopLong()
	s := uint32(v2) & 0x3f
	result := int64(uint64(v1) >> s)
	stack.PushLong(result)
}