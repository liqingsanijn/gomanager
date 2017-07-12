package command

import (
	"fmt"
	"os"
	"os/exec"
)

type restart struct {

}

func (s *restart) Help()  {
	fmt.Println("usage: start [arguments]")
	fmt.Println()
	fmt.Println("the arguments are:")
	fmt.Println(fmt.Sprintf("\t%s\t%s", "-n", "the project name"))
	fmt.Println(fmt.Sprintf("\t%s\t%s", "-p", "the go file path"))
	fmt.Println(fmt.Sprintf("\t%s\t%s", "--log", "the log file path"))
	os.Exit(0)
}

func (s *restart) Run(paramMap map[string]string)  {
	name := paramMap["-n"]
	path := paramMap["-p"]
	log := paramMap["--log"]
	if name == "" || path == "" {
		s.Help()
	}
	if log == "" {
		log = "/data/log/" + name
	}
	gopath := os.Getenv("GOPATH")
	cmd := exec.Command("/bin/sh", "-c", fmt.Sprintf("go build -o %s/bin/%s %s", gopath, name, path))
	outBytes, err := cmd.Output()
	checkErr(err)
	fmt.Println(string(outBytes))
	start := exec.Command("/bin/sh", "-c", fmt.Sprintf("%s >> %s 2>&1 &", name, log))
	startOut, err := start.Output()
	checkErr(err)
	fmt.Println(startOut)
}