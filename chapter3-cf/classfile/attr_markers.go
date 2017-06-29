package chapter3_cf

/**
	用于支持@Deprecated注解
 */
type DeprecatedAttribute struct {
	MarkerAttribute
}

/**
	用来标记源文件中不存在的、由编译器生成的类成员，主要为了支持嵌套类（内部类）和嵌套接口
 */
type SyntheticAttribute struct {
	MarkerAttribute
}

/**
	上面两个struct的父类，其中没有任何数据
 */
type MarkerAttribute struct {
}

func (self *MarkerAttribute) readInfo(reader *ClassReader) {
	//read nothing
}