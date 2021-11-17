package process

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

const (
	_procDir = "/proc"
)

// getPIDs 获取所有pid
func getPIDs() []int {
	var pidS []int
	fs, err := ioutil.ReadDir(_procDir)
	if err != nil {
		return nil
	}
	for _, f := range fs {
		if pid, err := strconv.Atoi(f.Name()); err == nil {
			pidS = append(pidS, pid)
		}
	}
	return pidS
}

// getExeLink 获取程序启动路径
func getExePath(pid int) string {
	if exePath, err := os.Readlink(fmt.Sprintf("/proc/%d/exe", pid)); err != nil {
		return ""
	} else {
		return exePath
	}
}
