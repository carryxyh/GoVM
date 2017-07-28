package lang

import (
	"GoVM/native"
	"GoVM/chapter4-rtdt"
	"math"
)

const jlFloat = "java/lang/Float"

func init() {
	native.Register(jlFloat, "floatToRawIntBits", "(F)I", floatToRawIntBits)
	native.Register(jlFloat, "intBitsToFloat", "(I)F", intBitsToFloat)
}

// public static native int floatToRawIntBits(float value);
// (F)I
func floatToRawIntBits(frame *chapter4_rtdt.Frame) {
	value := frame.LocalVars().GetFloat(0)
	bits := math.Float32bits(value)
	frame.OperandStack().PushInt(int32(bits))
}

// public static native float intBitsToFloat(int bits);
// (I)F
func intBitsToFloat(frame *chapter4_rtdt.Frame) {
	bits := frame.LocalVars().GetInt(0)
	value := math.Float32frombits(uint32(bits))
	frame.OperandStack().PushFloat(value)
}
