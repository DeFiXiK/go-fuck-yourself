package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/mitchellh/go-ps"
)

func RageFace(args ...interface{}) {
	fmt.Println()
	fmt.Print("  (╯°□°）╯︵")
	fmt.Println(Flip(fmt.Sprint(args...)))
	fmt.Println()
}

func ShockFace(args ...interface{}) {
	fmt.Println()
	fmt.Print("  (；￣Д￣) . o O( ")
	fmt.Print("It didn't work out as expected...")
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

func FindProcess(pattern string) (pid int, err error) {
	processes, err := ps.Processes()
	if err != nil {
		return 0, err
	}

	for _, p := range processes {
		if p.Pid() == os.Getpid() {
			continue
		}
		if !strings.Contains(p.Executable(), pattern) {
			continue
		}
		return p.Pid(), nil
	}

	return 0, fmt.Errorf("Process not found")
}

func FindAndKill(pname string) error {
	pid, err := FindProcess(pname)
	if err != nil {
		return err
	}

	osproc, err := os.FindProcess(pid)
	if err != nil {
		return err
	}
	osproc.Kill()

	return nil
}

func main() {
	pname, ok := ParseArgs()
	if !ok {
		fmt.Println("Usage: fuck [you] <process-name>")
		os.Exit(1)
	}

	if err := FindAndKill(pname); err != nil {
		ShockFace()
		os.Exit(1)
	}

	RageFace(pname)
}
