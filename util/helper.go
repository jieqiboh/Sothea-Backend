package util

import (
	"os/exec"
	"strings"
)

// GetGitRoot gets the git root of this project.
func GetGitRoot() (string, error) {
	cmd := exec.Command("git", "rev-parse", "--show-toplevel")
	sb := strings.Builder{}
	cmd.Stdout = &sb
	err := cmd.Run()
	if err != nil {
		return "", err
	}
	str := sb.String()
	return str[:len(str)-1], nil
}

// MustGitPath is util for getting relative path to git root
func MustGitPath(relative string) string {
	ret, err := GetGitRoot()
	if err != nil {
		panic(err)
	}
	return ret + "/" + relative
}
