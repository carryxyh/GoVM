package classpath

type CompositeEntry struct {

}

func (self *CompositeEntry) readClass(className string) ([]byte, Entry, error) {
	return nil, nil, nil
}

func (self *CompositeEntry) String() string {
	return nil
}

func newCompositeEntry(path string) *CompositeEntry {
	return nil
}