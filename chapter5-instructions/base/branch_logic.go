package base

import "GoVM/chapter4-rtdt"

/**
	跳转逻辑的基类
 */
func Branch(frame *chapter4_rtdt.Frame, offset int) {
	pc := frame.Thread().PC()
	nextPC := pc + offset
	frame.SetNextPC(nextPC)
}
