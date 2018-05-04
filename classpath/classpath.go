package classpath

import (
	"fmt"
	"os"
	"path/filepath"
)

type Classpath struct {
	bootClasspath Entry
	extClasspath  Entry
	userClasspath Entry
}

func Parse(jreOption, cpOption string) *Classpath {
	cp := &Classpath{}
	cp.parseBootClasspath(jreOption)
	cp.parseUserClasspath(cpOption)
	return cp
}

func (self *Classpath) parseBootClasspath(jreOption string) {
	jreDir := getJreDir(jreOption)
	jreLib := filepath.Join(jreDir, "lib", "*")
	fmt.Printf("jre lib %s", jreLib)
	self.bootClasspath = newWildcardEntry(jreLib)
	self.extClasspath = newWildcardEntry(filepath.Join(jreDir, "lib", "ext", "*"))
}

func (self *Classpath) parseUserClasspath(cp string) {
	if cp == "" {
		cp = "."
	}
	self.userClasspath = newEntry(cp)
}

func getJreDir(jreOption string) string {
	if jreOption != "" && exists(jreOption) {
		return jreOption
	}

	if exists("./jre") {
		return "./jre"
	}

	if jh := os.Getenv("JAVA_HOME"); jh != "" {
		return filepath.Join(jh, "jre")
	}

	panic("Cannot find jre folder!")
}

func exists(path string) bool {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}

	return true
}

func (self *Classpath) ReadClass(className string) ([]byte, Entry, error) {
	classFile := className + ".class"
	fmt.Printf("Load class %s\n", classFile)
	data, entry, err := self.bootClasspath.readClass(classFile)
	if err == nil {
		return data, entry, err
	}

	//fmt.Printf("Errors: %s", err)
	if data, entry, err := self.extClasspath.readClass(classFile); err == nil {
		return data, entry, err
	}

	return self.userClasspath.readClass(classFile)
}

func (self *Classpath) String() string {
	return self.userClasspath.String()
}
