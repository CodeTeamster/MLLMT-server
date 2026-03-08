package biz

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct{ BizDatasetApi }

var datasetService = service.ServiceGroupApp.BizServiceGroup.BizDatasetService
