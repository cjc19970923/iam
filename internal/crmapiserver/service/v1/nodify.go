package v1

import (
	"github.com/marmotedu/errors"
	nodifyParams "github.com/marmotedu/iam/internal/crmapiserver/model/params/nodify/pinterface"
	"github.com/marmotedu/iam/internal/crmapiserver/store"
	"github.com/marmotedu/iam/internal/crmapiserver/store/crm"
	"github.com/marmotedu/iam/internal/pkg/code"
)

type NodifySrv interface {
	CusNodify(params nodifyParams.NodifyParams, appId string) error
	//SetCrmStore(store crm.CrmStore)
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

//func (n *nodifyService) SetCrmStore(store crm.CrmStore) {
//	n.crmStore = store
//}

func (n *nodifyService) CusNodify(params nodifyParams.NodifyParams, appId string) error {
	action, ok := params.GetType()
	if !ok {
		return errors.WithCode(code.ErrCrmTypeIgnore, "")
	}

	applet, err := n.store.CrmApplet().Get(appId)
	if err != nil {
		return err
	}

	n.crmStore = params.GetStore(applet)

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
