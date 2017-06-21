package classpath

type DirEntry struct {
	absDir string
}

func newDirEntry(path string) *DirEntry {
	return nil
}

func (self *DirEntry) readClass(className string) ([]byte, Entry, error) {
	return nil, nil, nil
}