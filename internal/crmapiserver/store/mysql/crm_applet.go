package mysql

import (
	"github.com/marmotedu/errors"
	"github.com/marmotedu/iam/internal/crmapiserver/model"
	"github.com/marmotedu/iam/internal/pkg/code"
	"gorm.io/gorm"
)

type crmApplet struct {
	db *gorm.DB
}

func newCrmApplet(ds *datastore) *crmApplet {
	return &crmApplet{ds.db}
}

func (c *crmApplet) Get(appId string) (*model.CrmApplet, error) {
	applet := &model.CrmApplet{}
	err := c.db.Where("appkey = ?", appId).First(applet).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.WithCode(code.ErrCrmAppNotFound, err.Error())
		}

		return nil, errors.WithCode(code.ErrDatabase, err.Error())
	}

	if applet.Status == "disable" {
		return nil, errors.WithCode(code.ErrCrmAppDisable, "")
	}

	return applet, nil
}

func (c *crmApplet) GetAppletByType(crmType string, cusType string) (*model.CrmApplet, error) {
	applet := &model.CrmApplet{}
	err := c.db.Where("crm_type = ? and type = ? ").First(applet).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.WithCode(code.ErrCrmAppNotFound, err.Error())
		}

		return nil, errors.WithCode(code.ErrDatabase, err.Error())
	}

	if applet.Status == "disable" {
		return nil, errors.WithCode(code.ErrCrmAppDisable, "")
	}

	return applet, nil

}
