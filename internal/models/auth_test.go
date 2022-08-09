package models

import (
	"gitee.com/phper95/pkg/db"
	"shop/pkg/global"
	"testing"
)

func TestAuth(t *testing.T) {
	err := db.InitMysqlClient(db.DefaultClient, "root", "admin123", "127.0.0.1:3306", "shop")
	if err != nil {
		t.Error("InitMysqlClient error", err, "client", db.DefaultClient)
	}
	global.Db = db.GetMysqlClient(db.DefaultClient).DB

	if err := CreateAuthTable(); err != nil {
		t.Error("create table error", err)
	}

	auth := Auth{
		BusinessKey:       "AK20220808327988",
		BusinessSecret:    "xOBYfykyFVixXFziF8XN5F9crzpC0XrW",
		BusinessDeveloper: "",
		Remark:            "",
		IsUsed:            1, //启用
		CreatedUser:       "",
		UpdatedUser:       "",
	}

	//插入数据
	err = AddBusiness(&auth)
	if err != nil {
		t.Error("AddBusiness error", err)
	}
}
