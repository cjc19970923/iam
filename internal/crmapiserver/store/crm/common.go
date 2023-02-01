package crm

import (
	"github.com/marmotedu/iam/internal/crmapiserver/model"
	nodifyParams "github.com/marmotedu/iam/internal/crmapiserver/model/params/nodify/pinterface"
)

type Common struct {
	nodifyParams.NodifyParams
	*model.CrmApplet
}
