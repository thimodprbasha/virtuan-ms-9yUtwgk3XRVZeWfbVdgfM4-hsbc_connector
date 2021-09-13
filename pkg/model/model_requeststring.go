package model

import (
	"gorm.io/gorm"
)

type Requeststring struct {
	MID_                   string `bson:"_id,omitempty" gorm:"-" json:"-" xml:"mid_"`
	Objectuuid             string `bson:"objectuuid" json:"objectuuid" xml:"objectuuid"`
	Consentid              string `bson:"consentid" json:"consentid" xml:"consentid"`
	X_fapi_financial_id    string `bson:"x_fapi_financial_id" json:"x_fapi_financial_id" xml:"x_fapi_financial_id"`
	Authorization          string `bson:"authorization" json:"authorization" xml:"authorization"`
	Internationalpaymentid string `bson:"internationalpaymentid" json:"internationalpaymentid" xml:"internationalpaymentid"`
	Domesticpaymentid      string `bson:"domesticpaymentid" json:"domesticpaymentid" xml:"domesticpaymentid"`
	Opcode                 string `bson:"opcode" json:"opcode" xml:"opcode"`
}

func (Requeststring) TableName() string {
	return "requeststring"
}
func (m *Requeststring) PreloadRequeststring(db *gorm.DB) *gorm.DB {
	return db
}
