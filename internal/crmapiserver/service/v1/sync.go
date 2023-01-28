package v1

import "github.com/marmotedu/iam/internal/crmapiserver/store"

type SyncSrv interface {
	SyncCusFieldMap() error
	SyncLinkFieldMap() error
	SyncCus() error
}

type syncService struct {
	store store.Factory
}

func newSyncService(srv *service) *syncService {
	return &syncService{
		store: srv.store,
	}
}

func (*syncService) SyncCusFieldMap() error {
	return nil
}

func (*syncService) SyncLinkFieldMap() error {
	return nil
}

func (*syncService) SyncCus() error {
	return nil
}
