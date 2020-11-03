package classpath

import (
	"io/ioutil"
	"path/filepath"
)

type DirEntry struct {
	absDir string
}

func newDirEntry(path string) *DirEntry {
	// 将路径转换为绝对路径
	absPath, err := filepath.Abs(path)
	if err != nil {
		// 终止程序
		panic(err)
	}
	return &DirEntry{absPath}
}

func (self *DirEntry) readClass(className string) ([]byte, Entry, error) {
	// 将目录与class文件拼接成一个完整的路径
	fileName := filepath.Join(self.absDir, className)
	// 读取文件内容
	data, err := ioutil.ReadFile(fileName)
	return data, self, err
}

func (self *DirEntry) String() string {
	return self.absDir
}
