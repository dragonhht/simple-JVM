package classfile

type CodeAttribute struct {
	cp ConstantPool
	maxStack uint16
	maxLoacls uint16
	code []byte
	exceptionTable []*ExceptionTableEntry
	attributes []AttributeInfo
}

type ExceptionTableEntry struct {
	startPc uint16
	endPc uint16
	handlerPc uint16
	catchType uint16
}

func (self *CodeAttribute) readInfo(reader *ClassReader) {
	self.maxStack = reader.readUint16()
	self.maxLoacls = reader.readUint16()
	condeLength := reader.readUint32()
	self.code = reader.readBytes(condeLength)
	self.exceptionTable = readExceptionTable(reader)
	self.attributes = readAttributes(reader, self.cp)
}

func readExceptionTable(reader *ClassReader) []*ExceptionTableEntry {
	exceptionTableLen := reader.readUint16()
	exceptionTable := make([]*ExceptionTableEntry, exceptionTableLen)
	for i := range exceptionTable {
		exceptionTable[i] = &ExceptionTableEntry{
			startPc: reader.readUint16(),
			endPc: reader.readUint16(),
			handlerPc: reader.readUint16(),
			catchType: reader.readUint16(),
		}
	}
	return exceptionTable
}