package command

import (
	"fmt"
	"os"
)

func GetCommand(name string) command {
	switch name {
	case "help":
		return new(index)
	case "start":
		return new(start)
	case "restart":
		return new(restart)
	case "build":
		return new(build)
	default:
		return new(index)
	}
	return new(index)
}

type command interface {
	Help()
	Run(paramMap map[string]string)
}

type index struct {


}

func (i *index) Help()  {

}

func (i *index) Run(paramMap map[string]string)  {
	fmt.Println("Usage:", "gomanager command [arguments]")
	fmt.Println()
	fmt.Println("The commands are:")
	fmt.Println(fmt.Sprintf("\t%s\t%s", "start", "build the go file and start it"))
	os.Exit(0)
}

func checkErr(err error)  {
	if err != nil {
		fmt.Println(err.Error())
		fmt.Println("failed")
		os.Exit(1)
	}
}