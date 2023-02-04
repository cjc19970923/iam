package common

type NodifyParams interface {
	GetCrmType() string
	GetType() (string, bool)
	GetCusId() string
}
