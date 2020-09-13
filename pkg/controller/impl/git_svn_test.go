package impl

import (
	"context"
	"github.com/fatelei/grpc-svn/pkg/utils"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestGitSvnImpl_Clone(t *testing.T) {
	ctl := NewGivSvnController()
	ctx := context.Background()
	svnUrl := "svn://47.92.104.229:3690/authsvn"
	ctl.CleanRepoFromLocal(svnUrl)
	ctl.Clone(ctx, svnUrl, "", "admin", "123456")
	repoName := utils.GetRepoName(svnUrl)
	repoPath := ctl.GeneratePath(repoName)
	_, err := os.Stat(repoPath)
	assert.True(t, !os.IsNotExist(err))
}
