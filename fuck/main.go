package main

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/yamnikov-oleg/go-ps"
)

// RageFace prints text flipped by an outraged smiley.
func RageFace(s string) {
	fmt.Println()
	fmt.Println("  (╯°□°）╯︵", Flip(s))
	fmt.Println()
}

// ShockFace prints text in a mind bubble of a regretful smiley.
func ShockFace(s string) {
	fmt.Println()
	fmt.Println("  (；￣Д￣) . o O( ", s, " )")
	fmt.Println()
}

// IndexRune returns index by which rn is found in s.
// If s doesn't contain rn, returns -1.
func IndexRune(s []rune, rn rune) int {
	for i, v := range s {
		if v == rn {
			return i
		}
	}
	return -1
}

// ReverseRunes reverses runes slice.
func ReverseRunes(r []rune) {
	for i := 0; i < len(r)/2; i++ {
		j := len(r) - i - 1
		r[i], r[j] = r[j], r[i]
	}
}

// Regular-to-flipped char map.
// Each i'th rune at Chars has it's flipped version at Flipped[i]
var (
	Chars   = []rune(" -_.abcdefghijklmnopqrstuvwxyz1234567890")
	Flipped = []rune(" -_'ɐqɔpǝɟɓɥıɾʞlɯuodbɹsʇnʌʍxʎz⇂zƐㄣϛ9ㄥ860")
)

// Flip flips string around, replacing each character
// as if it was rotated by 180 degrees.
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

// ParseArgs reads application arguements for process name.
// If arguements are malformed, it returns empty string and false.
// Otherwise it returns the process name and true.
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

// ErrMultipleProcesses happens when FindProcess function finds several processes
// per pattern.
var ErrMultipleProcesses = errors.New("found multiple processes")

// FindProcess scans list of processes running on the system, looking for the one
// which's name contains pattern.
// It succeeds only if there was found exactly one process.
// Otherwise it retunes nil proc and appopriate error.
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
			return nil, ErrMultipleProcesses
		}
		proc = p
	}

	if proc == nil {
		err = fmt.Errorf("Process not found")
	}

	return
}

// FindAndKill kills the system process, which's name contains pname.
// It uses FindProcess() to find the process and then sends kill signal.
// On success it returnes an actual name of a victim.
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

// MinNameLen evaluates to minimal supported length
// of process name pattern string
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
	if err == ErrMultipleProcesses {
		ShockFace("Not sure which one you mean to abuse...")
		os.Exit(1)
	}
	if err != nil {
		ShockFace("It didn't work out as expected...")
		os.Exit(1)
	}

	RageFace(exec)
}
