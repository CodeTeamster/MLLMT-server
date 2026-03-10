package request

import "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"

type BizInferenceLogSearch struct {
	TaskHash    *string `json:"taskHash" form:"taskHash" binding:"required"`   //任务编号
	DatasetId   *int64  `json:"datasetId" form:"datasetId" binding:"required"` //数据集Id
	SampleId    *int64  `json:"sampleId" form:"sampleId" binding:"required"`   //样本Id
	AlgorithmId *int64  `json:"algorithmId" form:"algorithmId"`                //算法Id
}

type BizInferenceLogSample struct {
	request.PageInfo
	TaskHash *string `json:"taskHash" form:"taskHash" binding:"required"` //任务编号
}
