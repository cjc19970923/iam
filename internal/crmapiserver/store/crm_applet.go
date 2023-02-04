package store

import "github.com/marmotedu/iam/internal/crmapiserver/model"

type CrmAppletStore interface {
	Get(appId string) (*model.CrmApplet, error)
	GetAppletByType(crmType string, cusType string) (*model.CrmApplet, error)
}
