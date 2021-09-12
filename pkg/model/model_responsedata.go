package model

import (
	"gorm.io/gorm"
)

type Responsedata struct {
	MID_        string `bson:"_id,omitempty" gorm:"-" json:"-" xml:"mid_"`
	Objectuuid  string `bson:"objectuuid" json:"objectuuid" xml:"objectuuid"`
	Intpayresp  string `bson:"intpayresp" json:"intpayresp" xml:"intpayresp"`
	Dompayresp  string `bson:"dompayresp" json:"dompayresp" xml:"dompayresp"`
	Fundconresp string `bson:"fundconresp" json:"fundconresp" xml:"fundconresp"`
}

func (Responsedata) TableName() string {
	return "responsedata"
}
func (m *Responsedata) PreloadResponsedata(db *gorm.DB) *gorm.DB {
	return db
}
