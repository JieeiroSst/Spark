package bridge

import (
	"Spark/utils/cmap"
	"sync"

	"github.com/gin-gonic/gin"
)

type Bridge struct {
	creation int64
	using    bool
	uuid     string
	lock     *sync.Mutex
	Dst      *gin.Context
	Src      *gin.Context
	ext      any
	OnPull   func(bridge *Bridge)
	OnPush   func(bridge *Bridge)
	OnFinish func(bridge *Bridge)
}

var bridges = cmap.New[*Bridge]()
