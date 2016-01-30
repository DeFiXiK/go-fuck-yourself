package main

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/yamnikov-oleg/go-ps"
)

func RageFace(s string) {
	fmt.Println()
	fmt.Println("  (╯°□°）╯︵", Flip(s))
	fmt.Println()
}

func ShockFace(s string) {
	fmt.Println()
	fmt.Println("  (；￣Д￣) . o O( ", s, " )")
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

var (
	Chars   = []rune(" -_.abcdefghijklmnopqrstuvwxyz1234567890")
	Flipped = []rune(" -_'ɐqɔpǝɟɓɥıɾʞlɯuodbɹsʇnʌʍxʎz⇂zƐㄣϛ9ㄥ860")
)

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

func ParseArgs() (pname string, ok bool) {
	if len(os.Args) != 2 && len(os.Args) != 3 {
		return "", false
	}

	pname = os.Args[len(os.Args)-1]
	if pname == "you" {
		return "", false
	}

	return pname, true
}

var MultipleProcessesError = errors.New("found multiple processes")

func FindProcess(pattern string) (proc ps.Process, err error) {
	processes, err := ps.Processes()
	if err != nil {
		return nil, err
	}

	for _, p := range processes {
		if p.Pid() == os.Getpid() {
			continue
		}
		if !strings.Contains(p.Executable(), pattern) {
			continue
		}
		if proc != nil {
			return nil, MultipleProcessesError
		}
		proc = p
	}

	if proc == nil {
		err = fmt.Errorf("Process not found")
	}

	return
}

func FindAndKill(pname string) (string, error) {
	proc, err := FindProcess(pname)
	if err != nil {
		return "", err
	}

	osproc, err := os.FindProcess(proc.Pid())
	if err != nil {
		return "", err
	}
	osproc.Kill()

	return proc.Executable(), nil
}

const MinNameLen = 4

func main() {
	pname, ok := ParseArgs()
	if !ok {
		fmt.Println("Usage: fuck [you] <process-name>")
		os.Exit(1)
	}

	if len(pname) < MinNameLen {
		fmt.Printf("<process-name> can't be shorter than %v symbols\n", MinNameLen)
		os.Exit(1)
	}

	exec, err := FindAndKill(pname)
	if err == MultipleProcessesError {
		ShockFace("Not sure which one you mean to abuse...")
		os.Exit(1)
	}
	if err != nil {
		ShockFace("It didn't work out as expected...")
		os.Exit(1)
	}

	RageFace(exec)
}
