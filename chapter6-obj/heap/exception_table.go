package heap

import "GoVM/chapter3-cf/classfile"

type ExceptionTable []*ExceptionHandler

/**
	异常处理器
 */
type ExceptionHandler struct {
	startPc   int
	endPc     int
	handlerPc int
	catchType *ClassRef
}

func newExceptionTable(entries []*chapter3_cf.ExceptionTableEntry, cp *ConstantPool) ExceptionTable {
	table := make([]*ExceptionHandler, len(entries))
	for i, entry := range entries {
		table[i] = &ExceptionHandler{
			startPc: int(entry.StartPc()),
			endPc: int(entry.EndPc()),
			handlerPc: int(entry.HandlerPc()),
			catchType: getCatchType(uint(entry.CatchType()), cp),
		}
	}
	return table
}

/*
	如果常量池索引为0，在这里表示catch-all
 */
func getCatchType(index uint, cp *ConstantPool) *ClassRef {
	if index == 0 {
		return nil
	}
	return cp.GetConstant(index).(*ClassRef)
}

/**
	从异常表中找到异常处理器
 */
func (self ExceptionTable) findExceptionHandler(exClass *Class, pc int) *ExceptionHandler {
	for _, handler := range self {
		if pc >= handler.startPc && pc <= handler.endPc {
			if handler.catchType == nil {
				return handler
			}
			catchClass := handler.catchType.ResolvedClass()
			if catchClass == exClass || catchClass.IsSuperClassOf(exClass) {
				return handler
			}
		}
	}
	return nil
}