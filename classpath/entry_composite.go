package classpath

import (
	"errors"
	"strings"
)

type CompositeEntry []Entry

func newCompositeEntry(pathList string) CompositeEntry {
	compositeEntry := CompositeEntry{}
	for _, path := range strings.Split(pathList, pathListSeparator) {
		entry := newEntry(path)
		compositeEntry = append(compositeEntry, entry)
	}
	// fmt.Printf("jars %s\n", compositeEntry.Name())
	return compositeEntry
}

func (self CompositeEntry) readClass(className string) ([]byte, Entry, error) {
	for _, entry := range self {
		// fmt.Printf("jar %s\n", entry.Name())
		data, from, err := entry.readClass(className)
		if err == nil {
			// fmt.Printf("Errors: %s\n", err)
			return data, from, err
		}
	}
	return nil, nil, errors.New("Class cannot be found:" + className)
}

func (self CompositeEntry) String() string {
	strs := make([]string, len(self))
	for i, entry := range self {
		strs[i] = entry.String()
	}
	return strings.Join(strs, pathListSeparator)
}
