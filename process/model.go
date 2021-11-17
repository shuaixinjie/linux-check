package process

// Proc 进程实体
type Proc struct {
	UID  string `json:"uid"`
	PID  int    `json:"pid"`
	PPID string `json:"ppid"`
	Time string `json:"time"`
	Exe  string `json:"exe"`
}
