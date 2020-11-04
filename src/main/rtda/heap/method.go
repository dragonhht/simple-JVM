package heap

import "main/classfile"

type Method struct {
	ClassMember
	maxStack     uint16
	maxLocals    uint16
	code         []byte
	argSlotCount uint16
}

func newMethods(class *Class, cfMethods []*classfile.MemberInfo) []*Method {
	methods := make([]*Method, len(cfMethods))
	for i, cfMethod := range cfMethods {
		methods[i] = newMethod(class, cfMethod)
	}
	return methods
}

func newMethod(class *Class, cfMethod *classfile.MemberInfo) *Method {
	method := &Method{}
	method.class = class
	method.copyMemberInfo(cfMethod)
	method.copyAttributes(cfMethod)
	md := parseMethodDescriptor(method.descriptor)
	method.calcArgSlotCount(md.parameterTypes)
	if method.IsNative() {
		method.injectCodeAttribute(md.returnType)
	}
	return method
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
	return 0 != self.accessFlags&ACC_PUBLIC
}

func (self *Method) IsPrivate() bool {
	return 0 != self.accessFlags&ACC_PRIVATE
}

func (self *Method) IsProtected() bool {
	return 0 != self.accessFlags&ACC_PROTECTED
}

func (self *Method) IsStatic() bool {
	return 0 != self.accessFlags&ACC_STATIC
}

func (self *Method) IsFinal() bool {
	return 0 != self.accessFlags&ACC_FINAL
}

func (self *Method) IsSynchronized() bool {
	return 0 != self.accessFlags&ACC_SYNCHRONIZED
}

func (self *Method) IsBridge() bool {
	return 0 != self.accessFlags&ACC_BRIDGE
}

func (self *Method) IsVarargs() bool {
	return 0 != self.accessFlags&ACC_VARARGS
}

func (self *Method) IsNative() bool {
	return 0 != self.accessFlags&ACC_NATIVE
}

func (self *Method) IsAbstract() bool {
	return 0 != self.accessFlags&ACC_ABSTRACT
}

func (self *Method) IsStrict() bool {
	return 0 != self.accessFlags&ACC_STRICT
}

func (self *Method) IsSynthetic() bool {
	return 0 != self.accessFlags&ACC_SYNTHETIC
}

func (self *Method) ArgSlotCount() uint16 {
	return self.argSlotCount
}

func (self *Method) calcArgSlotCount(paramTypes []string) {
	for _, paramType := range paramTypes {
		self.argSlotCount++
		if paramType == "J" || paramType == "D" {
			self.argSlotCount++
		}
	}
	if !self.IsStatic() {
		self.argSlotCount++
	}
}

func (self *Method) injectCodeAttribute(returnType string) {
	self.maxStack = 4
	self.maxLocals = self.argSlotCount
	switch returnType[0] {
	case 'V': self.code = []byte{0xfe, 0xb1} // return
	case 'D': self.code = []byte{0xfe, 0xaf} // dreturn
	case 'F': self.code = []byte{0xfe, 0xae} // freturn
	case 'J': self.code = []byte{0xfe, 0xad} // lreturn
	case 'L', '[': self.code = []byte{0xfe, 0xb0} // areturn
	default:
		self.code = []byte{0xfe, 0xac} // ireturn
	}
}
