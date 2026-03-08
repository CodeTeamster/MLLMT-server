package biz

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct {
	BizDatasetApi
	BizSampleApi
}

var (
	datasetService = service.ServiceGroupApp.BizServiceGroup.BizDatasetService
	sampleService  = service.ServiceGroupApp.BizServiceGroup.BizSampleService
)
