package classfile

type MarkerAttribute struct {

}

func (self *MarkerAttribute) readInfo(reader *ClassReader) {
	// 不需要做操作
}

type DeprecatedAttribute struct {
	MarkerAttribute
}

type SyntheticAttribute struct {
	MarkerAttribute
}