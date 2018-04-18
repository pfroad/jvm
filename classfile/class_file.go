package classfile

import (
	"fmt"
)

/* java Class file struct
ClassFile {
u4 magic;
u2 minor_version;
u2 major_version;
u2 constant_pool_count;
cp_info constant_pool[constant_pool_count-1];
u2 access_flags;
u2 this_class;
u2 super_class;
u2 interfaces_count;
u2 interfaces[interfaces_count];
u2 fields_count;
field_info fields[fields_count];
u2 methods_count;
method_info methods[methods_count];
u2 attributes_count;
attribute_info attributes[attributes_count];
}
*/

// import "fmt"

// ClassFile class file info
type ClassFile struct {
	magic        uint32
	minorVersion uint16
	majorVersion uint16
	constantPool ConstantPool
	accessFlags  uint16
	thisClass    uint16
	superClass   uint16
	interfaces   []uint16
	fields       []*MemberInfo
	methods      []*MemberInfo
	attributes   []AttributeInfo
}

func Parse(data []byte) (cf *ClassFile, err error) {
	defer func() {
		if r := recover(); r != nil {
			// r.(type) if success to convert r to type return r, true else return r, false
			err, ok := r.(error)

			if !ok {
				err = fmt.Errorf("%v", r)
			}
		}
	}()
	cr := &ClassReader{data}
	cf = &ClassFile{}
	cf.read(cr)
	return cf, nil
}

func (cf *ClassFile) read(classReader *ClassReader) {
	cf.readAndCheckMagic(classReader)
	cf.readAndCheckVersion(classReader)
	cf.constantPool = readConstantPool(classReader)
	cf.accessFlags = classReader.readUint16()
	cf.thisClass = classReader.readUint16()
	cf.superClass = classReader.readUint16()
	cf.interfaces = classReader.readUint16s()
	cf.fields = readMembers(classReader, cf.constantPool)
	cf.methods = readMembers(classReader, cf.constantPool)
	cf.attributes = readAttributes(classReader, cf.constantPool)
}

// magic number is bytes which in the file header, java class file is 0xCAFEBABE
func (cf *ClassFile) readAndCheckMagic(classReader *ClassReader) {
	cf.magic = classReader.readUint32()
	if cf.magic != 0XCAFEBABE {
		panic("java.lang.ClassFormatError: magic!")
	}
}

func (cf *ClassFile) readAndCheckVersion(classReader *ClassReader) {
	cf.minorVersion = classReader.readUint16()
	cf.majorVersion = classReader.readUint16()
	if cf.majorVersion > 45 && cf.majorVersion <= 52 {
		if cf.minorVersion == 0 {
			return
		}
	} else if cf.majorVersion == 45 {
		return
	}
	panic("java.lang.UnsupportedClassVersionError")
}

// Magic class file magic number is 0xcafebabe
func (cf *ClassFile) Magic() uint32 {
	return cf.magic
}

// MinorVeresion getter java class minorVersion
func (cf *ClassFile) MinorVeresion() uint16 {
	return cf.minorVersion
}

// MajorVersion jdk majorVersion(45, 46 ... 52(jdk1.8))
func (cf *ClassFile) MajorVersion() uint16 {
	return cf.majorVersion
}

// AccessFlags flags is
/*  ACC_PUBLIC 0x0001 public
ACC_PRIVATE 0x0002 private
ACC_PROTECTED 0x0004 protected
ACC_STATIC 0x0008 static
ACC_FINAL 0x0010 final
ACC_VOLATILE 0x0040 volatile
ACC_TRANSIENT 0x0080 transient
ACC_SYNTHETIC 0x1000 编译器自动产生
ACC_ENUM 0x4000 enum
*/
func (cf *ClassFile) AccessFlags() uint16 {
	return cf.accessFlags
}

// ConstantPool getter class file constant pool
func (cf *ClassFile) ConstantPool() ConstantPool {
	return cf.constantPool
}

// Fields getter class fields
/* fields construct
field_info {
u2 access_flags;
u2 name_index;
u2 descriptor_index;
u2 attributes_count;
attribute_info attributes[attributes_count];
}*/
func (cf *ClassFile) Fields() []*MemberInfo {
	return cf.Fields
}

func (cf *ClassFile) Methods() []*MemberInfo {
	return cf.methods
}

func (cf *ClassFile) ClassName() string {
	return cf.constantPool.getClassName(cf.thisClass)
}

func (cf *ClassFile) SuperClassName() string {
	// java.lang.Object no super class
	if cf.superClass > 0 {
		return cf.constantPool.getSuperClassName(cf.superClass)
	}

	return ""
}

func (cf *ClassFile) InterfaceNames() []string {
	interfaceNames := make([]string, len(cf.interfaces))
	for i, cpInfo := range cf.interfaces {
		interfaceNames[i] = cf.constantPool.getClassName(cpInfo)
	}
	return interfaceNames
}
