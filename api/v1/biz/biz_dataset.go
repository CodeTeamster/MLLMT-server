package biz

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/biz"
	"github.com/flipped-aurora/gin-vue-admin/server/model/biz/request"
	bizReq "github.com/flipped-aurora/gin-vue-admin/server/model/biz/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type BizDatasetApi struct{}

// CreateBizDataset 创建数据集管理
// @Tags BizDataset
// @Summary 创建数据集管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body biz.BizDataset true "创建数据集管理"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /dataset/createBizDataset [post]
func (datasetApi *BizDatasetApi) CreateBizDataset(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var datasetReq request.BizDatasetCreate
	err := c.ShouldBindJSON(&datasetReq)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	creatorId := int64(utils.GetUserID(c))
	dataset := biz.BizDataset{
		DatasetName: datasetReq.DatasetName,
		Scope:       datasetReq.Scope,
		CreatorId:   &creatorId,
	}

	err = datasetService.CreateBizDataset(ctx, &dataset)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteBizDataset 删除数据集管理
// @Tags BizDataset
// @Summary 删除数据集管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body biz.BizDataset true "删除数据集管理"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /dataset/deleteBizDataset [delete]
func (datasetApi *BizDatasetApi) DeleteBizDataset(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	ID := c.Query("ID")
	err := datasetService.DeleteBizDataset(ctx, ID)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteBizDatasetByIds 批量删除数据集管理
// @Tags BizDataset
// @Summary 批量删除数据集管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /dataset/deleteBizDatasetByIds [delete]
func (datasetApi *BizDatasetApi) DeleteBizDatasetByIds(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	IDs := c.QueryArray("IDs[]")
	err := datasetService.DeleteBizDatasetByIds(ctx, IDs)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateBizDataset 更新数据集管理
// @Tags BizDataset
// @Summary 更新数据集管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body biz.BizDataset true "更新数据集管理"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /dataset/updateBizDataset [put]
func (datasetApi *BizDatasetApi) UpdateBizDataset(c *gin.Context) {
	// 从ctx获取标准context进行业务行为
	ctx := c.Request.Context()

	var datasetReq request.BizDatasetUpdate
	err := c.ShouldBindJSON(&datasetReq)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	dataset := biz.BizDataset{
		GVA_MODEL:   global.GVA_MODEL{ID: *datasetReq.ID},
		DatasetName: datasetReq.DatasetName,
		Scope:       datasetReq.Scope,
	}

	err = datasetService.UpdateBizDataset(ctx, dataset)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindBizDataset 用id查询数据集管理
// @Tags BizDataset
// @Summary 用id查询数据集管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param ID query uint true "用id查询数据集管理"
// @Success 200 {object} response.Response{data=biz.BizDataset,msg=string} "查询成功"
// @Router /dataset/findBizDataset [get]
func (datasetApi *BizDatasetApi) FindBizDataset(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	ID := c.Query("ID")
	redataset, err := datasetService.GetBizDataset(ctx, ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(redataset, c)
}

// GetBizDatasetList 分页获取数据集管理列表
// @Tags BizDataset
// @Summary 分页获取数据集管理列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query bizReq.BizDatasetSearch true "分页获取数据集管理列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /dataset/getBizDatasetList [get]
func (datasetApi *BizDatasetApi) GetBizDatasetList(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var pageInfo bizReq.BizDatasetSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	userId := utils.GetUserID(c)

	list, total, err := datasetService.GetBizDatasetInfoList(ctx, pageInfo, userId)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败:"+err.Error(), c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取成功", c)
}

// GetBizDatasetPublic 不需要鉴权的数据集管理接口
// @Tags BizDataset
// @Summary 不需要鉴权的数据集管理接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /dataset/getBizDatasetPublic [get]
func (datasetApi *BizDatasetApi) GetBizDatasetPublic(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	datasetService.GetBizDatasetPublic(ctx)
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的数据集管理接口信息",
	}, "获取成功", c)
}
