package generate

type checkClient struct {
	OS     string `json:"os" yaml:"os" form:"os" binding:"required"`
	Arch   string `json:"arch" yaml:"arch" form:"arch" binding:"required"`
	Host   string `json:"host" yaml:"host" form:"host" binding:"required"`
	Port   uint16 `json:"port" yaml:"port" form:"port" binding:"required"`
	Path   string `json:"path" yaml:"path" form:"path" binding:"required"`
	Secure string `json:"secure" yaml:"secure" form:"secure"`
}

type generateClient struct {
	OS     string `json:"os" yaml:"os" form:"os" binding:"required"`
	Arch   string `json:"arch" yaml:"arch" form:"arch" binding:"required"`
	Host   string `json:"host" yaml:"host" form:"host" binding:"required"`
	Port   uint16 `json:"port" yaml:"port" form:"port" binding:"required"`
	Path   string `json:"path" yaml:"path" form:"path" binding:"required"`
	Secure string `json:"secure" yaml:"secure" form:"secure"`
}
