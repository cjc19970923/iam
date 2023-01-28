package mysql

import (
	"github.com/marmotedu/iam/internal/crmapiserver/model"
	"gorm.io/gorm"
)

type admins struct {
	db *gorm.DB
}

func newAdmins(ds *datastore) *admins {
	return &admins{ds.db}
}

func (a *admins) Create(admin *model.UctAdmin) error {
	return a.db.Create(admin).Error
}
