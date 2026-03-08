package request

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type BizAlgorithmSearch struct {
	CreatedAtRange []time.Time `json:"createdAtRange" form:"createdAtRange[]"`
	request.PageInfo
}

type BizAlgorithmCreate struct {
	AlgorithmName *string `json:"algorithmName" form:"algorithmName" binding:"required"`
}
