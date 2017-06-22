package chapter3_cf

import "fmt"

type ClassFile struct {
	//魔数
	//magic uint32
	minorVersion uint32
	majorVersion uint16
	//constantPool ConstantPool
	accessFlags  uint16
	thisClass    uint16
	superClass   uint16
	interfaces   []uint16
	//fields       []*MemberInfo
	//methods      []*MemberInfo
	//attributes   []AttributeInfo
}

func Parse(classData []byte) (cf *ClassFile, err error) {
	defer func() {
		if r := recover(); r != nil {
			var ok bool
			err, ok = r.(error)
			if !ok {
				err = fmt.Errorf("%v", r)
			}
		}
	}()
	cr := &ClassReader{classData}
	cf = &ClassFile{}
	cf.read(cr)
	return
}

func (this *ClassFile) read(reader *ClassReader) {

}

func (this *ClassFile) readAndCheckMagic(reader *ClassReader) {

}

func (this *ClassFile) readAndCheckVersion(reader *ClassReader) {

}

func (this *ClassFile) MinorVersion() uint16 {
	return 0
}

func (this *ClassFile) MajorVersion() uint16 {
	return 0
}

//func (this *ClassFile) ConstantPool() ConstantPool {
//	return nil
//}
//
//func (this *ClassFile) AccessFlags() uint16 {
//
//}
//
//func (this *ClassFile) Fields() []*MemberInfo {
//	return nil
//}
//
//func (this *ClassFile) Methods() []*MemberInfo {
//	return nil
//}

func (this *ClassFile) ClassName() string {
	return ""
}

func (this *ClassFile) SuperClassName() string {
	return ""
}

func (this *ClassFile) InterfaceNames() []string {
	return ""
}