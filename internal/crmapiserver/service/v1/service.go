package v1

import "github.com/marmotedu/iam/internal/crmapiserver/store"

type Service interface {
	Nodify() NodifySrv
	Sync() SyncSrv
}

type service struct {
	store store.Factory
}

func NewService(store store.Factory) Service {
	return &service{
		store: store,
	}
}

func (s *service) Nodify() NodifySrv {
	return newNodifyService(s)
}

func (s *service) Sync() SyncSrv {
	return newSyncService(s)
}
