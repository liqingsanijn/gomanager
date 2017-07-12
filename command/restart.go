package command

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"io/ioutil"
	"errors"
	"encoding/json"
	"bufio"
)

type restart struct {

}

func (s *restart) Help()  {
	fmt.Println("usage: start [arguments]")
	fmt.Println()
	fmt.Println("the arguments are:")
	fmt.Println(fmt.Sprintf("\t%s\t%s", "-n", "the project name"))
	os.Exit(0)
}

func (s *restart) Run(paramMap map[string]string)  {
	name := paramMap["-n"]
	if name == "" {
		s.Help()
	}
	fileName := "/tmp/gomanager/projects"
	file, err:= os.Open(fileName)
	checkErr(err)
	fileBytes, err := ioutil.ReadAll(file)
	checkErr(err)
	fileContent := string(fileBytes)
	if fileContent == "" {
		checkErr(errors.New("no project existed"))
	}
	projects := []*project{}
	err = json.Unmarshal(fileBytes, &projects)
	checkErr(err)
	i, p := getFromProjects(name, projects)
	if p.Name == "" {
		checkErr(errors.New("project is not existed"))
	}
	fmt.Println(fmt.Sprintf( "ps aux | grep %s | grep -v 'grep' | awk '{print$2}'", name))
	cmd := exec.Command("/bin/sh", "-c", fmt.Sprintf( "ps aux | grep %s | grep -v '[grep|restart]' | awk '{print$2}'", name))
	outBytes, err := cmd.Output()
	checkErr(err)
	pidStr := string(outBytes)
	if pidStr != "" {
		pids := strings.Split(pidStr, "\n")
		fmt.Println(pids)
		str := ""
		for _, pid := range pids {
			str += pid + " "
		}
		fmt.Println("kill", str)
		kill := exec.Command("/bin/sh", "-c", fmt.Sprintf("kill %s", str))
		err := kill.Run()
		checkErr(err)
	}
	gopath := os.Getenv("GOPATH")
	bu := exec.Command("/bin/sh", "-c", fmt.Sprintf("go build -o %s/bin/%s %s", gopath, name, p.Path))
	err = bu.Run()
	checkErr(err)
	start := exec.Command("/bin/sh", "-c", fmt.Sprintf("%s >> %s 2>&1 &", name, p.Log))
	err = start.Run()
	checkErr(err)
	p.Restart = p.Restart + 1
	projects[i] = p
	jsonByte, err := json.Marshal(projects)
	write, err := os.Create(fileName)
	checkErr(err)
	w := bufio.NewWriter(write)
	w.Write(jsonByte)
	err = w.Flush()
	checkErr(err)
	fmt.Println(fmt.Sprintf("restart %s success", name))
}