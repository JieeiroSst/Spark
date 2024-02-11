package file

type removeDeviceFile struct {
	Files []string `json:"files" yaml:"files" form:"files" binding:"required"`
}

type listDeviceFile struct {
	Path string `json:"path" yaml:"path" form:"path" binding:"required"`
}

type getDeviceFile struct {
	Files   []string `json:"files" yaml:"files" form:"files" binding:"required"`
	Preview bool     `json:"preview" yaml:"preview" form:"preview"`
}

type deviceTextFile struct {
	File string `json:"file" yaml:"file" form:"file" binding:"required"`
}

type uploadToDevice struct {
	Path string `json:"path" yaml:"path" form:"path" binding:"required"`
	File string `json:"file" yaml:"file" form:"file" binding:"required"`
}
