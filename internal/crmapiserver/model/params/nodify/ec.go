package nodify

import "strconv"

var ecTypeMapping map[string]string = map[string]string{
	"1":  "addCus",
	"2":  "addCus",
	"3":  "addCus",
	"9":  "addCus",
	"10": "addCus",
	"11": "addCus",
	"7":  "mergeCus",
	"13": "addlink",
	"14": "addLink",
	"15": "addLink",
	"16": "delLink",
	"8":  "delCus",
}

type EcNodifyParams struct {
	Type  int `json:"type" binding:"required,number"`
	CrmId int `json:"crmId" binding:"required,number"`
}

func (e EcNodifyParams) GetCrmType() string {
	return "ec"
}

func (e EcNodifyParams) GetType() (string, bool) {
	action := strconv.Itoa(e.Type)
	action, ok := ecTypeMapping[action]
	return action, ok
}

func (e EcNodifyParams) GetCusId() string {
	return strconv.Itoa(e.CrmId)
}

//func NewEcNodifyParams() *EcNodifyParams {
//	return &EcNodifyParams{}
//}
