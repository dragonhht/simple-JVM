package classfile

type CodeAttribute struct {
	cp             ConstantPool
	maxStack       uint16
	maxLoacls      uint16
	code           []byte
	exceptionTable []*ExceptionTableEntry
	attributes     []AttributeInfo
}

type ExceptionTableEntry struct {
	startPc   uint16
	endPc     uint16
	handlerPc uint16
	catchType uint16
}

func (self *ExceptionTableEntry) StartPc() uint16 {
	return self.startPc
}

func (self *ExceptionTableEntry) EndPc() uint16 {
	return self.endPc
}

func (self *ExceptionTableEntry) HandlerPc() uint16 {
	return self.handlerPc
}

func (self *ExceptionTableEntry) CatchType() uint16 {
	return self.catchType
}

func (self *CodeAttribute) ExceptionTable() []*ExceptionTableEntry {
	return self.exceptionTable
}

func (self *CodeAttribute) Code() []byte {
	return self.code
}

func (self *CodeAttribute) MaxStack() uint16 {
	return self.maxStack
}

func (self *CodeAttribute) MaxLoacls() uint16 {
	return self.maxLoacls
}

func (self *CodeAttribute) readInfo(reader *ClassReader) {
	self.maxStack = reader.readUint16()
	self.maxLoacls = reader.readUint16()
	condeLength := reader.readUint32()
	self.code = reader.readBytes(condeLength)
	self.exceptionTable = readExceptionTable(reader)
	self.attributes = readAttributes(reader, self.cp)
}

func (self *CodeAttribute) LineNumberTableAttribute() *LineNumberTableAttribute {
	for _, attrInfo := range self.attributes {
		switch attrInfo.(type) {
		case *LineNumberTableAttribute:
			return attrInfo.(*LineNumberTableAttribute)
		}
	}
	return nil
}

func readExceptionTable(reader *ClassReader) []*ExceptionTableEntry {
	exceptionTableLen := reader.readUint16()
	exceptionTable := make([]*ExceptionTableEntry, exceptionTableLen)
	for i := range exceptionTable {
		exceptionTable[i] = &ExceptionTableEntry{
			startPc:   reader.readUint16(),
			endPc:     reader.readUint16(),
			handlerPc: reader.readUint16(),
			catchType: reader.readUint16(),
		}
	}
	return exceptionTable
}
