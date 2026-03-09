package initialize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/biz"
)

func bizModel() error {
	db := global.GVA_DB
	err := db.AutoMigrate(biz.BizDataset{}, biz.BizSample{}, biz.BizAlgorithm{}, biz.BizModel{}, biz.BizInferenceTask{})
	if err != nil {
		return err
	}
	return nil
}
