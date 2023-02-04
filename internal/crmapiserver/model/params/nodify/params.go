package nodify

import (
	"github.com/marmotedu/iam/internal/crmapiserver/model"
	"github.com/marmotedu/iam/internal/crmapiserver/model/params/nodify/common"
	"github.com/marmotedu/iam/internal/crmapiserver/store/crm"
)

// 映射对应erp客户的操作
// addCus 添加客户
// updCus 修改客户
// delCus 删除客户
// mergeCus 合并客户
// addLink 添加联系人
// updLink 修改联系人
// delLink 删除联系人
// mergeLink 合并联系人

var CrmTypeMapping map[string]common.NodifyParams = map[string]common.NodifyParams{
	"ec":   &EcNodifyParams{},
	"huke": &HukeNodifyParams{},
}

type Store interface {
	GetStore(*model.CrmApplet) crm.CrmStore
}
