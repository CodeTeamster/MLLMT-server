package biz

import api "github.com/flipped-aurora/gin-vue-admin/server/api/v1"

type RouterGroup struct{ BizDatasetRouter }

var datasetApi = api.ApiGroupApp.BizApiGroup.BizDatasetApi
