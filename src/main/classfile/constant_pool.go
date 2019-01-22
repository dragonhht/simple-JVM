package classfile

/*
	常量池结构体
 */
type ConstantPool []ConstantInfo

func readConstantPool(reader *ClassReader) ConstantPool {

}

func (self *ConstantPool) getConstantInfo(index uint16) ConstantInfo {

}

func (self *ConstantPool) getNameAndType(index uint16) (string, string) {

}

func (self *ConstantPool) getClassName(index uint16) string {

}

func (self *ConstantPool) getUtf8(index uint16) string {

}
