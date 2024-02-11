package process

type killDeviceProcess struct {
	Pid int32 `json:"pid" yaml:"pid" form:"pid" binding:"required"`
}
