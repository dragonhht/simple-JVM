package heap

import "github.com/kuangcp/simple-jvm/src/main/classfile"

type Method struct {
	ClassMember
	maxStack uint16
	maxLocals uint16
	code []byte
	argSlotCount uint
}

func newMethods(class *Class, cfMethods []*classfile.MemberInfo) []*Method {
	methods :=make([]*Method, len(cfMethods))
	for i, cfMethod := range cfMethods {
		methods[i] = &Method{}
		methods[i].class = class
		methods[i].copyMemberInfo(cfMethod)
		methods[i].copyAttributes(cfMethod)
		methods[i].calcArgSlotCount()
	}
	return methods
}

func (self *Method) MaxStack() uint16 {
	return self.maxStack
}

func (self *Method) Code() []byte {
	return self.code
}

func (self *Method) MaxLocals() uint16 {
	return self.maxLocals
}

func (self *Method) copyAttributes(cfMethod *classfile.MemberInfo) {
	if codeAttr := cfMethod.CodeAttribute(); codeAttr != nil {
		self.maxStack = codeAttr.MaxStack()
		self.maxLocals = codeAttr.MaxLoacls()
		self.code = codeAttr.Code()
	}
}

func (self *Method) IsPublic() bool {
	return 0 != self.accessFlags & ACC_PUBLIC
}

func (self *Method) IsPrivate() bool {
	return 0 != self.accessFlags & ACC_PRIVATE
}

func (self *Method) IsProtected() bool {
	return 0 != self.accessFlags & ACC_PROTECTED
}

func (self *Method) IsStatic() bool {
	return 0 != self.accessFlags & ACC_STATIC
}

func (self *Method) IsFinal() bool {
	return 0 != self.accessFlags & ACC_FINAL
}

func (self *Method) IsSynchronized() bool {
	return 0 != self.accessFlags & ACC_SYNCHRONIZED
}

func (self *Method) IsBridge() bool {
	return 0 != self.accessFlags & ACC_BRIDGE
}

func (self *Method) IsVarargs() bool {
	return 0 != self.accessFlags & ACC_VARARGS
}

func (self *Method) IsNative() bool {
	return 0 != self.accessFlags & ACC_NATIVE
}

func (self *Method) IsAbstract() bool {
	return 0 != self.accessFlags & ACC_ABSTRACT
}

func (self *Method) IsStrict() bool {
	return 0 != self.accessFlags & ACC_STRICT
}

func (self *Method) IsSynthetic() bool {
	return 0 != self.accessFlags & ACC_SYNTHETIC
}

func (self *Method) ArgSlotCount() uint {
	return self.argSlotCount
}

func (self *Method) calcArgSlotCount() {
	parsedDescriptor := parseMethodDescriptor(self.descriptor)
	for _, paramType := range parsedDescriptor.parameterTypes {
		self.argSlotCount++
		if paramType == "J" || paramType == "D" {
			self.argSlotCount++
		}
	}
	if !self.IsStatic() {
		self.argSlotCount++
	}
}