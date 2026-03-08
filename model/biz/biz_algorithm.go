// 自动生成模板BizAlgorithm
package biz

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 推理加速算法 结构体  BizAlgorithm
type BizAlgorithm struct {
	global.GVA_MODEL
	AlgorithmName *string `json:"algorithmName" form:"algorithmName" gorm:"column:algorithm_name;" binding:"required"`         //算法名
	CreatorName   *string `json:"creatorName" form:"creatorName" gorm:"column:creator_name;" binding:"required"`               //创建人昵称
	CreatorId     *int64  `json:"creatorId" form:"creatorId" gorm:"index;comment:创建人ID;column:creator_id;" binding:"required"` //创建人ID
}

// TableName 推理加速算法 BizAlgorithm自定义表名 biz_algorithm
func (BizAlgorithm) TableName() string {
	return "biz_algorithm"
}
