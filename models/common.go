package models

import (
	"gopkg.in/mgo.v2/bson"
)

// Predefined model error codes.
const (
	ErrDatabase = -1
	ErrSystem   = -2
	ErrDupRows  = -3
	ErrNotFound = -4
)

const (
	DBTablePage = "page"
	DBTableUser = "user"
	DBTablePromoCode = "promo_code"
	DBTableMenu = "menu"
	DBTableAgency = "agency"
	DBTableProduct = "product"
	DBTableOrder = "order"
	DBTableDeliveryCity = "delivery_city"
)

// CodeInfo definiton.
type CodeInfo struct {
	Code int    `json:"code"`
	Info string `json:"info"`
}

type ListResponse struct {
	Items interface{} `json:"items"`
	Total int         `json:"total"`
}

// NewErrorInfo return a CodeInfo represents error.
func NewErrorInfo(info string) *CodeInfo {
	return &CodeInfo{-1, info}
}

// NewNormalInfo return a CodeInfo represents OK.
func NewNormalInfo(info string) *CodeInfo {
	return &CodeInfo{0, info}
}

func JsonToBson(json string) interface{} {
	var bsonDoc interface{}
	bson.UnmarshalJSON([]byte(json),&bsonDoc)

	return bsonDoc
}
