package chapter1_main

import (
	"flag"
	"fmt"
	"os"
)

type Cmd struct {
	helpFlag    bool
	versionFlag bool
	//classpath option
	cpOption    string
	XjreOption  string
	class       string
	args        []string
}

/**
	从命令行获取cmd
 */
func parseCmd() *Cmd {
	cmd := &Cmd{}
	flag.Usage = printUsage
	fmt.Printf("%s    ", &cmd.helpFlag)
	flag.BoolVar(&cmd.helpFlag, "help", false, "print help message")
	flag.BoolVar(&cmd.helpFlag, "?", false, "print help message")
	flag.BoolVar(&cmd.versionFlag, "version", false, "version info")
	flag.StringVar(&cmd.cpOption, "classpath", "", "class path")
	flag.StringVar(&cmd.cpOption, "cp", "", "equals classpath")
	flag.StringVar(&cmd.XjreOption, "Xjre", "", "path to jre")
	flag.Parse()

	args := flag.Args()
	if len(args) > 0 {
		cmd.class = args[0]
		cmd.args = args[1:]
	}

	return cmd
}

func printUsage() {
	fmt.Printf("Usage: %s [-options] class [args...] \n", os.Args[0])
}

func main() {
	cmd := parseCmd()
	if cmd.versionFlag {
		fmt.Println("version 0.0.1")
	} else if cmd.helpFlag || cmd.class == "" {
		printUsage()
	} else {
		startJVM(cmd)
	}
}

func startJVM(cmd *Cmd) {
	fmt.Printf("classpath:%s class:%s args:%v \n", cmd.cpOption, cmd.class, cmd.args)
}