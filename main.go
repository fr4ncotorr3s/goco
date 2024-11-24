package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

var goosList = []string{
	"linux", "windows", "darwin", "android", "freebsd", "netbsd", "openbsd",
	"ios", "js", "solaris", "illumos", "aix", "dragonfly", "hurd", "plan9", "zos",
}

var goarchList = []string{
	"amd64", "386", "arm", "arm64", "ppc64le", "ppc64", "riscv64",
	"mips", "mipsle", "mips64", "mips64le", "s390x", "loong64", "amd64p32", "wasm",
}

func main() {
	fmt.Println("Select a target operating system (GOOS):")
	goos := selectFromList(goosList)

	fmt.Println("\nSelect a target architecture (GOARCH):")
	goarch := selectFromList(goarchList)

	setEnv(goos, goarch)

	fmt.Println("\nTo set these manually, use the following commands:")
	if runtime.GOOS == "windows" {
		fmt.Printf("  $env:GOOS = \"%s\"\n", goos)
		fmt.Printf("  $env:GOARCH = \"%s\"\n", goarch)
	} else {
		fmt.Printf("  export GOOS=%s\n", goos)
		fmt.Printf("  export GOARCH=%s\n", goarch)
	}

	fmt.Println("\nBye.")
}

func setEnv(goos, goarch string) {
	if runtime.GOOS == "windows" {
		exec.Command("powershell", "-Command", fmt.Sprintf("$env:GOOS='%s'; $env:GOARCH='%s'", goos, goarch)).Run()
	} else {
		cmd := fmt.Sprintf("export GOOS=%s; export GOARCH=%s", goos, goarch)
		exec.Command("bash", "-c", cmd).Run()
	}

	fmt.Printf("\nEnvironment variables set in the current shell:\n")
	fmt.Printf("  GOOS   = %s\n", goos)
	fmt.Printf("  GOARCH = %s\n", goarch)
}

func selectFromList(options []string) string {
	reader := bufio.NewReader(os.Stdin)

	for {
		for i, option := range options {
			fmt.Printf("%d) %s\n", i+1, option)
		}

		fmt.Printf("Enter the number of your choice (1-%d): ", len(options))
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		index, err := parseIndex(input, len(options))
		if err != nil {
			fmt.Println("Invalid input. Please try again.")
			continue
		}

		return options[index]
	}
}

func parseIndex(input string, max int) (int, error) {
	index := -1
	_, err := fmt.Sscanf(input, "%d", &index)
	if err != nil || index < 1 || index > max {
		return -1, fmt.Errorf("invalid index")
	}
	return index - 1, nil
}
