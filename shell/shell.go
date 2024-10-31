package shell

import (
	"os"
	"os/exec"
	"strings"
)

func execCmd(cmd string) *exec.Cmd {
	parts := strings.Split(cmd, " ")
	d := exec.Command(parts[0])
	if len(parts) > 1 {
		d = exec.Command(parts[0], parts[1:]...)
	}
	d.Stderr = os.Stdout
	return d
}

func BashStatusCode(cmd string) int {
	d := exec.Command("bash", "-c", cmd)
	d.Stderr = os.Stdout
	d.Stdout = os.Stdout
	err := d.Run()
	if err != nil {
		return err.(*exec.ExitError).ExitCode()
	}
	return 0
}

func BashExecOutput(cmd string) (string, error) {
	d := exec.Command("bash", "-c", cmd)
	d.Stderr = os.Stdout
	b, err := d.Output()
	return string(b), err
}

func ExecStatusCode(cmd string) int {
	d := execCmd(cmd)
	d.Stderr = os.Stderr
	d.Stdout = os.Stdout
	err := d.Run()
	if err != nil {
		return err.(*exec.ExitError).ExitCode()
	}
	return 0
}

func ExecOutput(cmd string) (string, error) {
	d := execCmd(cmd)
	b, err := d.Output()
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func ExecOutputVerbose(cmd string) error {
	d := execCmd(cmd)
	d.Stderr = os.Stderr
	d.Stdout = os.Stdout
	return d.Run()

}
