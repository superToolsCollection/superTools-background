package dao

import (
	"strings"
	"superTools-background/global"
	"superTools-background/pkg/db"
	"superTools-background/pkg/idGenerator"
	"superTools-background/pkg/setting"
	"testing"
)

/**
* @Author: super
* @Date: 2020-11-25 14:21
* @Description:
**/

func TestToolManager_Insert(t *testing.T) {
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
	err = idGenerator.InitSnowflake()
	if err != nil {
		t.Error(err)
	}
	toolManager := NewToolManager("tools", conn)
	tool := &Tool{
		ID:          idGenerator.GenerateID(),
		CreatedBy:   "susu",
		Name:        "test",
		API:         "/api/v1/products/all",
		APIDescribe: "test",
	}
	result, err := toolManager.Insert(tool)
	if err != nil {
		t.Error(err)
	}
	t.Log(result)
}
