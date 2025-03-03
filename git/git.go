package git

import (
	"fmt"
	"log"
	"strings"

	"github.com/de1ux/gitstuff/shell"
	"github.com/de1ux/gitstuff/stack"
)

func cleanBranchName(branch string) string {
	branch = strings.Replace(branch, "refs/heads/", "", -1)
	branch = strings.Replace(branch, "\n", "", -1)
	return branch
}

func UntrackedFiles() (bool, error) {
	out, err := shell.ExecOutput("git status -s --untracked-files=no")
	if err != nil {
		return false, err
	}
	return out != "", nil
}

func IsBranch(branch string) (bool, error) {
	out, err := shell.ExecOutput("git symbolic-ref -q HEAD")
	if err != nil {
		return false, err
	}
	return cleanBranchName(out) == branch, nil
}

func Push(branch string, force bool) error {
	currentBranch := cleanBranchName(branch)

	cmd := "git push origin " + currentBranch
	if force {
		cmd += " --force-with-lease"
	}
	_, err := shell.ExecOutput(cmd)
	return err
}

func CommitNoEdit() error {
	code := shell.BashStatusCode("git commit -a --no-edit")
	if code == 1 {
		log.Print("No changes, skipping commit")
		return nil
	}
	if code == 0 {
		return nil
	}
	return fmt.Errorf("failed: commit --no-edit returned exit code %d", code)
}

func Commit(message string) error {
	code := shell.BashStatusCode("git commit -am \"" + message + "\"")
	if code == 1 {
		log.Print("No changes, skipping commit")
		return nil
	}
	if code == 0 {
		return nil
	}
	return fmt.Errorf("failed: commit returned exit code %d", code)
}

func CurrentBranch() (string, error) {
	out, err := shell.ExecOutput("git symbolic-ref HEAD")
	if err != nil {
		out, err = shell.ExecOutput("git rev-parse --short HEAD")
		if err != nil {
			return "", err
		}
	}
	return cleanBranchName(out), nil
}

func DeleteBranch(branch string) error {
	_, err := shell.ExecOutput("git branch -D " + branch)
	return err
}

func Checkout(branch string) error {
	_, err := shell.ExecOutput("git checkout " + branch)
	if err != nil {
		return err
	}
	s, err := stack.Load()
	if err != nil {
		return err
	}
	defer s.Save()

	s.Push(branch)
	return nil
}

func CheckoutNew(branch string) error {
	_, err := shell.ExecOutput("git checkout -b " + branch)
	if err != nil {
		return err
	}
	s, err := stack.Load()
	if err != nil {
		return err
	}
	defer s.Save()

	s.Push(branch)
	return nil
}

func InMergeConflict() bool {
	out, err := shell.BashExecOutput("git diff --diff-filter=U --name-only | wc -l")
	if err != nil {
		panic("failed to detect if in a merge conflict: " + err.Error())
	}
	if strings.TrimSpace(out) == "0" {
		return false
	}
	return true
}

func TicketNumber(prefix, branch string) string {
	prefixLess := strings.Replace(branch, prefix+"/", "", 1)
	parts := strings.Split(prefixLess, "-")
	if len(parts) == 1 {
		return ""
	}
	return parts[0] + "-" + parts[1]
}

func ResetLastCommit() error {
	_, err := shell.ExecOutput("git reset HEAD~1")
	return err
}

func Pull(branch string) error {
	var err error
	if branch == "" {
		branch, err = CurrentBranch()
		if err != nil {
			return err
		}
	}
	err = shell.ExecOutputVerbose("git pull origin " + branch)
	return err
}

func BranchExists(branch string) (bool, error) {
	code := shell.ExecStatusCode("git rev-parse --verify " + branch)
	if code != 0 {
		return false, nil
	}
	return true, nil
}

func CurrentOrgAndRepo() (string, string, error) {
	out, err := shell.ExecOutput("git config --get remote.origin.url")
	if err != nil {
		return "", "", err
	}
	parts := strings.Split(out, "/")
	if len(parts) < 2 {
		return "", "", fmt.Errorf("failed to get repo name from url: %s", out)
	}
	orgName := parts[len(parts)-2]
	orgName = strings.Replace(orgName, "\n", "", -1)

	repoName := strings.TrimSuffix(parts[len(parts)-1], ".git")
	repoName = strings.Replace(repoName, "\n", "", -1)
	return orgName, repoName, nil
}
