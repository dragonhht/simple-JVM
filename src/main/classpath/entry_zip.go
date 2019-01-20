package classpath

import (
	"archive/zip"
	"errors"
	"io/ioutil"
	"path/filepath"
)

// 处理zip和jar类型的文件

type ZipEntry struct {
	absPath string
	zipRc *zip.ReadCloser
}

func newZipEntry(path string) *ZipEntry {
	absPath, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	return &ZipEntry{absPath, nil}
}

func (self *ZipEntry) readClass(className string) ([]byte, Entry, error) {
	// 打开zip文件
	if self.zipRc == nil {
		err := self.openJar()
		if err != nil {
			return nil, self, err
		}
	}

	classFile := self.findClass(className)

	if classFile == nil {
		return nil, self, errors.New("class not found: " + className)
	}

	data, err := readClass(classFile)
	return data, self, err
}

func readClass(classFile *zip.File) ([]byte, error) {
	rc, err := classFile.Open()
	if err != nil {
		return nil, err
	}

	// 读取文件内容
	data, err := ioutil.ReadAll(rc)
	rc.Close()
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (self *ZipEntry) String() string {
	return self.absPath
}

func (self *ZipEntry) openJar() error {
	r, err := zip.OpenReader(self.absPath)
	if err == nil {
		self.zipRc = r
	}
	return err
}

func (self *ZipEntry) findClass(className string) *zip.File {
	for _, f := range self.zipRc.File {
		if f.Name == className {
			return f
		}
	}
	return nil
}
