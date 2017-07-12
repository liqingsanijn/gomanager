package command

import (
	"fmt"
	"os"
	"os/exec"
	"io/ioutil"
)

type start struct {

}

func (s *start) Help()  {
	fmt.Println("usage: start [arguments]")
	fmt.Println()
	fmt.Println("the arguments are:")
	fmt.Println(fmt.Sprintf("\t%s\t%s", "-n", "the project name"))
	fmt.Println(fmt.Sprintf("\t%s\t%s", "-p", "the go file path"))
	fmt.Println(fmt.Sprintf("\t%s\t%s", "--log", "the log file path"))
	os.Exit(0)
}

func (s *start) Run(paramMap map[string]string)  {
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
	fmt.Println(fmt.Sprintf("%s >> %s 2>&1 &", name, log))
	start := exec.Command("/bin/sh", "-c", fmt.Sprintf("~/go/bin/%s >> %s 2>&1 &", name, log))
	//start := exec.Command("/bin/sh", "-c", fmt.Sprintf("echo %s %s", name, log))
	startOut, err := start.StdoutPipe()
	checkErr(err)
	err = start.Start()
	checkErr(err)
	fmt.Println(fmt.Sprintf("start %s success", name))
	st, _ := ioutil.ReadAll(startOut)
	fmt.Println(st)
}
