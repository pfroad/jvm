package data

const (
	ACC_PUBLIC       = 0x0001 // class field method
	ACC_PRIVATE      = 0x0002 // field method
	ACC_PROTECTED    = 0x0004 // field method
	ACC_STATIC       = 0x0008 // field method
	ACC_FINAL        = 0x0010 // class field method
	ACC_SUPER        = 0x0020 // class
	ACC_SYNCHRONIZED = 0x0020 // method
	ACC_VOLATILE     = 0x0040 // field
	ACC_BRIDGE       = 0x0040 // method
	ACC_TRANSIENT    = 0x0080 // field
	ACC_VARARGS      = 0x0080 // method
	ACC_NATIVE       = 0x0100 // method
	ACC_INTERFACE    = 0x0200 // class
	ACC_ABSTRACT     = 0x0400 // class method
	ACC_STRICT       = 0x0800 // method
	ACC_SYNTHETIC    = 0x1000 // class field method
	ACC_ANNOTATION   = 0x2000 // class
	ACC_ENUM         = 0x4000 // class field
)

type AccessFlags struct {
	accessFlags uint16
}

func (aFlags *AccessFlags) IsPublic() bool {
	return 0 != aFlags.accessFlags&ACC_PUBLIC
}

func (aFlags *AccessFlags) IsPrivate() bool {
	return 0 != aFlags.accessFlags&ACC_PRIVATE
}

func (aFlags *AccessFlags) IsProtected() bool {
	return 0 != aFlags.accessFlags&ACC_PROTECTED
}

func (aFlags *AccessFlags) IsStatic() bool {
	return 0 != aFlags.accessFlags&ACC_STATIC
}

func (aFlags *AccessFlags) IsFinal() bool {
	return 0 != aFlags.accessFlags&ACC_FINAL
}

func (aFlags *AccessFlags) IsSuper() bool {
	return 0 != aFlags.accessFlags&ACC_SUPER
}

func (aFlags *AccessFlags) IsSynchronized() bool {
	return 0 != aFlags.accessFlags&ACC_SYNCHRONIZED
}

func (aFlags *AccessFlags) IsVolatile() bool {
	return 0 != aFlags.accessFlags&ACC_VOLATILE
}

func (aFlags *AccessFlags) IsBridge() bool {
	return 0 != aFlags.accessFlags&ACC_BRIDGE
}

// ACC_TRANSIENT
func (aFLags *AccessFlags) IsTransient() bool {
	return 0 != aFLags.accessFlags&ACC_TRANSIENT
}

// ACC_VARARGS
func (aFlags *AccessFlags) IsVarArgs() bool {
	return 0 != aFlags.accessFlags&ACC_VARARGS
}

// ACC_NATIVE
func (aFlags *AccessFlags) IsNative() bool {
	return 0 != aFlags.accessFlags&ACC_NATIVE
}

// ACC_INTERFACE
func (aFlags *AccessFlags) IsInterface() bool {
	return 0 != aFlags.accessFlags&ACC_INTERFACE
}

// ACC_ABSTRACT
func (aFlags *AccessFlags) IsAbstract() bool {
	return 0 != aFlags.accessFlags&ACC_ABSTRACT
}

// ACC_STRICT
func (aFlags *AccessFlags) IsStrict() bool {
	return 0 != aFlags.accessFlags&ACC_STRICT
}

// ACC_SYNTHETIC
func (aFlags *AccessFlags) IsSynthetic() bool {
	return 0 != aFlags.accessFlags&ACC_SYNTHETIC
}

// ACC_ANNOTATION
func (aFlags *AccessFlags) IsAnnotation() bool {
	return 0 != aFlags.accessFlags&ACC_ANNOTATION
}

// ACC_ENUM
func (aFlags *AccessFlags) name() bool {
	return 0 != aFlags.accessFlags&ACC_ENUM
}
