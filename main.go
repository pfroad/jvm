package main

// import "fmt"
// import "flag"
import (
	"fmt"
	"jvm/classpath"
	"jvm/runtime/data"
	"strings"
)

func startJVM(cmd *Cmd) {
	cp := classpath.Parse(cmd.XjreOption, cmd.cpOption)
	classLoader := data.NewClassLoader(cp, cmd.verboseClassFlag)
	className := strings.Replace(cmd.class, ".", "/", -1)
	class := classLoader.LoadClass(className)
	//fmt.Printf("classpath: %s, class: %s args: %v \n", cmd.cpOption, cmd.class, cmd.args)

	//cf := loadClass(className, cp)
	mainMethod := class.GetMainMethod()

	if mainMethod != nil {
		interpret(mainMethod, cmd.verboseInstFlag)
	} else {
		fmt.Printf("Main method not found in class %s\n", cmd.class)
	}
	//data, _, err := cp.ReadClass(className)
	//
	//if err != nil {
	//	fmt.Printf("Cannot find or load main class %s\n", cmd.class)
	//	fmt.Printf("Error:%s", err)
	//	return
	//}
	//
	//fmt.Printf("class data:%v\n", data)
	//loadClass(className, cp)
	//frame := runtime.NewFrame(100, 100)
	//testLocalVars(frame.LocalVars())
	//testOperandStack(frame.OperandStack())
}

//func getMainMethod(cf *classfile.ClassFile) *classfile.MemberInfo {
//	for _, m := range cf.Methods() {
//		if m.Name() == "main" && m.Descriptor() == "([Ljava/lang/Name;)V" {
//			return m
//		}
//	}
//
//	return nil
//}

//func testLocalVars(vars runtime.LocalVars) {
//	vars.SetInt(0, 100)
//	vars.SetInt(1, -100)
//	vars.SetLong(2, 2997924580)
//	vars.SetLong(4, -2997924580)
//	vars.SetFloat(6, 3.1415926)
//	vars.SetDouble(7, 2.71828182845)
//	vars.SetRef(9, nil)
//
//	println(vars.GetInt(0))
//	println(vars.GetInt(1))
//	println(vars.GetLong(2))
//	println(vars.GetLong(4))
//	println(vars.GetFloat(6))
//	println(vars.GetDouble(7))
//	println(vars.GetRef(9))
//}
//
//func testOperandStack(stack *runtime.OperandStack) {
//	stack.PushInt(100)
//	stack.PushInt(-100)
//	stack.PushLong(2997924580)
//	stack.PushLong(-2997924580)
//	stack.PushFloat(3.1415926)
//	stack.PushDouble(2.71828182845)
//	stack.PushRef(nil)
//
//	println(stack.PopRef())
//	println(stack.PopDouble())
//	println(stack.PopFloat())
//	println(stack.PopLong())
//	println(stack.PopLong())
//	println(stack.PopInt())
//	println(stack.PopInt())
//}

//func loadClass(className string, cp *classpath.Classpath) *classfile.ClassFile {
//	data, _, err := cp.ReadClass(className)
//
//	if err != nil {
//		//fmt.Printf("Cannot find or load main class %s\n", cmd.class)
//		//fmt.Printf("Error:%s", err)
//		//return
//		panic(err)
//	}
//
//	cf, err := classfile.Parse(data)
//
//	if err != nil {
//		panic(err)
//	}
//
//	return cf
//}
//
//func printClassInfo(cf *classfile.ClassFile) {
//	fmt.Printf("version: %v.%v\n", cf.MajorVersion(), cf.MinorVeresion())
//	fmt.Printf("flags: 0X%x\n", cf.AccessFlags())
//	fmt.Printf("constant count: %d\n", len(cf.ConstantPool()))
//	fmt.Printf("this class: %v\n", cf.ClassName())
//	fmt.Printf("supre class: %v\n", cf.SuperClassName())
//	fmt.Printf("interfaces: %v\n", cf.InterfaceNames())
//
//	fmt.Printf("fields count: %d\n", len(cf.Fields()))
//	for _, field := range cf.Fields() {
//		fmt.Printf("%v\n", field.Name())
//	}
//
//	fmt.Printf("methods count: %d\n", len(cf.Methods()))
//	for _, method := range cf.Methods() {
//		fmt.Printf("%v\n", method.Name())
//	}
//}

func main() {
	cmd := parseCmd()
	if cmd.versionFlag {
		fmt.Println("1.0.0")
	} else if cmd.helpFlag || cmd.class == "" {
		printUsage()
	} else {
		startJVM(cmd)
	}
}
