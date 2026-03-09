package biz

import api "github.com/flipped-aurora/gin-vue-admin/server/api/v1"

type RouterGroup struct {
	BizDatasetRouter
	BizSampleRouter
	BizAlgorithmRouter
	BizModelRouter
	BizInferenceTaskRouter
}

var (
	datasetApi       = api.ApiGroupApp.BizApiGroup.BizDatasetApi
	sampleApi        = api.ApiGroupApp.BizApiGroup.BizSampleApi
	algorithmApi     = api.ApiGroupApp.BizApiGroup.BizAlgorithmApi
	modelApi         = api.ApiGroupApp.BizApiGroup.BizModelApi
	inferenceTaskApi = api.ApiGroupApp.BizApiGroup.BizInferenceTaskApi
)
