package chapter3_cf

/**
	字段/方法
 */
type MemberInfo struct {
	//常量池
	constantPool    ConstantPool
	accessFlags     uint16
	nameIndex       uint16
	//描述符
	descriptorIndex uint16
	//属性表
	attributes      []AttributeInfo
}

/**
	从class中获取字段/方法
 */
func readMembers(reader *ClassReader, cp ConstantPool) []*MemberInfo {
	memberCount := reader.readUint16()
	members := make([]*MemberInfo, memberCount)
	for i := range members {
		members[i] = readMember(reader, cp)
	}
	return members
}

func readMember(reader *ClassReader, cp ConstantPool) *MemberInfo {
	return &MemberInfo{
		constantPool: cp,
		accessFlags: reader.readUint16(),
		nameIndex: reader.readUint16(),
		descriptorIndex: reader.readUint16(),
		attributes: readAttributes(reader, cp),
	}
}

func (this *MemberInfo) AccessFlags() uint16 {
	return this.accessFlags
}

func (this *MemberInfo) Name() string {
	return this.constantPool.getUtf8(this.nameIndex)
}

func (this *MemberInfo) Descriptor() string {
	return this.constantPool.getUtf8(this.descriptorIndex)
}

func (this *MemberInfo) CodeAttribute() *CodeAttribute {
	for _, attrInfo := range this.attributes {
		switch attrInfo.(type) {
		case *CodeAttribute:
			return attrInfo.(*CodeAttribute)
		}
	}
	return nil
}

func (this *MemberInfo) ConstantValueAttribute() *ConstantValueAttribute {
	for _, attrInfo := range this.attributes {
		switch attrInfo.(type) {
		case *ConstantValueAttribute :
			return attrInfo.(*ConstantValueAttribute)
		}
	}
	return nil
}