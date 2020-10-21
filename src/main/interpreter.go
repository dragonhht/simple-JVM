package main

import (
	"fmt"
	"main/instructuins"
	"main/instructuins/base"
	"main/rtda"
	"main/rtda/heap"
)

/*
	简单的解释器
 */

func interpret(method *heap.Method) {

	thread := rtda.NewThread()
	frame := thread.NewFrame(method)
	thread.PushFrame(frame)

	defer catchErr(frame)
	loop(thread, method.Code())
}

func catchErr(frame *rtda.Frame) {
	if r := recover(); r != nil {
		fmt.Printf("LocalVars:%v\n", frame.LocalVars())
		fmt.Printf("OperandStack:%v\n", frame.OperandStack())
		panic(r)
	}
}

func loop(thread *rtda.Thread, byteCode []byte) {
	frame := thread.PopFrame()
	reader := &base.BytecodeReader{}
	for {
		pc := frame.NextPC()
		thread.SetPC(pc)

		reader.Reset(byteCode, pc)
		opCode := reader.ReadUint8()
		inst := instructuins.NewInstruction(opCode)
		inst.FetchOperands(reader)
		frame.SetNextPC(reader.PC())

		fmt.Printf("pc:%2d inst:%T %v\n", pc, inst, inst)
		inst.Execute(frame)
	}
}