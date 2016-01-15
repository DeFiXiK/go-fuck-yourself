package main

import (
	"fmt"
	"os"

	"github.com/mitchellh/go-ps"
)

const Usage = "Usage: fuck [you] <process-name>"
const Chars = " -_.abcdefghijklmnopqrstuvwxyz1234567890"
const Flipped = " -_'ɐqɔpǝɟɓɥıɾʞlɯuodbɹsʇnʌʍxʎz⇂zƐㄣϛ9ㄥ860"

var ProcessName string

func main() {
	if len(os.Args) < 2 || len(os.Args) > 3 {
		Fatal(Usage)
	}
	ProcessName = os.Args[len(os.Args)-1]
	if ProcessName == "you" {
		Fatal(Usage)
	}
	pid, err := FindProcess(ProcessName)
	if err != nil {
		Fatal(err)
	}
	err = KillerProcess(pid)
	if err != nil {
		Fatal(err)
	}
	RageFace(ProcessName)
	ShokFace()
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

func KillerProcess(pid int) (err error) {
	osproc, err := os.FindProcess(pid)
	if err != nil {
		return err
	}
	osproc.Kill()
	return nil
}

func Fatal(err interface{}) {
	fmt.Println(err)
	os.Exit(1)
}

func RageFace(args ...interface{}) {
	fmt.Println()
	fmt.Print("  (╯°□°）╯︵")
	fmt.Println(args...)
	fmt.Println()
}
func ShokFace(args ...interface{}) {
	fmt.Println()
	fmt.Print("  (；￣Д￣) . o O( ")
	fmt.Print(args...)
	fmt.Println(" )")
	fmt.Println()
}

func Flip(str string) {

}
