package crm

import (
	"github.com/marmotedu/iam/internal/crmapiserver/model"
	"github.com/marmotedu/iam/internal/crmapiserver/model/params/nodify"
)

type Common struct {
	nodify.NodifyParams
	*model.CrmApplet
}
