package heap

/*
	Object结构体
*/
type Object struct {
	class *Class
	data  interface{}
	extra interface{}
}

func (self *Object) GetExtra() interface{} {
	return self.extra
}

func (self *Object) SetExtra(extra interface{}) {
	self.extra = extra
}

func (self *Object) Fields() Slots {
	return self.data.(Slots)
}

func (self *Object) Class() *Class {
	return self.class
}

func (self *Object) IsInstanceOf(class *Class) bool {
	return class.IsAssignableFrom(self.class)
}

func (self *Object) setRefVar(name string, descriptor string, ref *Object) {
	field := self.class.getField(name, descriptor, false)
	slots := self.data.(Slots)
	slots.SetRef(field.slotId, ref)
}

func (self *Object) GetRefVar(name string, descriptor string) *Object {
	field := self.class.getField(name, descriptor, false)
	slots := self.data.(Slots)
	return slots.GetRef(field.slotId)
}
