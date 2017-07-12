package command

import (
	"fmt"
	"os"
	"os/exec"
)

type build struct {

}

func (s *build) Help()  {
	fmt.Println("usage: start [arguments]")
	fmt.Println()
	fmt.Println("the arguments are:")
	fmt.Println(fmt.Sprintf("\t%s\t%s", "-n", "the project name"))
	fmt.Println(fmt.Sprintf("\t%s\t%s", "-p", "the go file path"))
	os.Exit(0)
}

func (s *build) Run(paramMap map[string]string)  {
	name := paramMap["-n"]
	path := paramMap["-p"]
	if name == "" || path == "" {
		s.Help()
	}
	gopath := os.Getenv("GOPATH")
	cmd := exec.Command("/bin/sh", "-c", fmt.Sprintf("go build -o %s/bin/%s %s", gopath, name, path))
	outBytes, err := cmd.Output()
	checkErr(err)
	fmt.Println(string(outBytes))
}