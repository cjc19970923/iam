package sync

import (
	"github.com/gin-gonic/gin"
	srvv1 "github.com/marmotedu/iam/internal/crmapiserver/service/v1"
	"github.com/marmotedu/iam/internal/crmapiserver/store"
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

}
