package heap

/**
	存放字段和方法符号引用公有的信息
 */
type MemberRef struct {
	SymRef
	name       string
	descriptor string
}
