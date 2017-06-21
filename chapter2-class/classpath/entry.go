package classpath

import (
	"os"
	"strings"
)

//路径分隔符
const pathListSeparator = string(os.PathListSeparator)

/**
	类路径的接口
 */
type Entry interface {
	readClass(className string) ([]byte, Entry, error)
	String() string
}

/**
	根据路径创建不同的Entry
 */
func newEntry(path string) Entry {
	if strings.Contains(path, pathListSeparator) {
		return newDirEntry(path)
	}
	if strings.HasSuffix(path, "*") {
		return newWildCardEntry(path)
	}
	if strings.HasSuffix(path, ".jar") || strings.HasSuffix(path, "JAR") || strings.HasSuffix(path, ".zip") || strings.HasSuffix(path, ".ZIP") {
		return newZipEntry(path)
	}
	return newDirEntry(path)
}