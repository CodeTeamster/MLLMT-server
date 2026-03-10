package biz

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/biz"
	bizReq "github.com/flipped-aurora/gin-vue-admin/server/model/biz/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type BizInferenceLogApi struct{}

// CreateBizInferenceLog 创建推理详细记录
// @Tags BizInferenceLog
// @Summary 创建推理详细记录
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body biz.BizInferenceLog true "创建推理详细记录"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /inferenceLog/createBizInferenceLog [post]
func (inferenceLogApi *BizInferenceLogApi) CreateBizInferenceLog(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var inferenceLog biz.BizInferenceLog
	err := c.ShouldBindJSON(&inferenceLog)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = inferenceLogService.CreateBizInferenceLog(ctx, &inferenceLog)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteBizInferenceLog 删除推理详细记录
// @Tags BizInferenceLog
// @Summary 删除推理详细记录
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body biz.BizInferenceLog true "删除推理详细记录"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /inferenceLog/deleteBizInferenceLog [delete]
func (inferenceLogApi *BizInferenceLogApi) DeleteBizInferenceLog(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	ID := c.Query("ID")
	err := inferenceLogService.DeleteBizInferenceLog(ctx, ID)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteBizInferenceLogByIds 批量删除推理详细记录
// @Tags BizInferenceLog
// @Summary 批量删除推理详细记录
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /inferenceLog/deleteBizInferenceLogByIds [delete]
func (inferenceLogApi *BizInferenceLogApi) DeleteBizInferenceLogByIds(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	IDs := c.QueryArray("IDs[]")
	err := inferenceLogService.DeleteBizInferenceLogByIds(ctx, IDs)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateBizInferenceLog 更新推理详细记录
// @Tags BizInferenceLog
// @Summary 更新推理详细记录
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body biz.BizInferenceLog true "更新推理详细记录"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /inferenceLog/updateBizInferenceLog [put]
func (inferenceLogApi *BizInferenceLogApi) UpdateBizInferenceLog(c *gin.Context) {
	// 从ctx获取标准context进行业务行为
	ctx := c.Request.Context()

	var inferenceLog biz.BizInferenceLog
	err := c.ShouldBindJSON(&inferenceLog)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = inferenceLogService.UpdateBizInferenceLog(ctx, inferenceLog)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindBizInferenceLog 用id查询推理详细记录
// @Tags BizInferenceLog
// @Summary 用id查询推理详细记录
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param ID query uint true "用id查询推理详细记录"
// @Success 200 {object} response.Response{data=biz.BizInferenceLog,msg=string} "查询成功"
// @Router /inferenceLog/findBizInferenceLog [get]
func (inferenceLogApi *BizInferenceLogApi) FindBizInferenceLog(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	ID := c.Query("ID")
	reinferenceLog, err := inferenceLogService.GetBizInferenceLog(ctx, ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(reinferenceLog, c)
}

// GetBizInferenceLogSample 分页获取当前任务所有样本内容
// @Tags BizInferenceLog
// @Summary 分页获取当前任务所有样本内容
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query bizReq.BizInferenceLogSample true "分页获取当前任务所有样本内容"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /inferenceLog/getBizInferenceLogSample [get]
func (inferenceLogApi *BizInferenceLogApi) GetBizInferenceLogSample(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var pageInfo bizReq.BizInferenceLogSample
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := inferenceLogService.GetBizInferenceLogSample(ctx, pageInfo)
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

// GetBizInferenceLogList 分页获取推理详细记录列表
// @Tags BizInferenceLog
// @Summary 分页获取推理详细记录列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query bizReq.BizInferenceLogSearch true "分页获取推理详细记录列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /inferenceLog/getBizInferenceLogList [get]
func (inferenceLogApi *BizInferenceLogApi) GetBizInferenceLogList(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var pageInfo bizReq.BizInferenceLogSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := inferenceLogService.GetBizInferenceLogInfoList(ctx, pageInfo)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败:"+err.Error(), c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:  list,
		Total: total,
	}, "获取成功", c)
}

// GetBizInferenceLogPublic 不需要鉴权的推理详细记录接口
// @Tags BizInferenceLog
// @Summary 不需要鉴权的推理详细记录接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /inferenceLog/getBizInferenceLogPublic [get]
func (inferenceLogApi *BizInferenceLogApi) GetBizInferenceLogPublic(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	inferenceLogService.GetBizInferenceLogPublic(ctx)
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的推理详细记录接口信息",
	}, "获取成功", c)
}
