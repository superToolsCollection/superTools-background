package dao

import "github.com/jinzhu/gorm"

/**
* @Author: super
* @Date: 2021-02-13 19:48
* @Description:
**/

type SpAttribute struct {
	AttrID     int    `json:"attr_id"`
	AttrName   string `json:"attr_name"`
	CatID      int    `json:"cat_id"`
	AttrSel    string `json:"attr_sel"`
	AttrWrite  string `json:"attr_write"`
	AttrVals   string `json:"attr_vals"`
	DeleteTime int64  `json:"delete_time"`
}

type ISpAttribute interface {
}

type SpAttributeManager struct {
	table string
	conn  *gorm.DB
}

func NewSpAttributeManager(table string, conn *gorm.DB) ISpAttribute {
	return &SpAttributeManager{table: table, conn: conn}
}
