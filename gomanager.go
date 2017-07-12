package main

import (
	"fmt"
	"os"
	"gomanager/command"
)

func main()  {
	paramMap := parseArgument(os.Args)
	cmdName := paramMap["command"]
	cmd := command.GetCommand(cmdName)
	cmd.Run(paramMap)

	//cmd := exec.Command("/bin/sh", "-c", "ps aux | grep node | grep -v 'grep' | awk '{print$2}'")
	//pidsByte, err := cmd.Output()
	//if err != nil {
	//	fmt.Println(err)
	//}
	//pids := string(pidsByte)
	//if pids == "" {
	//	fmt.Println("pid不存在")
	//	os.Exit(1)
	//}
	//fmt.Println(pids)
}



func parseArgument(params []string) map[string]string {
	paramMap := make(map[string]string)
	if len(params) == 1 {
		help()
	}
	paramMap["command"] = params[1]
	for i := 2; i < len(params); i+=2 {
		paramMap[params[i]] = params[i + 1]
	}
	return paramMap
}

func help()  {
	fmt.Println("Usage:", "gomanager command [arguments]")
	fmt.Println()
	fmt.Println("The commands are:")
	fmt.Println(fmt.Sprintf("\t%s\t%s", "start", "build the go file and start it"))
	os.Exit(0)
}
