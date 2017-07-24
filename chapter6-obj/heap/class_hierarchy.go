package heap

/**
	一个类是否是可以看成是另外一个
 */
func (self *Class) IsAssignableFrom(other *Class) bool {
	//s -> other  t -> self
	s, t := other, self
	if s == t {
		return true
	}

	if !s.IsArray() {
		if !s.IsInterface() {
			// s is class
			if !t.IsInterface() {
				// t is not interface
				return s.IsSubClassOf(t)
			} else {
				// t is interface
				return s.IsImplements(t)
			}
		} else {
			// s is interface
			if !t.IsInterface() {
				// t is not interface
				return t.isJlObject()
			} else {
				// t is interface
				return t.isSuperInterfaceOf(s)
			}
		}
	} else {
		// s is array
		if !t.IsArray() {
			if !t.IsInterface() {
				// t is class
				return t.isJlObject()
			} else {
				// t is interface
				return t.isJlCloneable() || t.isJioSerializable()
			}
		} else {
			// t is array
			sc := s.ComponentClass()
			tc := t.ComponentClass()
			return sc == tc || tc.IsAssignableFrom(sc)
		}
	}

	return false
}

/**
	一个类是不是另外一个类的子类
 */
func (self *Class) isSubClassOf(other *Class) bool {
	for c := self.superClass; c != nil; c = c.superClass {
		if c == other {
			return true
		}
	}
	return false
}

// self implements iface
func (self *Class) IsImplements(iface *Class) bool {
	for c := self; c != nil; c = c.superClass {
		for _, i := range c.interfaces {
			if i == iface || i.isSubInterfaceOf(iface) {
				return true
			}
		}
	}
	return false
}

/**
	一个类是不是一个接口的实现类
 */
func (self *Class) isImplement(inface *Class) bool {
	for c := self; c != nil; c = c.superClass {
		for _, i := range c.interfaces {
			if i == inface || i.isSubInterfaceOf(inface) {
				return true
			}
		}
	}
	return false
}

// iface extends self
func (self *Class) isSuperInterfaceOf(iface *Class) bool {
	return iface.isSubInterfaceOf(self)
}

//一个接口是不是另一个接口的子类
func (self *Class) isSubInterfaceOf(inface *Class) bool {
	for _, superInterface := range self.interfaces {
		if superInterface == inface || superInterface.isSubInterfaceOf(inface) {
			return true
		}
	}
	return false
}