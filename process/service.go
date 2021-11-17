package process

import "sort"

type ProcSvc struct {
}

// NewProcSvc new
func NewProcSvc() *ProcSvc {
	return new(ProcSvc)
}

// Get 获取进程信息
func (p *ProcSvc) Get() []Proc {
	var procS []Proc
	pidS := getPIDs()
	for _, pid := range pidS {
		exePath := getExePath(pid)
		procS = append(procS, Proc{
			UID:  "",
			PID:  pid,
			PPID: "",
			Time: "",
			Exe:  exePath,
		})
	}
	sort.Sort(procSort(procS))
	return procS
}

// procSort 实现一个排序，之后可以丰富成支持各种排序
type procSort []Proc

func (p procSort) Len() int {
	return len(p)
}
func (p procSort) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}
func (p procSort) Less(i, j int) bool {
	return p[i].PID < p[j].PID
}
