package classpath

import (
	"os"
	"strings"
)

const pathListSeparator = string(os.PathListSeparator)


type Entry interface {
	/*
		加载Class文件
	 */
	readClass(className string) ([]byte, Entry, error)
	String() string
}

/*
	根据参数类型创建实例
 */
func newEntry(path string) Entry {
	// 是否包含分隔符
	if strings.Contains(path, pathListSeparator) {
		return newCompositeEntry(path)
	}
	// 是否以 * 结尾
	if strings.HasSuffix(path, "*") {
		return newWildcardEntry(path)
	}
	// zip或jar文件
	if strings.HasSuffix(path, ".jar") || strings.HasSuffix(path, ".JAR") ||
		strings.HasSuffix(path, ".zip") || strings.HasSuffix(path, ".ZIP"){
		return newZipEntry(path)
	}
	return newDirEntry(path)
}
