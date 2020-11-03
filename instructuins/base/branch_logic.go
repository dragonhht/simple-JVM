package base

import "simple-jvm/rtda"

/*
	跳转逻辑
 */
func Branch(frame *rtda.Frame, offset int)  {
	pc := frame.Thread().PC()
	nextPC := pc + offset
	frame.SetNextPC(nextPC)
}
