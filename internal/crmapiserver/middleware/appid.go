package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/marmotedu/component-base/pkg/core"
	"github.com/marmotedu/errors"
	"github.com/marmotedu/iam/internal/pkg/code"
)

func GetAppId(c *gin.Context) {
	fmt.Println("appId middleware...")
	id := ""
	crmType := "ec"
	if hukeId := c.Request.Header.Get("ur-appkey"); hukeId != "" { //网易互客
		id = hukeId
		crmType = "huke"
	} else if ecId := c.Request.Header.Get("x-ec-cid"); ecId != "" { //ec
		id = ecId
	}

	if id == "" {
		core.WriteResponse(c, errors.WithCode(code.ErrCrmAppIdIsNull, "appId is null"), nil)
		c.Abort()
	}

	//ca,err := store.Client().CrmApplet().Get(id)
	//if err != nil {
	//	core.WriteResponse(c,errors.WrapC(err,code.ErrCrmAppNotFound,""),nil)
	//}

	c.Set("crmInfo", map[string]string{"appId": id, "crmType": crmType})
	c.Next()
}
