package heap

func (self *Class) IsAssignableFrom(other *Class) bool {
	//s -> other  t -> self
	s, t := other, self
	if s == t {
		return true
	}

	if !t.IsInterface() {
		return s.isSubClassOf(t)
	} else {
		return s.isImplement(t)
	}
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

func (self *Class) isSubInterfaceOf(inface *Class) bool {
	for _, superInterface := range self.interfaces {
		if superInterface == inface || superInterface.isSubInterfaceOf(inface) {
			return true
		}
	}
	return false
}