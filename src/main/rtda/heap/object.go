package heap

/*
	Object结构体
 */
type Object struct {
	class *Class
	fields Slots
}

func (self *Object) Fields() Slots {
	return self.fields
}

func (self *Object) Class() *Class {
	return self.class
}