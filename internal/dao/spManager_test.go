package dao

import (
	"strings"
	"superTools-background/global"
	"superTools-background/pkg/db"
	"superTools-background/pkg/setting"
	"testing"
)

/**
* @Author: super
* @Date: 2021-01-25 10:01
* @Description:
**/

func TestAddSpManager(t *testing.T) {
	newSetting, err := setting.NewSetting(strings.Split("/Users/super/develop/superTools-background/configs", ",")...)
	if err != nil {
		t.Error(err)
	}
	err = newSetting.ReadSection("Database", &global.DatabaseSetting)
	if err != nil {
		t.Error(err)
	}
	conn, err := db.NewDBEngine(global.DatabaseSetting)
	if err != nil {
		t.Error(err)
	}
	spManager := NewSpManagerManger("sp_manager", conn)
	p := &SpManager{
		MgName:   "12345",
		MgPwd:    "aaaa",
		MgEmail:  "sssss@qq.com",
		MgMobile: "1666666666",
	}
	result, err := spManager.Insert(p)
	if err != nil {
		t.Error(err)
	}
	t.Log(result)
}
