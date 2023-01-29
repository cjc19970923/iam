package nodify

import (
	"github.com/gin-gonic/gin"
	"github.com/marmotedu/component-base/pkg/core"
	"github.com/marmotedu/errors"
	"github.com/marmotedu/iam/internal/crmapiserver/model/params/nodify"
	srvv1 "github.com/marmotedu/iam/internal/crmapiserver/service/v1"
	"github.com/marmotedu/iam/internal/crmapiserver/store"
	"github.com/marmotedu/iam/internal/pkg/code"
	"github.com/marmotedu/iam/pkg/log"
	"reflect"
)

type NodifyController struct {
	srv srvv1.Service
}

func NewNodifyController(store store.Factory) *NodifyController {
	return &NodifyController{
		srv: srvv1.NewService(store),
	}
}

func (n *NodifyController) Nodify(c *gin.Context) {
	//fmt.Println("这是nodify函数")
	//判断是哪个应用 是否存在这个应用 应用是否开启了
	//判断应用是上游还是下游
	//根据各个crm传过来的code 判断是什么操作
	//实例化出实现了CrmApp接口的struct
	//调用对应的操作
	//如果缺少资料 调用通知接口
	//实现回调 访问客户标签系统 短信等等

	log.L(c).Info("crm nodify")

	v, _ := c.Get("crmInfo")
	crmInfo, ok := v.(map[string]string)
	if !ok {
		err := errors.New("crmInfo is not map")
		core.WriteResponse(c, err, nil)
		return
	}

	//var r nodify.NodifyParams
	//switch crmInfo["crmType"] {
	//case "ec":
	//	r = &nodify.EcNodifyParams{}
	//case "huke":
	//	r = &nodify.HukeNodifyParams{}
	//default:
	//	err := errors.WithCode(code.ErrCrmAppNotFound, "")
	//	core.WriteResponse(c, err, nil)
	//	return
	//}

	r, ok := nodify.CrmTypeMapping[crmInfo["crmType"]]
	if !ok {
		err := errors.WithCode(code.ErrCrmAppNotFound, "")
		core.WriteResponse(c, err, nil)
		return
	}

	val := reflect.New(reflect.ValueOf(r).Elem().Type()).Interface().(nodify.NodifyParams)

	if err := c.ShouldBindJSON(val); err != nil {
		core.WriteResponse(c, errors.WithCode(code.ErrBind, err.Error()), nil)
		return
	}

	//设置crmstore层
	//n.srv.Nodify().SetCrmStore()

	err := n.srv.Nodify().CusNodify(val, crmInfo["appId"])
	if err != nil {
		core.WriteResponse(c, err, "cusNodify service")
		return
	}

	core.WriteResponse(c, errors.WithCode(code.ErrSuccess, ""), nil)

}
