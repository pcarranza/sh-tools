package main

import (
	"flag"
	"fmt"
	"net/url"
	"os"
	"path"
	"strings"
)

func main() {
	flag.Parse()

	if flag.NArg() == 0 {
		fmt.Printf("Not enough arguments, I need at least 1 url\n")
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
	u, err := url.Parse(gitURL)
	if err != nil {
		return fmt.Errorf("could not parse git url %s: %s", gitURL, err)
	}

	gopath := os.Getenv("GOPATH")
	if gopath == "" {
		gopath = path.Join(os.Getenv("HOME"), "Go")
	}

	repopath := u.Path
	if strings.HasSuffix(repopath, ".git") {
		repopath = repopath[:len(repopath)-4]
	}

	r := path.Join(gopath, "src", u.Hostname(), repopath)

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

	fmt.Printf("command to run: git clone %s %s", gitURL, r)

	// cmd := exec.Command("git", "clone", gitURL, r)
	// cmd.Stdout = os.Stdout
	// cmd.Stderr = os.Stderr
	// cmd.Stdin = os.Stdin
	// if err := cmd.Run(); err != nil {
	// 	return fmt.Errorf("could not clone git repo %s into %s: %s", gitURL, r, err)
	// }
	return nil
}

func dirExists(p string) bool {
	_, err := os.Stat(p)
	if err != nil {
		return false
	}
	return true
}
