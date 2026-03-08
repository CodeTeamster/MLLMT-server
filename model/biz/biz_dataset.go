// 自动生成模板BizDataset
package biz

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 数据集管理 结构体  BizDataset
type BizDataset struct {
	global.GVA_MODEL
	DatasetName *string `json:"datasetName" form:"datasetName" gorm:"column:dataset_name;" binding:"required"`                   //数据集名称
	Scope       *int64  `json:"scope" form:"scope" gorm:"default:0;comment:0: 仅自己可见  1: 所有人可见;column:scope;" binding:"required"` //权限
	CreatorId   *int64  `json:"creatorId" form:"creatorId" gorm:"column:creator_id;"`                                            //创建人id
}

// TableName 数据集管理 BizDataset自定义表名 biz_dataset
func (BizDataset) TableName() string {
	return "biz_dataset"
}
