package utils

import (
	"fmt"
	"os/exec"
	"strings"
)

func GetRepoName(url string) string {
	index := strings.LastIndex(url, "/")
	return url[index+1:len(url)]
}


func GetAuthorsFileName(repoName string) string {
	return fmt.Sprintf("%s-authors.txt", repoName)
}


func ExecCommand(command string, arg ...string) (*string, error) {
	out, err := exec.Command(command, arg...).CombinedOutput()
	if err != nil {
		return nil, err
	}
	humanOutput := string(out)
	return &humanOutput, nil
}
