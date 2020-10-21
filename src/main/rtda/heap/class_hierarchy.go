package heap

func (self *Class) isAssignableFrom(class *Class) bool {
	s ,t := class, self
	if s == t {
		return true
	}
	if !t.IsInterface() {
		return s.isSubClassOf(t)
	} else {
		return s.isImplements(t)
	}
}

func (self *Class) isImplements(class *Class) bool {
	for c := self; c != nil; c = c.superClass {
		for _, i := range c.interfaces {
			if i == class || i.isSubInterfaceOf(class) {
				return true
			}
		}
	}
	return false
}

func (self *Class) isSubClassOf(class *Class) bool {
	for k := class.superClass; k != nil; k = k.superClass {
		if k == class {
			return true
		}
	}
	return false
}

func (self *Class) isSubInterfaceOf(class *Class) bool {
	for _, superInterface := range self.interfaces {
		if superInterface == class || superInterface.isSubInterfaceOf(class) {
			return true
		}
	}
	return false
}