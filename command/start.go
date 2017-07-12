package command

import (
	"fmt"
	"os"
	"os/exec"
	"io/ioutil"
	"encoding/json"
	"bufio"
	"errors"
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
	fileName := "/tmp/gomanager/projects"
	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)
	//file, err := os.Create(fileName)
	checkErr(err)
	fileBytes, err := ioutil.ReadAll(file)
	checkErr(err)
	projects := []*project{}
	if string(fileBytes) != "" {
		err = json.Unmarshal(fileBytes, &projects)
	}
	if i, _ := getFromProjects(name, projects); i != -1 {
		checkErr(errors.New("name has already existed"))
	}
	checkErr(err)
	cmd := exec.Command("/bin/sh", "-c", fmt.Sprintf("go build -o %s/bin/%s %s", gopath, name, path))
	err = cmd.Run()
	checkErr(err)
	start := exec.Command("/bin/sh", "-c", fmt.Sprintf("%s >> %s 2>&1 &", name, log))
	//start := exec.Command("/bin/sh", "-c", fmt.Sprintf("echo %s %s", name, log))
	err = start.Run()
	checkErr(err)
	fmt.Println("std")
	p :=  new(project)
	p.Name = name
	p.State = 1
	p.Log = log
	p.Restart = 0
	p.Path = path
	projects = append(projects, p)
	jsonByte, err := json.Marshal(projects)
	write, err := os.Create(fileName)
	checkErr(err)
	w := bufio.NewWriter(write)
	w.Write(jsonByte)
	err = w.Flush()
	checkErr(err)
	fmt.Println(fmt.Sprintf("start %s success", name))

}

type project struct {
	Name string
	State int  //运行状态 1 running 0 stopped
	Log string
	Restart int
	Path string
}

func getFromProjects(name string, s []*project) (int, *project) {
	for i, v := range s {
		if name == v.Name {
			return i, v
		}
	}
	return -1, new(project)
}