package classfile

// tag常量值
const (
	CONSTANT_CLASS = 7
	CONSTANT_FIELD_REF = 9
	CONSTANT_METHOD_REF = 10
	CONSTANT_INTERFACE_METHOD_REF = 11
	CONSTANT_STRING = 8
	CONSTANT_INTEGER = 3
	CONSTANT_FLOAT = 4
	CONSTANT_LONG = 5
	CONSTANT_DOUBLE = 6
	CONSTANT_NAME_TYPE = 12
	CONSTANT_UTF8 = 1
	// TODO 还有三种
	CONSTANT_METHOD_HANDLE = 15
	CONSTANT_METHOD_TYPE = 16
	CONSTANT_INVOKE_DYNAMIC = 18
)

// 表示常量信息的接口
type ConstantInfo interface {
	readInfo(reader *ClassReader)
}

func readConstantInfo(reader *ClassReader, cp ConstantPool) ConstantInfo {
	tag := reader.readUint8()
	c := newConstantInfo(tag, cp)
	c.readInfo(reader)
	return c
}

func newConstantInfo(tag uint8, cp ConstantPool) ConstantInfo {
	switch tag {
	case CONSTANT_INTEGER:
		return &ConstantIntegerInfo{}
	case CONSTANT_FLOAT:
		return &ConstantFloatInfo{}
	case CONSTANT_LONG:
		return &ConstantLongInfo{}
	case CONSTANT_DOUBLE:
		return &ConstantDoubleInfo{}
	case CONSTANT_UTF8:
		return &ConstantUtf8Info{}
	case CONSTANT_STRING:
		return &ConstantStringInfo{cp: cp}
	case CONSTANT_CLASS:
		return &ConstantClassInfo{cp:cp}
	case CONSTANT_FIELD_REF:
		return &ConstantFieldrefInfo{ConstantMemberRefInfo{cp: cp}}
	case CONSTANT_METHOD_REF:
		return &ConstantMethodrefInfo{ConstantMemberRefInfo{cp: cp}}
	case CONSTANT_INTERFACE_METHOD_REF:
		return &ConstantInterfaceMethodrefInfo{ConstantMemberRefInfo{cp: cp}}
	case CONSTANT_NAME_TYPE:
		return &ConstantNameAndTypeInfo{}
	default:
		panic("java.lang.ClassFormatErr: constant pool tag!")
	}
}