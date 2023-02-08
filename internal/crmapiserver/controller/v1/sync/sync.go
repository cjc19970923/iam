package sync

import (
	"github.com/gin-gonic/gin"
	"github.com/marmotedu/component-base/pkg/core"
	"github.com/marmotedu/errors"
	"github.com/marmotedu/iam/internal/crmapiserver/model/params/sync"
	srvv1 "github.com/marmotedu/iam/internal/crmapiserver/service/v1"
	"github.com/marmotedu/iam/internal/crmapiserver/store"
	"github.com/marmotedu/iam/internal/pkg/code"
	"github.com/marmotedu/iam/pkg/log"
)

type SyncController struct {
	srv srvv1.Service
}

func NewSyncController(store store.Factory) *SyncController {
	return &SyncController{
		srv: srvv1.NewService(store),
	}
}

func (s *SyncController) SyncCusFieldMap(c *gin.Context) {
	log.L(c).Info("SyncCusFieldMap")
	//r := &sync.SyncParams{}
	//if err := c.ShouldBindJSON(r); err != nil {
	//	core.WriteResponse(c, errors.WithCode(code.ErrBind, ""), err.Error())
	//}
	r, ok := validationParams(c)
	if !ok {
		return
	}

}

func validationParams(c *gin.Context) (*sync.SyncParams, bool) {
	r := &sync.SyncParams{}
	if err := c.ShouldBindJSON(r); err != nil {
		core.WriteResponse(c, errors.WithCode(code.ErrBind, ""), err.Error())
		return nil, false
	}
	return r, true
}