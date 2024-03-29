package main

import (
	"os"
	"os/exec"
	"path"
	"strings"
	"syscall"
)

const (
	baseDir         = "src"
	lowerCaseOption = false
)

func usage() {
	println("Usage git get <URL>")
}

func main() {
	if len(os.Args) < 2 {
		usage()
		os.Exit(1)
	}

	gitUrl := os.Args[1]

	if len(gitUrl) == 0 {
		usage()
		os.Exit(1)
	}

	h, p, err := parseUrl(gitUrl)
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}

	if lowerCaseOption {
		h = strings.ToLower(h)
		p = strings.ToLower(p)
	}

	fullPath := path.Join(rootDirectory(), h, p)

	cloneErr := gitClone(gitUrl, fullPath)
	if cloneErr != nil {
		println(cloneErr)
		os.Exit(254)
	}
}

func rootDirectory() string {
	return path.Join(os.Getenv("HOME"), baseDir)
}

func gitClone(url string, path string) error {
	gitBin, lookErr := exec.LookPath("git")
	if lookErr != nil {
		return lookErr
	}

	args := []string{"git", "clone", url, path}
	env := os.Environ()

	err := syscall.Exec(gitBin, args, env)

	if err != nil {
		os.Exit(1)
	}

	return nil
}
