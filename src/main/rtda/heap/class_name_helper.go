package heap

var primitiveTypes = map[string]string {
	"void": "V",
	"boolean": "Z",
	"byte": "B",
	"short": "S",
	"int": "I",
	"long": "J",
	"char": "C",
	"float": "F",
	"double": "D",
}

func getArrayClassName(name string) string {
	return "[" + toDescriptor(name)
}

func toDescriptor(name string) string {
	if name[0] == '[' {
		return name
	}
	if d, ok := primitiveTypes[name]; ok {
		return d
	}
	return "L" + name + ";"
}

func getComponentClassName(name string) string {
	if name[0] == '[' {
		componentTypeDescriptor := name[1:]
		return toClassName(componentTypeDescriptor)
	}
	panic("Not array: " + name)
}

func toClassName(descriptor string) string {
	if descriptor[0] == '[' {
		return descriptor
	}
	if descriptor[0] == 'L' {
		return descriptor[1 : len(descriptor) - 1]
	}
	for className, d := range primitiveTypes {
		if d == descriptor {
			return className
		}
	}
	panic("Invalid descriptor: " + descriptor)
}