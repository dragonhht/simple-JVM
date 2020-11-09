package classfile

import "fmt"

type ClassFile struct {
	// 魔数
	//magic uint32
	// 小版本，JSE1.2之后都为0
	minorVersion uint16
	// 大版本，JSE1.2及以前为45，之后每个大版本加1
	majorVersion uint16
	// 常量池
	constantPool ConstantPool
	// 类访问标志，指定是类还是接口,访问级别为public还是private
	accessFlags uint16
	// 类索引
	thisClass uint16
	// 超类索引
	superClass uint16
	// 接口索引表
	interfaces []uint16
	fields     []*MemberInfo
	methods    []*MemberInfo
	attributes []AttributeInfo
}

/*
	将字节解析为ClassFile结构体
*/
func Parse(classData []byte) (cf *ClassFile, err error) {
	defer func() {
		if r := recover(); r != nil {
			var ok bool
			err, ok = r.(error)
			if !ok {
				err = fmt.Errorf("%v", r)
			}
		}
	}()

	cr := &ClassReader{classData}
	cf = &ClassFile{}
	cf.read(cr)
	return
}

func (self *ClassFile) read(reader *ClassReader) {
	self.readAndCheckMagic(reader)
	self.readAndCheckVersion(reader)
	self.constantPool = readConstantPool(reader)
	self.accessFlags = reader.readUint16()
	self.thisClass = reader.readUint16()
	self.superClass = reader.readUint16()
	self.interfaces = reader.readUint16s()
	self.fields = readMembers(reader, self.constantPool)
	self.methods = readMembers(reader, self.constantPool)
	self.attributes = readAttributes(reader, self.constantPool)
}

func (self *ClassFile) ClassName() string {
	return self.constantPool.getClassName(self.thisClass)
}

func (self *ClassFile) SuperClassName() string {
	if self.superClass > 0 {
		return self.constantPool.getClassName(self.superClass)
	}
	// java.lang.Object没有父类
	return ""
}

func (self *ClassFile) InterfaceNames() []string {
	interfaceNames := make([]string, len(self.interfaces))
	for i, cpIndex := range self.interfaces {
		interfaceNames[i] = self.constantPool.getClassName(cpIndex)
	}
	return interfaceNames
}

/*
	读取并校验Class的魔数
*/
func (self *ClassFile) readAndCheckMagic(reader *ClassReader) {
	magic := reader.readUint32()
	if magic != 0xCAFEBABE {
		panic("java.lang.ClassFormatError: magic!")
	}
}

/*
	校验Class版本
*/
func (self *ClassFile) readAndCheckVersion(reader *ClassReader) {
	self.minorVersion = reader.readUint16()
	self.majorVersion = reader.readUint16()
	switch self.majorVersion {
	case 45:
		return
	case 46, 47, 48, 49, 50, 51, 52:
		if self.minorVersion == 0 {
			return
		}
	}
	panic("java.lang.UnsupportedClassVersionError!")
}

func (self *ClassFile) MajorVersion() uint16 {
	return self.majorVersion
}

func (self *ClassFile) Interfaces() []uint16 {
	return self.interfaces
}

func (self *ClassFile) SuperClass() uint16 {
	return self.superClass
}

func (self *ClassFile) ThisClass() uint16 {
	return self.thisClass
}

func (self *ClassFile) AccessFlags() uint16 {
	return self.accessFlags
}

func (self *ClassFile) MinorVersion() uint16 {
	return self.minorVersion
}

func (self *ClassFile) ConstantPool() ConstantPool {
	return self.constantPool
}

func (self *ClassFile) Attributes() []AttributeInfo {
	return self.attributes
}

func (self *ClassFile) Methods() []*MemberInfo {
	return self.methods
}

func (self *ClassFile) Fields() []*MemberInfo {
	return self.fields
}

func (self *ClassFile) SourceFileAttribute() *SourceFileAttribute {
	for _, attrInfo := range self.attributes {
		switch attrInfo.(type) {
		case *SourceFileAttribute:
			return attrInfo.(*SourceFileAttribute)
		}
	}
	return nil
}
