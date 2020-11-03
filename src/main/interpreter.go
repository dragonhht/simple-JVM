package main

import (
	"fmt"
	"github.com/kuangcp/simple-jvm/src/main/instructuins"
	"github.com/kuangcp/simple-jvm/src/main/instructuins/base"
	"github.com/kuangcp/simple-jvm/src/main/rtda"
	"github.com/kuangcp/simple-jvm/src/main/rtda/heap"
)

/*
	简单的解释器
 */

func interpret(method *heap.Method, logInst bool) {

	thread := rtda.NewThread()
	frame := thread.NewFrame(method)
	thread.PushFrame(frame)

	defer catchErr(thread)
	loop(thread, logInst)
}

func catchErr(thread *rtda.Thread) {
	if r := recover(); r != nil {
		logFrames(thread)
		panic(r)
	}
}

func logFrames(thread *rtda.Thread) {
	for !thread.IsStackEmpty()  {
		frame := thread.PopFrame()
		method := frame.Method()
		className := method.Class().Name()
		fmt.Printf(">> pc: %4d %v.%v%v \n", frame.NextPC(), className, method.Name(), method.Descriptor())
	}
}

func loop(thread *rtda.Thread, logInst bool) {
	reader := &base.BytecodeReader{}
	for {
		frame := thread.CurrentFrame()
		pc := frame.NextPC()
		thread.SetPC(pc)

		reader.Reset(frame.Method().Code(), pc)
		opCode := reader.ReadUint8()
		inst := instructuins.NewInstruction(opCode)
		inst.FetchOperands(reader)
		frame.SetNextPC(reader.PC())
		if logInst {
			logInstruction(frame, inst)
		}
		inst.Execute(frame)
		if thread.IsStackEmpty() {
			break
		}
	}
}

func logInstruction(frame *rtda.Frame, instruction base.Instruction) {
	method := frame.Method()
	className := method.Class().Name()
	methodName := method.Name()
	pc := frame.Thread().PC()
	fmt.Printf("%v.%v() #%2d %T %v\n", className, methodName, pc, instruction, instruction)
}