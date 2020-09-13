package impl

import (
	"bytes"
	"context"
	"fmt"
	pipe "github.com/b4b4r07/go-pipe"
	"github.com/fatelei/grpc-svn/config"
	"github.com/fatelei/grpc-svn/pkg/controller"
	"github.com/fatelei/grpc-svn/pkg/utils"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
)

type GitSvnImpl struct{}

var _ controller.GitSvn = (*GitSvnImpl)(nil)

func NewGivSvnController() *GitSvnImpl {
	return &GitSvnImpl{}
}

func (p *GitSvnImpl) Clone(ctx context.Context, svnUrl string, mailSuffix string, username string, password string) error {
	var out string
	var err error
	var b bytes.Buffer
	var authorFilePath string
	authorFilePath, err = p.genAuthorsForRepo(svnUrl, mailSuffix, username, password)
	if err != nil {
		panic(err)
		return err
	}
	commands := make([]*exec.Cmd, 0)
	if len(password) > 0 {
		commands = append(commands, exec.Command("echo", password))
	}
	commands = append(commands, exec.Command(
		"git", "svn", "clone", svnUrl, "-s", fmt.Sprintf("--authors-file=%s", authorFilePath), "--username", username))
	if err = pipe.Command(&b, commands...); err != nil {
		log.Fatal(err)
	}

	out = b.String()
	fmt.Printf("output: %s\n", out)
	return nil
}

func (p *GitSvnImpl) Update(ctx context.Context, svnUrl string) error {
	repoName := utils.GetRepoName(svnUrl)
	path := p.GeneratePath(repoName)
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


func (p *GitSvnImpl) genAuthorsForRepo(svnUrl string, mailSuffix string, username string, password string) (path string, err error) {
	var b bytes.Buffer
	commands := make([]*exec.Cmd, 3, 3)
	if len(username) > 0 && len(password) > 0 {
		commands[0] = exec.Command(
			"svn", "log", "--non-interactive", "--username", username, "--password", password, "-q", "-r", "1:HEAD", svnUrl)
	} else {
		commands[0] = exec.Command("svn", "log", "-q", "-r", "1:HEAD", svnUrl)
	}
	commands[1] = exec.Command("grep", `^r`)
	commands[2] = exec.Command("awk","-F", `|`, `{print $2}`)

	if err = pipe.Command(&b, commands...); err != nil {
		log.Fatal(err)
	}

	out := b.String()
	rawLogs := strings.Split(out, "\n")
	contents := make([]string, 0)
	for _, ele := range rawLogs {
		item := strings.TrimSpace(ele)
		if len(item) > 0 {
			if len(mailSuffix) == 0 {
				contents = append(contents, fmt.Sprintf("%s = %s <>", item, item))
			} else {
				contents = append(contents, fmt.Sprintf("%s = %s <%s@%s>", item, item, item, mailSuffix))
			}
		}
	}
	authorFileName := utils.GetAuthorsFileName(utils.GetRepoName(svnUrl))
	path = p.GeneratePath(authorFileName)
	isValidPath, err := p.checkPath(path)
	if !isValidPath {
		content := []byte(strings.Join(contents, "\n"))
		err = ioutil.WriteFile(path, content, 0644)
		if err != nil {
			fmt.Errorf("gen authors files failed: %+v\n", err)
			return
		}
		return
	}
	return
}

func (p *GitSvnImpl) GeneratePath(filename string) string {
	var path string
	if len(config.DefaultConfig.DestPath) > 0 {
		path = fmt.Sprintf("%s/%s", config.DefaultConfig.DestPath, filename)
	} else {
		cwd, _ := os.Getwd()
		path = fmt.Sprintf("%s/%s", cwd, filename)
	}
	return path
}

func (p *GitSvnImpl) CleanRepoFromLocal(svnUrl string) {
	repoName := utils.GetRepoName(svnUrl)
	path := p.GeneratePath(repoName)
	if err := os.RemoveAll(path); err != nil {
		fmt.Printf("remove %s error %v\n", svnUrl, err)
	}
}

func (p *GitSvnImpl) checkPath(path string) (bool, error) {
	if _, err := os.Stat(path); err != nil {
		return false, err
	}
	return true, nil
}
