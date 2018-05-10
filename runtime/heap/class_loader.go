package heap

import (
	"jvm/classpath"
	"fmt"
	"jvm/classfile"
	"jvm/runtime"
)

type ClassLoader struct {
	cp          *classpath.Classpath
	loadedClass map[string]*Class
}

func NewClassLoader(cp *classpath.Classpath) *ClassLoader {
	return &ClassLoader{cp: cp,
		loadedClass: make(map[string]*Class),
	}
}

// loading class
func (loader *ClassLoader) LoadClass(name string) *Class {
	if class, ok := loader.loadedClass[name]; ok {
		return class
	}

	return loader.loadNonArrayClass(name)
}

func (loader *ClassLoader) loadNonArrayClass(name string) *Class {
	data, entry := loader.readClass(name)
	class := loader.defineClass(data)
	link(class)
	fmt.Printf("[Load %s from %s\n]", name, entry)
	return class
}

// linking class
func link(class *Class) {
	verify(class)
	prepare(class)
}

// Preparation involves creating the static fields for a class or interface and initializing
// such fields to their default values. This does not require the execution
// of any Java Virtual Machine code; explicit initializers for static fields are executed
// as part of initialization (§5.5), not preparation
func prepare(class *Class) {
	countInstantFields(class)
	countStaticFields(class)
	allocAndInitStaticVars(class)
}

func allocAndInitStaticVars(class *Class) {
	class.staticVars = runtime.NewVariables(class.staticCount)
	for _, field := range class.fields {
		if field.accessFlags.IsStatic() && field.accessFlags.IsFinal() {
			initStaticFinalVar(class, field)
		}
	}
}

func initStaticFinalVar(class *Class, field *Field) {
	staticVars := class.staticVars
	cp := class.cp
	cpIndex := field.ConstValueIndex()
	fieldId := field.fieldId
}

func countStaticFields(class *Class) {
	fieldId := uint(0)
	for _, field := range class.fields {
		if field.accessFlags.IsStatic() {
			field.fieldId = fieldId
			fieldId ++

			if field.isLongOrDouble() {
				fieldId ++
			}
		}
	}
	class.staticCount = fieldId
}

func countInstantFields(class *Class) {
	fieldId := uint(0)
	if class.superClass != nil {
		fieldId = class.superClass.instanceCount
	}

	for _, field := range class.fields {
		if !field.accessFlags.IsStatic() {
			field.fieldId = fieldId
			fieldId ++

			if field.isLongOrDouble() {
				fieldId ++
			}
		}
	}

	class.instanceCount = fieldId
}

// verification
func verify(class *Class) {
	// todo
}

func (loader *ClassLoader) readClass(name string) ([]byte, classpath.Entry) {
	data, entry, err := loader.cp.ReadClass(name)
	if err != nil {
		panic("java.lang.ClassNotFoundException:" + name)
	}
	return data, entry
}

func (loader *ClassLoader) defineClass(data []byte) *Class {
	cf := parseClassFile(data)
	class := NewClass(cf)
	class.classLoader = loader
	resolveSuperClass(class, cf.SuperClassName())
	resolveInterfaces(class, cf.InterfaceNames())
	loader.loadedClass[class.className] = class
	return class
}

func resolveInterfaces(class *Class, interfaceNames []string) {
	len := len(interfaceNames)

	if len > 0{
		interfaces := make([]*Class, len)
		for i, interfaceName := range interfaceNames {
			interfaces[i] = class.classLoader.LoadClass(interfaceName)
		}
		class.interfaces = interfaces
	}
}

func resolveSuperClass(class *Class, name string) {
	if name != "java/lang/Object" {
		class.superClass = class.classLoader.LoadClass(name)
	}
}

func parseClassFile(data []byte) *classfile.ClassFile {
	cf, err := classfile.Parse(data)

	if err != nil {
		panic("java.lang.ClassFormatError")
	}

	return cf
}




