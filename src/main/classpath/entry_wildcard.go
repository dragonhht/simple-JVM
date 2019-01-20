package classpath

import (
	"os"
	"path/filepath"
	"strings"
)

func newWildcardEntry(path string) CompositeEntry {
	// 去除结尾处的*
	baseDir := path[:len(path) - 1]
	compositeEbtry := []Entry{}

	walkFn := func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return filepath.SkipDir
		}
		if strings.HasSuffix(path, ".jar") || strings.HasSuffix(path, ".JAR") {
			jarEntry := newZipEntry(path)
			compositeEbtry = append(compositeEbtry, jarEntry)
		}
		return nil
	}
	filepath.Walk(baseDir, walkFn)

	return compositeEbtry
}
