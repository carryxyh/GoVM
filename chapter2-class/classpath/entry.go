package classpath

import "os"

//路径分隔符
const pathListSeparator = string(os.PathListSeparator)

/**
	类路径的接口
 */
type Entry interface {
	readClass(className string) ([]byte, Entry, error)
	String() string
}

func newEntry(path string) Entry {
	return nil
}