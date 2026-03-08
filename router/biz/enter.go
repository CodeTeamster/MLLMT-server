package biz

import api "github.com/flipped-aurora/gin-vue-admin/server/api/v1"

type RouterGroup struct {
	BizDatasetRouter
	BizSampleRouter
	BizAlgorithmRouter
}

var (
	datasetApi   = api.ApiGroupApp.BizApiGroup.BizDatasetApi
	sampleApi    = api.ApiGroupApp.BizApiGroup.BizSampleApi
	algorithmApi = api.ApiGroupApp.BizApiGroup.BizAlgorithmApi
)
