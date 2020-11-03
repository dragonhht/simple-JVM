package base

import "github.com/kuangcp/simple-jvm/rtda"

/*
	解释器结构
 */
type Instruction interface {
	// 读取操作数
	FetchOperands(reader *BytecodeReader)
	// 执行逻辑
	Execute(frame *rtda.Frame)
}

/*
	没有操作数的指令
 */
type NoOperandsInstruction struct {}

func (self *NoOperandsInstruction) FetchOperands(reader *BytecodeReader) {
	// 没有指定，所以不做操作
}

/*
	跳转指令.
 */
type BranchInstruction struct {
	Offset int
}

func (self *BranchInstruction) FetchOperands(reader *BytecodeReader) {
	self.Offset = int(reader.ReadInt16())
}

/*
	存储和加载指令
 */
type Index8Instruction struct {
	Index uint
}

func (self *Index8Instruction) FetchOperands(reader *BytecodeReader) {
	self.Index = uint(reader.ReadUint8())
}

/*
	需访问运行时常量池的指令
 */
type Index16Instruction struct {
	Index uint
}

func (self *Index16Instruction) FetchOperands(reader *BytecodeReader) {
	self.Index = uint(reader.ReadUint16())
}