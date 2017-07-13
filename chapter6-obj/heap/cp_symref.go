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

func (self *SymRef) ResolvedClass() *Class {
	if self.class == nil {
		self.resolveClassRef()
	}
	return self.class
}

/**
	   通过符号引用N
	类D -------------> 类C
	解析过程：
		D的加载器加载C
		D是否有权限访问C -> 没有抛异常

	如果类D通过符号引用N引用类C的话，要解析N，先用D的类加载器加载C，然后检查D是否有权限访问C，如果没有抛出异常
 */
func (self *SymRef) resolveClassRef() {
	d := self.cp.class
	c := d.loader.LoadClass(self.className)
	if !c.isAccessibleTo(d) {
		panic("java.lang.IllegalAccessError")
	}
	self.class = c
}