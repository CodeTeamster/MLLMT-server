// 自动生成模板BizModel
package biz

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 模型管理 结构体  BizModel
type BizModel struct {
	global.GVA_MODEL
	ModelName     *string `json:"modelName" form:"modelName" gorm:"comment:模型名称;column:model_name;" binding:"required"`                           //模型名称
	AlgorithmName *string `json:"algorithmName" form:"algorithmName" gorm:"comment:支持算法名;column:algorithm_name;" binding:"required"`              //支持算法名
	AlgorithmId   *int64  `json:"algorithmId" form:"algorithmId" gorm:"comment:支持算法ID;column:algorithm_id;" binding:"required"`                   //支持算法ID
	Type          *int64  `json:"type" form:"type" gorm:"default:0;comment:模型类型：0: Local 1: OpenAI 2: Anthropic;column:type;" binding:"required"` //模型类型
	Enable        *int64  `json:"enable" form:"enable" gorm:"default:1;comment:是否启用：0: 禁用 1:启用;column:enable;" binding:"required"`                //是否启用
	CreatorName   *string `json:"creatorName" form:"creatorName" gorm:"comment:创建人用户名;column:creator_name;" binding:"required"`                   //创建人用户名
	CreatorId     *int64  `json:"creatorId" form:"creatorId" gorm:"comment:创建人ID;column:creator_id;" binding:"required"`                          //创建人ID
}

// TableName 模型管理 BizModel自定义表名 biz_model
func (BizModel) TableName() string {
	return "biz_model"
}
