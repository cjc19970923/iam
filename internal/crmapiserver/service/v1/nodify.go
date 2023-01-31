package v1

import (
	"github.com/marmotedu/errors"
	"github.com/marmotedu/iam/internal/crmapiserver/store"
	"github.com/marmotedu/iam/internal/crmapiserver/store/crm"
	"github.com/marmotedu/iam/internal/pkg/code"
)

type NodifySrv interface {
	CusNodify(params crm.CrmStore) error
	SetCrmStore(store crm.CrmStore)
}

type nodifyService struct {
	store    store.Factory
	crmStore crm.CrmStore
}

func newNodifyService(srv *service) *nodifyService {
	return &nodifyService{
		store: srv.store,
	}
}

func (n *nodifyService) SetCrmStore(store crm.CrmStore) {
	n.crmStore = store
}

func (n *nodifyService) CusNodify(store crm.CrmStore) error {
	action, ok := store.GetType()
	if !ok {
		return errors.WithCode(code.ErrCrmTypeIgnore, "")
	}

	applet, err := n.store.CrmApplet().Get(store.GetAppId())
	if err != nil {
		return err
	}
	store.SetApplet(applet)

	switch action {
	case "addCus":
		return nil
	case "updCus":
		return nil
	}
	return nil
}

func (n *nodifyService) addCus() {
	
}

func (n *nodifyService) updCus() {

}
