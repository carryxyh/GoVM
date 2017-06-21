package main

import (
	"flag"
	"fmt"
	"os"
	"GoVM/chapter2-class/classpath"
	"strings"
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
	cp := classpath.Parse(cmd.XjreOption, cmd.cpOption)
	fmt.Printf("classpath: %+v class: %v args:%v  \n", cp, cmd.class, cmd.args)
	className := strings.Replace(cmd.class, ".", "/", -1)
	classData, _, err := cp.ReadClass(className)
	if err != nil {
		fmt.Printf("Could not find or load main class %s \n", cmd.class)
		return
	}

	fmt.Printf("class data:%v  \n", classData)
}