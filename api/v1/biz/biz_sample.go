package biz

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/biz"
	bizReq "github.com/flipped-aurora/gin-vue-admin/server/model/biz/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type BizSampleApi struct{}

// CreateBizSample 创建数据集样本管理
// @Tags BizSample
// @Summary 创建数据集样本管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body biz.BizSample true "创建数据集样本管理"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /sample/createBizSample [post]
func (sampleApi *BizSampleApi) CreateBizSample(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var sample biz.BizSample
	err := c.ShouldBindJSON(&sample)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = sampleService.CreateBizSample(ctx, &sample)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteBizSample 删除数据集样本管理
// @Tags BizSample
// @Summary 删除数据集样本管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body biz.BizSample true "删除数据集样本管理"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /sample/deleteBizSample [delete]
func (sampleApi *BizSampleApi) DeleteBizSample(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	ID := c.Query("ID")
	err := sampleService.DeleteBizSample(ctx, ID)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteBizSampleByIds 批量删除数据集样本管理
// @Tags BizSample
// @Summary 批量删除数据集样本管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /sample/deleteBizSampleByIds [delete]
func (sampleApi *BizSampleApi) DeleteBizSampleByIds(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	IDs := c.QueryArray("IDs[]")
	err := sampleService.DeleteBizSampleByIds(ctx, IDs)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateBizSample 更新数据集样本管理
// @Tags BizSample
// @Summary 更新数据集样本管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body biz.BizSample true "更新数据集样本管理"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /sample/updateBizSample [put]
func (sampleApi *BizSampleApi) UpdateBizSample(c *gin.Context) {
	// 从ctx获取标准context进行业务行为
	ctx := c.Request.Context()

	var sample biz.BizSample
	err := c.ShouldBindJSON(&sample)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = sampleService.UpdateBizSample(ctx, sample)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindBizSample 用id查询数据集样本管理
// @Tags BizSample
// @Summary 用id查询数据集样本管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param ID query uint true "用id查询数据集样本管理"
// @Success 200 {object} response.Response{data=biz.BizSample,msg=string} "查询成功"
// @Router /sample/findBizSample [get]
func (sampleApi *BizSampleApi) FindBizSample(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	ID := c.Query("ID")
	resample, err := sampleService.GetBizSample(ctx, ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(resample, c)
}

// GetBizSampleList 分页获取数据集样本管理列表
// @Tags BizSample
// @Summary 分页获取数据集样本管理列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query bizReq.BizSampleSearch true "分页获取数据集样本管理列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /sample/getBizSampleList [get]
func (sampleApi *BizSampleApi) GetBizSampleList(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var pageInfo bizReq.BizSampleSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := sampleService.GetBizSampleInfoList(ctx, pageInfo)
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

// GetBizSamplePublic 不需要鉴权的数据集样本管理接口
// @Tags BizSample
// @Summary 不需要鉴权的数据集样本管理接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /sample/getBizSamplePublic [get]
func (sampleApi *BizSampleApi) GetBizSamplePublic(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	sampleService.GetBizSamplePublic(ctx)
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的数据集样本管理接口信息",
	}, "获取成功", c)
}
