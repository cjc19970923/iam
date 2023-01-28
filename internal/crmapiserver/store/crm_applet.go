package store

import "github.com/marmotedu/iam/internal/crmapiserver/model"

type CrmAppletStore interface {
	Get(appId string) (*model.CrmApplet, error)
}
