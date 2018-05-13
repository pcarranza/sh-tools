package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path"

	"github.com/pcarranza/sh-tools/git"
)

func main() {
	flag.Parse()

	if flag.NArg() == 0 {
		fmt.Println("Error: not enough arguments, I need at least 1 url")
		fmt.Println("")
		fmt.Println("Usage:", os.Args[0], "git-url...")
		flag.PrintDefaults()
		os.Exit(1)
	}

	var returnCode int
	for _, arg := range flag.Args() {
		if err := cloneGitURL(arg); err != nil {
			fmt.Printf("Could not clone url %s: %s\n", arg, err)
			returnCode = 1
			continue
		}
	}
	os.Exit(returnCode)
}

func cloneGitURL(gitURL string) error {
	u, err := git.Parse(gitURL)
	if err != nil {
		return fmt.Errorf("could not parse git url %s: %s", gitURL, err)
	}

	gopath := os.Getenv("GOPATH")
	if gopath == "" {
		gopath = path.Join(os.Getenv("HOME"), "Go")
	}

	r := path.Join(gopath, "src", u.ToGoPath())

	if dirExists(r) {
		return fmt.Errorf("destination folder %s already exists", r)
	}

	d, _ := path.Split(r)
	if !dirExists(d) {
		fmt.Printf("creating path %s...\n", d)
		if err := os.MkdirAll(d, 755); err != nil {
			return fmt.Errorf("could not create destination folder %s: %s", d, err)
		}
	}

	cmd := exec.Command("git", "clone", gitURL, r)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("could not clone git repo %s into %s: %s", gitURL, r, err)
	}
	return nil
}

func dirExists(p string) bool {
	_, err := os.Stat(p)
	if err != nil {
		return false
	}
	return true
}
