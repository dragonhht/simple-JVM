package main

import (
	"flag"
	"fmt"
	"os"
)

/*
 * 定义命令的数据结构
 */
type Cmd struct {
	// 帮助参数
	helpFlag bool
	// 版本参数
	versionFlag bool
	// -classpath
	cpOption string
	class string
	// jre目录
	XjreOption string
	args []string
}

/*
 * 解析命令
 */
func parseCmd() *Cmd {
	cmd := &Cmd{}

	flag.Usage = printUsage
	flag.BoolVar(&cmd.helpFlag, "help", false, "print hellp message")
	flag.BoolVar(&cmd.helpFlag, "?", false, "print help message")
	flag.BoolVar(&cmd.versionFlag, "version", false, "print version and exit")
	flag.StringVar(&cmd.cpOption, "classPath", "", "classpath")
	flag.StringVar(&cmd.cpOption, "cp", "", "classpath")
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
	fmt.Printf("Usage: %s [-option] class [args...]\n", os.Args[0])
}

