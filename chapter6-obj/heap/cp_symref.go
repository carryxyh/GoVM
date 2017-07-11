package heap

/**
	类、字段、方法和接口符号的引用引用 都是常量
 */
type SymRef struct {
	cp        *ConstantPool
	className string
	class     *Class
}
