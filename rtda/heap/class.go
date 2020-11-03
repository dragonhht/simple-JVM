package heap

import (
	"simple-jvm/classfile"
	"strings"
)

type Class struct {
	accessFlags uint16
	name string
	superClassName string
	interfaceNames []string
	constantPool *ConstantPool
	fields []*Field
	methods []*Method
	loader *ClassLoader
	superClass *Class
	interfaces []*Class
	instanceSlotCount uint
	staticSlotCount uint
	staticVars Slots
	initStarted bool
}

func newClass(cf *classfile.ClassFile) *Class {
	class := &Class{}
	class.accessFlags = cf.AccessFlags()
	class.name = cf.ClassName()
	class.superClassName = cf.SuperClassName()
	class.interfaceNames = cf.InterfaceNames()
	class.constantPool = newConstantPool(class, cf.ConstantPool())
	class.fields = newFields(class, cf.Fields())
	class.methods = newMethods(class, cf.Methods())
	return class
}

func (self *Class) InitStarted() bool {
	return self.initStarted
}

func (self *Class) Loader() *ClassLoader {
	return self.loader
}

func (self *Class) StartInit() {
	self.initStarted = true
}

func (self *Class) StaticVars() Slots {
	return self.staticVars
}

func (self *Class) SuperClass() *Class {
	return self.superClass
}

func (self *Class) ConstantPool() *ConstantPool {
	return self.constantPool
}

func (self *Class) IsPublic() bool {
	return 0 != self.accessFlags & ACC_PUBLIC
}

func (self *Class) IsFinal() bool {
	return 0 != self.accessFlags & ACC_FINAL
}

func (self *Class) IsSuper() bool {
	return 0 != self.accessFlags & ACC_SUPER
}

func (self *Class) IsInterface() bool {
	return 0 != self.accessFlags & ACC_INTERFACE
}

func (self *Class) IsAbstract() bool {
	return 0 != self.accessFlags & ACC_ABSTRACT
}

func (self *Class) IsSynthetic() bool {
	return 0 != self.accessFlags & ACC_SYNTHETIC
}

func (self *Class) IsAnnotation() bool {
	return 0 != self.accessFlags & ACC_ANNOTATION
}

func (self *Class) IsEnum() bool {
	return 0 != self.accessFlags & ACC_ENUM
}

func (self *Class) isAccessibleTo(class *Class) bool {
	return self.IsPublic() || self.GetPackageName() == class.GetPackageName()
}

func (self *Class) GetPackageName() string {
	if i := strings.LastIndex(self.name, "/"); i >= 0 {
		return self.name[:i]
	}
	return ""
}

func (self *Class) NewObject() *Object {
	return newObject(self)
}

func (self *Class) Name() string {
	return self.name
}

func (self *Class) GetMainMethod() *Method {
	return self.getStaticMethod("main", "([Ljava/lang/String;)V")
}

func (self *Class) getStaticMethod(name string, descriptor string) *Method {
	for _, m := range self.methods {
		if m.IsStatic() && m.Name() == name && m.Descriptor() == descriptor {
			return m
		}
	}
	return nil
}

func (self *Class) GetClinitMethod() *Method {
	return self.getStaticMethod("<clinit>", "()V")
}

func (self *Class) ArrayClass() *Class {
	arrayClassName := getArrayClassName(self.name)
	return self.loader.LoadClass(arrayClassName)
}

func newObject(class *Class) *Object {
	return &Object{
		class:  class,
		data: newSlots(class.instanceSlotCount),
	}
}

func (self *Class) isJlObject() bool {
	return self.name == "java/lang/Object"
}
func (self *Class) isJlCloneable() bool {
	return self.name == "java/lang/Cloneable"
}
func (self *Class) isJioSerializable() bool {
	return self.name == "java/io/Serializable"
}