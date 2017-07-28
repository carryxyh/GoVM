package lang

import (
	"GoVM/native"
	"GoVM/chapter4-rtdt"
	"math"
)

const jlDouble = "java/lang/Double"

func init() {
	native.Register(jlDouble, "doubleToRawLongBits", "(D)J", doubleToRawLongBits)
	native.Register(jlDouble, "longBitsToDouble", "(J)D", longBitsToDouble)
}

// public static native long doubleToRawLongBits(double value);
// (D)J
func doubleToRawLongBits(frame *chapter4_rtdt.Frame) {
	value := frame.LocalVars().GetDouble(0)
	bits := math.Float64bits(value)
	frame.OperandStack().PushLong(int64(bits))
}

// public static native double longBitsToDouble(long bits);
// (J)D
func longBitsToDouble(frame *chapter4_rtdt.Frame) {
	bits := frame.LocalVars().GetLong(0)
	value := math.Float64frombits(uint64(bits))
	frame.OperandStack().PushDouble(value)
}