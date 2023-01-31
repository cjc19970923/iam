package crm

import (
	"strconv"
)

var hukeTypeMapping map[string]string = map[string]string{
	"1010": "addCus",
	"1001": "mergeCus",
	"1002": "addCus",
	"1005": "addLink",
	"1006": "updLink",
	"1007": "delLink",
}

type HukeStore struct {
	common
	Type        int              `json:"type" binding:"required,number"`
	Cid         int              `json:"cid" binding:"required,number"`
	Eid         int              `json:"eid"`
	MergedCid   []int            `json:"mergedCid"`
	CidList     []int            `json:"cidList"`
	ContactList []map[string]int `json:"contactList"`
}

func (h *HukeStore) GetAppId() string {
	return h.AppId
}

func (h *HukeStore) GetCrmType() string {
	return "huke"
}

func (h *HukeStore) GetType() (string, bool) {
	action := strconv.Itoa(h.Type)
	action, ok := ecTypeMapping[action]
	return action, ok
}

func (h *HukeStore) GetCusId() string {
	return strconv.Itoa(h.Cid)
}
