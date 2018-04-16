package classpath

import (
	"os"
	"strings"
)

const pathListSeparator = string(os.PathListSeparator)
type Entry interface {
	readClass(className string) (bety[], Entry, error)
	String() string
}

func newEntry() Entry {
	if strings.Contains(path, pathListSeparator) {
		return newCompositeEntry(path)
	}
	if strings.HasSuffix(path, "*") {
		return newWildcardEntry(path)
	}
	if strings.HasSuffix(path, ".jar") || strings.HasSuffix(path, ".JAR") || 
	strings.HasSuffix(path, ".zip") || strings.HasSuffix(path, ".ZIP") {
		return newZipEntry(path)
	}
	return newDirEntry(path)
}
