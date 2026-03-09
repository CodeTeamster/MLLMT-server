package biz

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct {
	BizDatasetApi
	BizSampleApi
	BizAlgorithmApi
	BizModelApi
	BizInferenceTaskApi
}

var (
	datasetService       = service.ServiceGroupApp.BizServiceGroup.BizDatasetService
	sampleService        = service.ServiceGroupApp.BizServiceGroup.BizSampleService
	algorithmService     = service.ServiceGroupApp.BizServiceGroup.BizAlgorithmService
	modelService         = service.ServiceGroupApp.BizServiceGroup.BizModelService
	inferenceTaskService = service.ServiceGroupApp.BizServiceGroup.BizInferenceTaskService
)
