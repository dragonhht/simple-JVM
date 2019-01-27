package classpath

import (
	"os"
	"path/filepath"
)

type Classpath struct {
	// 启动类路径
	bootClasspath Entry
	// 扩展类路径
	extClasspath Entry
	// 用户类路径
	userClasspath Entry
}

func Parse(jreOption, cpOption string) *Classpath {
	cp := &Classpath{}
	cp.parseBootAndExtClasspath(jreOption)
	cp.parseUserClasspath(cpOption)
	return cp
}

func (self *Classpath) ReadClass(className string) ([]byte, Entry, error) {
	className = className + ".class"
	if data, entry, err := self.bootClasspath.readClass(className); err == nil {
		return data, entry, err
	}
	if data, entry, err := self.extClasspath.readClass(className); err == nil {
		return data, entry, err
	}
	return self.userClasspath.readClass(className)
}

func (self *Classpath) String() string {
	return self.userClasspath.String()
}

func (self *Classpath) parseBootAndExtClasspath(jreOption string) {
	jreDir := getJreDir(jreOption)

	// jre/lib下的class
	jreLibPath := filepath.Join(jreDir, "lib", "*")
	self.bootClasspath = newWildcardEntry(jreLibPath)

	// jre/lib/ext/* 下的class
	jreExtPath := filepath.Join(jreDir, "lib", "ext", "*")
	self.extClasspath = newWildcardEntry(jreExtPath)
}

func (self *Classpath) parseUserClasspath(cpOption string) {
	if cpOption == "" {
		cpOption = "."
	}
	self.userClasspath = newEntry(cpOption)
}

func getJreDir(jreDir string) string {
	if jreDir != "" && exists(jreDir) {
		return jreDir
	}
	if exists("./jre") {
		return "./jre"
	}
	if javaHonme := os.Getenv("JAVA_HOME"); javaHonme != "" {
		return filepath.Join(javaHonme, "jre")
	}
	panic("Can not find jre folder!")
}

/*
	判断目录是否存在
 */
func exists(path string) bool {
	if _, err := os.Stat(path); err != nil {
		if os.IsExist(err) {
			return false
		}
	}
	return true
}
