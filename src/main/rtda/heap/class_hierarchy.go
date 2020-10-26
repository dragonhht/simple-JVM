package heap

func (self *Class) IsAssignableFrom(class *Class) bool {
	s ,t := class, self
	if s == t {
		return true
	}
	if !t.IsInterface() {
		return s.IsSubClassOf(t)
	} else {
		return s.IsImplements(t)
	}
}

func (self *Class) IsImplements(class *Class) bool {
	for c := self; c != nil; c = c.superClass {
		for _, i := range c.interfaces {
			if i == class || i.IsSubInterfaceOf(class) {
				return true
			}
		}
	}
	return false
}

func (self *Class) IsSubClassOf(class *Class) bool {
	for k := class.superClass; k != nil; k = k.superClass {
		if k == class {
			return true
		}
	}
	return false
}

func (self *Class) IsSubInterfaceOf(class *Class) bool {
	for _, superInterface := range self.interfaces {
		if superInterface == class || superInterface.IsSubInterfaceOf(class) {
			return true
		}
	}
	return false
}

func (self *Class) IsSuperClassOf(other *Class) bool {
	return other.IsSubClassOf(self)
}