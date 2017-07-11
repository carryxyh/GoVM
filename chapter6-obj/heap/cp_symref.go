package heap

/**
	类、字段、方法和接口符号的引用引用 都是常量
	运行时常量池主要存放两类信息：字面量 和 符号引用
	字面量包括：整数、浮点数和字符串字面量
	符号引用包括：类符号引用、字段符号引用、方法符号引用和接口符号引用
 */
type SymRef struct {
	cp        *ConstantPool
	className string
	class     *Class
}
