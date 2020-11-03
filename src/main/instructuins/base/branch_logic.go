package base

import "github.com/kuangcp/simple-jvm/src/main/rtda"

/*
	跳转逻辑
 */
func Branch(frame *rtda.Frame, offset int)  {
	pc := frame.Thread().PC()
	nextPC := pc + offset
	frame.SetNextPC(nextPC)
}
