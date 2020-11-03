package classfile

type ConstantNameAndTypeInfo struct {
	nameIndex       uint16
	descriptorIndex uint16
}

func (self *ConstantNameAndTypeInfo) readInfo(read *ClassReader) {
	self.nameIndex = read.readUint16()
	self.descriptorIndex = read.readUint16()
}
