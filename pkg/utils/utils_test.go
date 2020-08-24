package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetRepoName(t *testing.T) {
	name := GetRepoName("file:////Users/leo/work/merico/play/svn/play-ground/svnrepo/project1")
	assert.Equal(t, "project1", name)
}

func TestGetAuthorsFileName(t *testing.T) {
	name := GetAuthorsFileName("project1")
	assert.Equal(t, "project1-authors.txt", name)
}

func TestExecCommand(t *testing.T) {
	output, _ := ExecCommand("ls")
	assert.True(t, len(*output) > 0)

	_, err := ExecCommand("this is a wrong command")
	assert.True(t, err != nil)
}