package classpath

type WildCardEntry struct {

}

func (self *WildCardEntry) readClass(classname string) ([]byte, Entry, error) {
	return nil, nil, nil
}

func (self *WildCardEntry) String() string {
	return nil
}

func newWildCardEntry(path string) *WildCardEntry {
	return nil
}
