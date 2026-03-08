// 自动生成模板BizSample
package biz

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 数据集样本管理 结构体  BizSample
type BizSample struct {
	global.GVA_MODEL
	DatasetId *int64  `json:"datasetId" form:"datasetId" gorm:"index;comment:绑定的数据集ID;column:dataset_id;" binding:"required"` //数据集ID
	Prompt    *string `json:"prompt" form:"prompt" gorm:"comment:用户提示词;column:prompt;" binding:"required"`                    //用户提示词
	Img       *string `json:"img" form:"img" gorm:"comment:样本图片url;column:img;" binding:"required"`                           //样本图片
}

// TableName 数据集样本管理 BizSample自定义表名 biz_sample
func (BizSample) TableName() string {
	return "biz_sample"
}
