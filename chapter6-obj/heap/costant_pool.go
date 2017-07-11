package heap

import (
	"GoVM/chapter3-cf/classfile"
	"fmt"
)

//能存放任何类型
type Constant interface {

}

type ConstantPool struct {
	class  *Class
	consts []Constant
}

/**
	把classFile中的常量池转化成运行时常量池
 */
func newConstantPool(class *Class, cfConstantPool chapter3_cf.ConstantPool) *ConstantPool {
	cpCount := len(cfConstantPool)
	consts := make([]Constant, cpCount)
	rtConstantPool := &ConstantPool{class, consts}

	for i := 1; i < cpCount; i++ {
		cpInfo := cfConstantPool[i]
		switch cpInfo.(type) {

		case chapter3_cf.ConstantIntegerInfo:
			intInfo := cpInfo.(chapter3_cf.ConstantIntegerInfo)
			consts[i] = intInfo.Value()
		case *chapter3_cf.ConstantFloatInfo:
			floatInfo := cpInfo.(*chapter3_cf.ConstantFloatInfo)
			consts[i] = floatInfo.Value()
		case *chapter3_cf.ConstantLongInfo:
			longInfo := cpInfo.(*chapter3_cf.ConstantLongInfo)
			consts[i] = longInfo.Value()
			i++ // -> long double 在常量在常量池中都是占据着两个位置
		case *chapter3_cf.ConstantDoubleInfo:
			doubleInfo := cpInfo.(*chapter3_cf.ConstantDoubleInfo)
			consts[i] = doubleInfo.Value()
		case *chapter3_cf.ConstantStringInfo:
			stringInfo := cpInfo.(*chapter3_cf.ConstantStringInfo)
			consts[i] = stringInfo.String()
		case *chapter3_cf.ConstantClassInfo:
			classInfo := cpInfo.(*chapter3_cf.ConstantClassInfo)
			consts[i] = newClassRef(rtConstantPool, classInfo)
		case *chapter3_cf.ConstantFieldrefInfo:
			fieldrefInfo := cpInfo.(*chapter3_cf.ConstantFieldrefInfo)
			consts[i] = newFieldRef(rtConstantPool, fieldrefInfo)
		case *chapter3_cf.ConstantMethodrefInfo:
			methodrefInfo := cpInfo.(*chapter3_cf.ConstantMethodrefInfo)
			consts[i] = newMethodRef(rtConstantPool, methodrefInfo)
		case *chapter3_cf.ConstantInterfaceMethodrefInfo:
			methodrefInfo := cpInfo.(*chapter3_cf.ConstantInterfaceMethodrefInfo)
			consts[i] = newInterfaceMethodRef(rtConstantPool, methodrefInfo)
		default:
		// todo
		}
	}
	return rtConstantPool
}

/**
	根据索引返回常量
 */
func (self *ConstantPool) GetConstant(index uint) Constant {
	if c := self.consts[index]; c != nil {
		return c
	}
	panic(fmt.Sprintf("No constants at index %d", index))
}