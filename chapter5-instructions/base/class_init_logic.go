package base

import (
	"GoVM/chapter4-rtdt"
	"GoVM/chapter6-obj/heap"
)

func InitClass(thread chapter4_rtdt.Thread, class *heap.Class) {
	class.StartInit()
	scheduleClinit(thread, class)
	initSuperClass(thread, class)
}

/**
	调用clinit
 */
func scheduleClinit(thread chapter4_rtdt.Thread, class *heap.Class) {
	clinit := class.GetClinitMethod()
	if clinit != nil {
		newFrame := thread.NewFrame(clinit)
		thread.PushFrame(newFrame)
	}
}

/**
	递归的初始化父类
 */
func initSuperClass(thread chapter4_rtdt.Thread, class *heap.Class) {
	if !class.IsInterface() {
		superClass := class.SuperClass()
		if superClass != nil && !superClass.InitStarted() {
			InitClass(thread, superClass)
		}
	}
}