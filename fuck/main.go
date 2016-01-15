package main

import (
	"fmt"
	"os"

	"github.com/mitchellh/go-ps"
)

const Usage = "Usage: fuck [you] <process-name>"

var ProcessName string

func main() {
	if len(os.Args) < 2 || len(os.Args) > 3 {
		fmt.Println(Usage)
		os.Exit(1)
	}
	ProcessName = os.Args[len(os.Args)-1]
	if ProcessName == "you" {
		fmt.Println(Usage)
		os.Exit(1)
	}
	pid, err := FindProcess(ProcessName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("ID процесса: %v\n", pid)
}

func FindProcess(pattern string) (pid int, err error) {
	processes, err := ps.Processes()
	if err != nil {
		return 0, err
	}

	for _, p := range processes {
		if pattern == p.Executable() && os.Getpid() != p.Pid() {
			return p.Pid(), nil
		}
	}
	return 0, fmt.Errorf("Process not found")
}
