package impl

import (
	"context"
	"testing"
)

func TestGitSvnImpl_Clone(t *testing.T) {
	ctl := NewGivSvnController()
	ctx := context.Background()
	ctl.Clone(ctx, "svn://47.92.104.229:3690/authsvn", "", "123456")
}
