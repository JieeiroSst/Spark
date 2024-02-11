package utility

import "Spark/modules"

type checkForm struct {
	Conn   string `json:"uuid" yaml:"uuid" form:"uuid"`
	Device string `json:"device" yaml:"device" form:"device"`
}

type onDevicePack struct {
	Code   int            `json:"code,omitempty"`
	Act    string         `json:"act,omitempty"`
	Msg    string         `json:"msg,omitempty"`
	Device modules.Device `json:"data"`
}

type checkUpdate struct {
	OS     string `form:"os" binding:"required"`
	Arch   string `form:"arch" binding:"required"`
	Commit string `form:"commit" binding:"required"`
}

type execDeviceCmd struct {
	Cmd  string `json:"cmd" yaml:"cmd" form:"cmd" binding:"required"`
	Args string `json:"args" yaml:"args" form:"args"`
}