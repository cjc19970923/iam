package crm

import "github.com/marmotedu/iam/internal/crmapiserver/model"

type common struct {
	AppId string
	*model.CrmApplet
}

func (c *common) SetApplet(applet *model.CrmApplet) {
	c.CrmApplet = applet
}
