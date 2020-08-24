package impl

import (
	"context"
	"fmt"
	"github.com/fatelei/grpc-svn/config"
	"github.com/fatelei/grpc-svn/pkg/controller"
	"github.com/fatelei/grpc-svn/pkg/utils"
	"io/ioutil"
	"os"
	"strings"
)

type GitSvnImpl struct{}

var _ controller.GitSvn = (*GitSvnImpl)(nil)

func NewGivSvnController() *GitSvnImpl {
	return &GitSvnImpl{}
}

func (p *GitSvnImpl) Clone(ctx context.Context, svnUrl string, mailSuffix string) error {
	var out *string
	var err error
	err = p.genAuthorsForRepo(svnUrl, mailSuffix)
	if err != nil {
		return err
	}
	authorFileName := utils.GetAuthorsFileName(utils.GetRepoName(svnUrl))
	path := fmt.Sprintf("%s/%s", config.DefaultConfig.DestPath, authorFileName)
	out, err = utils.ExecCommand("git", "svn", "clone", "-s", fmt.Sprintf("--author-files=%s", path), svnUrl)
	if err != nil {
		return err
	}
	fmt.Printf("output: %s\n", *out)
	return nil
}

func (p *GitSvnImpl) Update(ctx context.Context, svnUrl string) error {
	path := p.generatePath(svnUrl)
	_, err := p.checkPath(path)
	if err != nil {
		return err
	}
	out, err := utils.ExecCommand(fmt.Sprintf("current_dir=$PWD;cd %s;git svn rebase;cd $current_dir;", path))
	if err != nil {
		return err
	}
	fmt.Printf("update response: %s", *out)
	return nil
}


func (p *GitSvnImpl) genAuthorsForRepo(svnUrl string, mailSuffix string) error {
	out, err := utils.ExecCommand(fmt.Sprintf(`svn log -q -r 1:HEAD %s | grep '^r' | awk -F'|' '!x[$2]++{print$2}'`, svnUrl))
	if err != nil {
		fmt.Errorf("gen authors failed: %+v\n", err)
		return err
	}

	rawLogs := strings.Split(*out, "\n")
	contents := make([]string, 0)
	for _, ele := range rawLogs {
		item := strings.TrimSpace(ele)
		if len(item) > 0 {
			contents = append(contents, fmt.Sprintf("%s = %s <%s@%s>", item, item, item, mailSuffix))
		}
	}

	authorFileName := utils.GetAuthorsFileName(utils.GetRepoName(svnUrl))
	path := fmt.Sprintf("%s/%s", config.DefaultConfig.DestPath, authorFileName)
	isValidPath, err := p.checkPath(path)
	if isValidPath {
		content := []byte(strings.Join(contents, "\n"))
		err = ioutil.WriteFile(path, content, 0644)
		if err != nil {
			fmt.Errorf("gen authors files failed: %+v\n", err)
			return err
		}
		return nil
	}
	return err
}

func (p *GitSvnImpl) generatePath(svnUrl string) string {
	repoName := utils.GetRepoName(svnUrl)
	return fmt.Sprintf("%s/%s", config.DefaultConfig.DestPath, repoName)
}


func (p *GitSvnImpl) checkPath(path string) (bool, error) {
	if _, err := os.Stat(path); err != nil {
		return false, err
	}
	return true, nil
}
