package heap

import "github.com/kuangcp/simple-jvm/src/main/classfile"

type ClassMember struct {
	accessFlags uint16
	name string
	descriptor string
	class *Class
}

func (self *ClassMember) copyMemberInfo(memberInfo *classfile.MemberInfo) {
	self.accessFlags = memberInfo.AccessFlags()
	self.name = memberInfo.Name()
	self.descriptor = memberInfo.Descriptor()
}

func (self *ClassMember) Class() *Class {
	return self.class
}

func (self *ClassMember) Name() string {
	return self.name
}

func (self *ClassMember) IsPublic() bool {
	return 0 != self.accessFlags & ACC_PUBLIC
}

func (self *ClassMember) IsPrivate() bool {
	return 0 != self.accessFlags & ACC_PRIVATE
}

func (self *ClassMember) IsProtected() bool {
	return 0 != self.accessFlags & ACC_PROTECTED
}

func (self *ClassMember) Descriptor() string {
	return self.descriptor
}

func (self *ClassMember) isAccessibleTo(class *Class) bool {
	if self.IsPublic() {
		return true
	}
	c := self.class
	if self.IsProtected() {
		return class == c || class.IsSubClassOf(c) ||
			c.GetPackageName() == class.GetPackageName()
	}
	if !self.IsPrivate() {
		return c.GetPackageName() == class.GetPackageName()
	}
	return class == c
}