package request

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type BizDatasetSearch struct {
	CreatedAtRange []time.Time `json:"createdAtRange" form:"createdAtRange[]"`
	request.PageInfo
}

type BizDatasetCreate struct {
	DatasetName *string `json:"datasetName" form:"datasetName" binding:"required"` // 数据集名称
	Scope       *int64  `json:"scope" form:"scope" binding:"required"`             // 权限
}

type BizDatasetUpdate struct {
	ID          *uint   `json:"ID" binding:"required"`          // 主键ID
	DatasetName *string `json:"datasetName" form:"datasetName"` // 数据集名称
	Scope       *int64  `json:"scope" form:"scope"`             // 权限
}
