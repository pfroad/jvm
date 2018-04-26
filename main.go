package main

// import "fmt"
// import "flag"
import "fmt"
import "jvm/classpath"
import (
	"strings"
	"jvm/classfile"
)

func startJVM(cmd *Cmd) {
	cp := classpath.Parse(cmd.XjreOption, cmd.cpOption)
	fmt.Printf("classpath: %s, class: %s args: %v \n", cmd.cpOption, cmd.class, cmd.args)
	className := strings.Replace(cmd.class, ".", "/", -1)
	//data, _, err := cp.ReadClass(className)
	//
	//if err != nil {
	//	fmt.Printf("Cannot find or load main class %s\n", cmd.class)
	//	fmt.Printf("Error:%s", err)
	//	return
	//}
	//
	//fmt.Printf("class data:%v\n", data)
	loadClass(className, cp)
}

func loadClass(className string, cp *classpath.Classpath) {
	data, _, err := cp.ReadClass(className)

	if err != nil {
		//fmt.Printf("Cannot find or load main class %s\n", cmd.class)
		//fmt.Printf("Error:%s", err)
		//return
		panic(err)
	}

	cf, err := classfile.Parse(data)

	if err != nil {
		panic(err)
	}

	printClassInfo(cf)
}

func printClassInfo(cf *classfile.ClassFile) {
	fmt.Printf("version: %v.%v\n", cf.MajorVersion(), cf.MinorVeresion())
	fmt.Printf("flags: %v\n", cf.AccessFlags())
	fmt.Printf("constant count: %d\n", len(cf.ConstantPool()))
	fmt.Printf("this class: %v\n", cf.ClassName())
	fmt.Printf("supre class: %v\n", cf.SuperClassName())
	fmt.Printf("interfaces: %v\n", cf.InterfaceNames())

	fmt.Printf("fields count: %d\n", len(cf.Fields()))
	for _, field := range cf.Fields() {
		fmt.Printf("%v\n", field)
	}

	fmt.Printf("methods count: %d\n", len(cf.Methods()))
	for _, method := range cf.Methods() {
		fmt.Printf("%v\n", method)
	}
}

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
