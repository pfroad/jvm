package main

// import "fmt"
// import "flag"
import "fmt"

func startJVM(cmd *Cmd) {
	fmt.Printf("classpath: %s, class: %s args: %v \n", cmd.cpOption, cmd.class, cmd.args)
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
