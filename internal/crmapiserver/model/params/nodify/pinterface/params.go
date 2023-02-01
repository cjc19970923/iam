package pinterface

import (
	"github.com/marmotedu/iam/internal/crmapiserver/model"
	"github.com/marmotedu/iam/internal/crmapiserver/store/crm"
)

type NodifyParams interface {
	GetCrmType() string
	GetType() (string, bool)
	GetCusId() string
	GetStore(*model.CrmApplet) crm.CrmStore
}
