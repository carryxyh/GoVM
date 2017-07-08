package base

import "GoVM/chapter4-rtdt"

func Branch(frame *chapter4_rtdt.Frame, offset int) {
	pc := frame.Thread().pc()
	nextPC := pc + offset
	frame.setNextPC(nextPC)
}
