package utils

import (
	"os/exec"
)

func RunLinuxCmd(cmd string) (string, error) {
	bs, err := exec.Command("/bin/sh", "-c", cmd).Output()
	if err != nil {
		return "", err
	}
	return string(bs), nil
}
