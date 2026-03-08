package request

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type BizSampleSearch struct {
	CreatedAtRange []time.Time `json:"createdAtRange" form:"createdAtRange[]"`
	request.PageInfo
	DatasetId *int64 `json:"datasetId" form:"datasetId" binding:"required"` //数据集ID
}
