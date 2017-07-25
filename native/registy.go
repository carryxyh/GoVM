package native

import "GoVM/chapter4-rtdt"

type NativeMethod func(frame *chapter4_rtdt.Frame)

var registry = map[string]NativeMethod{}

/**
	空实现
 */
func emptyNateveMethod(frame *chapter4_rtdt.Frame) {
}

func Register(className, methodName, methodDescriptor string, method NativeMethod) {
	key := className + "~" + methodName + "~" + methodDescriptor
	registry[key] = method
}

func FindNativeMethod(className, methodName, methodDescriptor string) NativeMethod {
	key := className + "~" + methodName + "~" + methodDescriptor
	if method, ok := registry[key]; ok {
		return method
	}

	if methodDescriptor == "()V" && methodName == "registerNatives" {
		return emptyNateveMethod
	}
	return nil
}
