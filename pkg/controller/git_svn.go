package controller

import "context"

type GitSvn interface {
	Clone(ctx context.Context, svnUrl string, mailSuffix string) error
	Update(ctx context.Context, svnUrl string) error
}
