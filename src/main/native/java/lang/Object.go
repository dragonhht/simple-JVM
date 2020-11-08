package lang

import (
	"main/native"
	"main/rtda"
	"unsafe"
)

func init() {
	// 注册本地方法 getClass
	native.Register("java/lang/Object", "getClass",
		"()Ljava/lang/Class;", getClass)
	// 注册hashCode方法
	native.Register("java/lang/Object", "hashCode",
		"()I", hashCode)
	// 注册clone方法
	native.Register("java/lang/Object", "clone",
		"()Ljava/lang/Object;", clone)
}

func clone(frame *rtda.Frame) {
	this := frame.LocalVars().GetThis()
	cloneable := this.Class().Loader().LoadClass("java/lang/Cloneable")
	if !this.Class().IsImplements(cloneable) {
		panic("java.lang.CloneNotSupportedException")
	}
	frame.OperandStack().PushRef(this.Clone())
}

func hashCode(frame *rtda.Frame) {
	this := frame.LocalVars().GetThis()
	hash := int32(uintptr(unsafe.Pointer(this)))
	frame.OperandStack().PushInt(hash)
}

/*
	实现本地方法 getClass
*/
func getClass(frame *rtda.Frame) {
	this := frame.LocalVars().GetThis()
	class := this.Class().JClass()
	frame.OperandStack().PushRef(class)
}