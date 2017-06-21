package classpath

type ZipEntry struct {

}

func (self *ZipEntry) readClass(className string) ([]byte, Entry, error) {
	return nil, nil, nil
}

func newZipEntry(path string) *ZipEntry {
	return nil
}