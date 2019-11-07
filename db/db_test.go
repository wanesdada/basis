package db

import (
	"qp_web_server/config"
	"testing"
)

type AccountsInfo struct {
	//`UserID` int(11) NOT NULL AUTO_INCREMENT COMMENT '用户标识',
	//`FaceID` smallint(6) NOT NULL DEFAULT '0' COMMENT '头像标识',
	//`Accounts` varchar(32) NOT NULL COMMENT '用户帐号',
	UserID   int64
	FaceID   int64
	Accounts string
}


func TestInitDB(t *testing.T) {
	config.InitConfig("/../config/config.yml")
	db, err := InitDB(config.GetDb())
	if err != nil {
		t.Fatalf("init db err %v", err)
	}
	defer func() {
		err := db.Close()
		if err != nil {
			t.Fatalf("db close err %v", err)
		}
	}()
	var accountsInfo []AccountsInfo
	err = db.Where("UserID = ?", 30).Find(&accountsInfo)
	if err != nil {
		t.Fatalf("db find err %v", err)
	}
	t.Logf("accountsInfo %v", accountsInfo)
}
