package data

/*
ACC_SUPER was introduced to correct a problem with the invocation of super methods. The ACC_SUPER flag marks a class as compiled for the changed semantics of the opcode 183 instruction. It's purpose is similar to that of the class file version number as it allows the JVM to detect whether a class was compiled for the older or newer semantics of that instruction. Java 1.0.2 did not set and ignored ACC_SUPER while Java 1.1 and later always sets ACC_SUPER.

Before Java 1.1 the byte code instruction with opcode 183 that is now called invokespecial was called invokenonvirtual and had a partially different specification. It was used whenever instance methods had to be invoked without a virtual method lookup. This was the case for private methods, instance initializers (constructors) and to implement method invocations on super. But the latter case caused problems with evolving class libraries.

A method reference in byte code (CONSTANT_Methodref_info) not only defines the name and the argument and return types of a method but also the class to which it belongs. Opcode 183 gets such a method reference parameter and was meant to directly invoke the referenced method from the specified class without further lookups. In the case of invocations on super it was the compilers responsibility to resolve the closest super class that implements this method and generate a reference to it into the byte code.

Since Java 1.1 it was changed to essentially ignore the class referenced in CONSTANT_Methodref_info and to instead do the lookup for the closest super method with the given method name and signature in the JVM. This is usually done now when the class gets loaded or right before the instruction is executed or JIT compiled the first time.

Here is an example why this change was neccessary. In Java 1.0.2 the AWT classes Container and Component were defined this way:

class Component
{
    public void paint( Graphics g ) {}
}

class Container extends Component
{
    // inherits paint from Component but doesn't override it
}
In Java 1.1 the class Conatiner was changed to have it's own implementation of paint:

class Container extends Component
{
    public void paint( Graphics g ) {
		// todo
	}
}
Now when you had a direct or indirect subclass of Container that made a call on super.paint(g) and compiled it for 1.0.2 it generated a invokenonvirtual instruction for Component.paint since this was the first parent that had this method. But if you used this compiled class on a JVM that also had Container.paint it would still have called Component.paint which is not what you would expect.

On the other hand, when you compiled the class for 1.1 and executed it on a 1.0.2 JVM it would throw a AbstractMethodError or more likely for VMs of that era simply crash. To avoid the crash you had to write ((Component)super).paint(g) and compile it with a 1.1 compiler to get the desired behaviour in either VM. This would set ACC_SUPER but still generate the instruction to call Component.paint. A 1.0.2 VM would ignore ACC_SUPER and go straight to invoke Component.paint which is fine while a 1.1 VM would find ACC_SUPER set and thus do the lookup itself which would make it invoke Container.paint even though the byte code method reference was Component.paint.
*/
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
