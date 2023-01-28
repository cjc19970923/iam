package store

import "github.com/marmotedu/iam/internal/crmapiserver/model"

type AdminStore interface {
	Create(admin *model.UctAdmin) error
}
