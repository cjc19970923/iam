package sync

import (
	srvv1 "github.com/marmotedu/iam/internal/crmapiserver/service/v1"
	"github.com/marmotedu/iam/internal/crmapiserver/store"
)

type SyncController struct {
	srv srvv1.Service
}

func NewSyncController(store store.Factory) *SyncController {
	return &SyncController{
		srv: srvv1.NewService(store),
	}
}

func (s *SyncController) SyncCusFieldMap() {

}