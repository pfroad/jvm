package data

import (
	"jvm/classfile"
	"strings"
)

type Class struct {
	accessFlags   AccessFlags
	className     string
	cp            *ConstantPool
	fields        []*Field
	methods       []*Method
	classLoader   *ClassLoader
	superClass    *Class
	interfaces    []*Class
	instanceCount uint
	staticCount   uint
	staticVars    Variables
	initStarted   bool
}

func NewClass(cf *classfile.ClassFile) *Class {
	class := &Class{}
	class.accessFlags = AccessFlags{cf.AccessFlags()}
	class.className = cf.ClassName()
	class.cp = newConstantPool(class, cf.ConstantPool())
	class.fields = newFields(class, cf.Fields())
	class.methods = newMethods(class, cf.Methods())
	return class
}

func (class *Class) isAccessTo(other *Class) bool {
	return class.accessFlags.IsPublic() || class.GetPackageName() == other.GetPackageName()
}

func (class *Class) GetPackageName() string {
	if i := strings.LastIndex(class.className, "/"); i >= 0 {
		return class.className[:i]
	}

	return ""
}

func (class *Class) ClassName() string {
	return class.className
}

func (class *Class) ConstantPool() *ConstantPool {
	return class.cp
}

func (class *Class) IsInterface() bool {
	return class.accessFlags.IsInterface()
}

func (class *Class) IsAbstract() bool {
	return class.accessFlags.IsAbstract()
}

func (class *Class) NewObject() *Object {
	return &Object{class: class, data: NewVariables(class.instanceCount)}
}

func (class *Class) StaticVars() Variables {
	return class.staticVars
}

func (class *Class) SuperClass() *Class {
	return class.superClass
}

func (class *Class) isAssignableFrom(other *Class) bool {
	if other == class {
		return true
	}

	if !class.IsInterface() {
		//if superClass := other.superClass; superClass != nil {
		//	return class.isAssignableFrom(superClass)
		//}
		if other.isExtendClass(class) {
			return true
		}
	} else {
		//for c := other; c != nil; c = c.superClass {
		//	for _, iface := range c.interfaces {
		//		return class.isAssignableFrom(iface)
		//	}
		//}
		if other.isImplements(class) {
			return true
		}
	}

	return false
}

func (class *Class) isExtendClass(other *Class) bool {
	for c := class.superClass; c != nil; c = c.superClass {
		if c == other {
			return true
		}
	}

	return false
}

func (class *Class) IsSubClassOf(other *Class) bool {
	return class.isExtendClass(other)
}

func (class *Class) IsSuperClassOf(other *Class) bool {
	return other.isExtendClass(class)
}

func (class *Class) isExtendInterface(other *Class) bool {
	for _, iface := range class.interfaces {
		if iface == other || iface.isExtendInterface(other) {
			return true
		}
	}

	return false
}

func (class *Class) isImplements(other *Class) bool {
	for c := class; c != nil; c = c.superClass {
		for _, iface := range c.interfaces {
			if iface == other || iface.isExtendInterface(other) {
				return true
			}
		}
	}

	return false
}

func (class *Class) GetMainMethod() *Method {
	method := class.getStaticMethod("main", "([Ljava/lang/String;)V")

	if method != nil && method.IsPublic() {
		return method
	}

	return nil
}

func (class *Class) getStaticMethod(methodName string, descriptor string) *Method {
	for _, method := range class.methods {
		if method.accessFlags.IsStatic() &&
			method.Name() == methodName && method.descriptor == descriptor {
			return method
		}
	}

	return nil
}

func (class *Class) IsSuper() bool {
	return class.accessFlags.IsSuper()
}

func (class *Class) InitStarted() bool {
	return class.initStarted
}

func (class *Class) StartInit() {
	class.initStarted = true
}

/* GetClinitMethod init is the (or one of the) constructor(s) for the instance, and non-static field initialization.
 clinit are the static initialization blocks for the class, and static field initialization.
Classfile /F:/test/demo/target/classes/com/pfroad/demo/FibonacciTest.class
  Last modified 2018-5-18; size 778 bytes
  MD5 checksum 481cce926cfb66e997ebbe9a597bc821
  Compiled from "FibonacciTest.java"
public class com.pfroad.demo.FibonacciTest
  minor version: 0
  major version: 50
  flags: ACC_PUBLIC, ACC_SUPER
Constant pool:
   #1 = Methodref          #10.#31        // java/lang/Object."<init>":()V
   #2 = Long               30l
   #4 = Methodref          #9.#32         // com/pfroad/demo/FibonacciTest.fibonacci:(J)J
   #5 = Fieldref           #33.#34        // java/lang/System.out:Ljava/io/PrintStream;
   #6 = Methodref          #35.#36        // java/io/PrintStream.println:(J)V
   #7 = Long               2l
   #9 = Class              #37            // com/pfroad/demo/FibonacciTest
  #10 = Class              #38            // java/lang/Object
  #11 = Utf8               <init>
  #12 = Utf8               ()V
  #13 = Utf8               Code
  #14 = Utf8               LineNumberTable
  #15 = Utf8               LocalVariableTable
  #16 = Utf8               this
  #17 = Utf8               Lcom/pfroad/demo/FibonacciTest;
  #18 = Utf8               main
  #19 = Utf8               ([Ljava/lang/String;)V
  #20 = Utf8               args
  #21 = Utf8               [Ljava/lang/String;
  #22 = Utf8               x
  #23 = Utf8               J
  #24 = Utf8               fibonacci
  #25 = Utf8               (J)J
  #26 = Utf8               n
  #27 = Utf8               StackMapTable
  #28 = Utf8               <clinit>
  #29 = Utf8               SourceFile
  #30 = Utf8               FibonacciTest.java
  #31 = NameAndType        #11:#12        // "<init>":()V
  #32 = NameAndType        #24:#25        // fibonacci:(J)J
  #33 = Class              #39            // java/lang/System
  #34 = NameAndType        #40:#41        // out:Ljava/io/PrintStream;
  #35 = Class              #42            // java/io/PrintStream
  #36 = NameAndType        #43:#44        // println:(J)V
  #37 = Utf8               com/pfroad/demo/FibonacciTest
  #38 = Utf8               java/lang/Object
  #39 = Utf8               java/lang/System
  #40 = Utf8               out
  #41 = Utf8               Ljava/io/PrintStream;
  #42 = Utf8               java/io/PrintStream
  #43 = Utf8               println
  #44 = Utf8               (J)V
{
  public com.pfroad.demo.FibonacciTest();
    descriptor: ()V
    flags: ACC_PUBLIC
    Code:
      stack=1, locals=1, args_size=1
         0: aload_0
         1: invokespecial #1                  // Method java/lang/Object."<init>":()V
         4: return
      LineNumberTable:
        line 3: 0
      LocalVariableTable:
        Start  Length  Slot  Name   Signature
            0       5     0  this   Lcom/pfroad/demo/FibonacciTest;

  public static void main(java.lang.String[]);
    descriptor: ([Ljava/lang/String;)V
    flags: ACC_PUBLIC, ACC_STATIC
    Code:
      stack=3, locals=3, args_size=1
         0: ldc2_w        #2                  // long 30l
         3: invokestatic  #4                  // Method fibonacci:(J)J
         6: lstore_1
         7: getstatic     #5                  // Field java/lang/System.out:Ljava/io/PrintStream;
        10: lload_1
        11: invokevirtual #6                  // Method java/io/PrintStream.println:(J)V
        14: return
      LineNumberTable:
        line 9: 0
        line 10: 7
        line 11: 14
      LocalVariableTable:
        Start  Length  Slot  Name   Signature
            0      15     0  args   [Ljava/lang/String;
            7       8     1     x   J

  static {};
    descriptor: ()V
    flags: ACC_STATIC
    Code:
      stack=0, locals=0, args_size=0
         0: return
      LineNumberTable:
        line 6: 0
}

Java code:
static {
}

method: name -> <clinit>, descriptor -> ()V
*/
func (class *Class) GetClinitMethod() *Method {
	return class.getStaticMethod("<clinit>", "()V")
}
