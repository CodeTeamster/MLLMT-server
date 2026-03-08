package request

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type BizModelSearch struct {
	CreatedAtRange []time.Time `json:"createdAtRange" form:"createdAtRange[]"`
	request.PageInfo
}

type BizModelCreate struct {
	ModelName     *string `json:"modelName" form:"modelName" binding:"required"`         //模型名称
	AlgorithmName *string `json:"algorithmName" form:"algorithmName" binding:"required"` //支持算法名
	AlgorithmId   *int64  `json:"algorithmId" form:"algorithmId" binding:"required"`     //支持算法ID
	Type          *int64  `json:"type" form:"type" binding:"required"`                   //模型类型
	Enable        *int64  `json:"enable" form:"enable" binding:"required"`               //是否启用
}
