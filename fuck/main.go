package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/mitchellh/go-ps"
)

const Usage = "Usage: fuck [you] <process-name>"

var (
	Chars   = []rune(" -_.abcdefghijklmnopqrstuvwxyz1234567890")
	Flipped = []rune(" -_'ɐqɔpǝɟɓɥıɾʞlɯuodbɹsʇnʌʍxʎz⇂zƐㄣϛ9ㄥ860")
)

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
	fmt.Println(Flip(fmt.Sprint(args...)))
	fmt.Println()
}
func ShokFace(args ...interface{}) {
	fmt.Println()
	fmt.Print("  (；￣Д￣) . o O( ")
	fmt.Print(args...)
	fmt.Println(" )")
	fmt.Println()
}

func IndexRune(s []rune, rn rune) int {
	for i, v := range s {
		if v == rn {
			return i
		}
	}
	return -1
}

func ReverseRunes(r []rune) {
	for i := 0; i < len(r)/2; i++ {
		j := len(r) - i - 1
		r[i], r[j] = r[j], r[i]
	}
}

func Flip(str string) string {
	str = strings.ToLower(str)

	buf := make([]rune, 0, len(str))
	for _, srcRune := range str {
		srcIndex := IndexRune(Chars, srcRune)
		if srcIndex < 0 {
			continue
		}
		dstRune := Flipped[srcIndex]
		buf = append(buf, dstRune)
	}
	ReverseRunes(buf)

	return string(buf)
}
