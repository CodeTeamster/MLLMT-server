package initialize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/biz"
)

func bizModel() error {
	db := global.GVA_DB
	err := db.AutoMigrate(biz.BizDataset{})
	if err != nil {
		return err
	}
	return nil
}
