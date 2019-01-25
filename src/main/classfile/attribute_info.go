package classfile

type AttributeInfo interface {
	readInfo(reader *ClassReader)
}

func readAttributes(reader *ClassReader, cp ConstantPool) []AttributeInfo {
	attributesCount := reader.readUint16()
	attributes := make([]AttributeInfo, attributesCount)
	for i := range attributes {
		attributes[i] = readAttribute(reader, cp)
	}
	return attributes
}

func readAttribute(reader *ClassReader, cp ConstantPool) AttributeInfo {
	attrNameIndex := reader.readUint16()
	attName := cp.getUtf8(attrNameIndex)
	attrLen := reader.readUint32()
	attrInfo := newAttributeInfo(attName, attrLen, cp)
	attrInfo.readInfo(reader)
	return attrInfo
}

func newAttributeInfo(attName string, attLen uint32, cp ConstantPool) AttributeInfo {

}