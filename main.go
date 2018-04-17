package main

// import "fmt"
// import "flag"
import "fmt"
import "jvm/classpath"
import "strings"

func startJVM(cmd *Cmd) {
	cp := classpath.Parse(cmd.XjreOption, cmd.cpOption)
	fmt.Printf("classpath: %s, class: %s args: %v \n", cmd.cpOption, cmd.class, cmd.args)
	className := strings.Replace(cmd.class, ".", "/", -1)
	data, _, err := cp.ReadClass(className)

	if err != nil {
		fmt.Printf("Cannot find or load main class %s\n", cmd.class)
		fmt.Printf("Error:%s", err)
		return
	}

	fmt.Printf("class data:%v\n", data)
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
