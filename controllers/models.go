package controllers

import (
	"github.com/gin-gonic/gin"
	"gin-restful-best-practice/models"
	"gin-restful-best-practice/services"
	"net/http"
)

type FetchModelListForm struct {
	Offset int `form:"offset" binding:"number"`
	Limit  int `form:"limit" binding:"required,number"`
	UserID int `form:"user_id" binding:"number"`
}

func FetchModelList(ctx *gin.Context) {
	var form FetchModelListForm
	if err := ctx.ShouldBindQuery(&form); err != nil {
		ErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	modelList, err := models.GetModelList(map[string]interface{}{
		"yn": true,
	}, form.Offset, form.Limit)
	if err != nil {
		ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data":   modelList,
		"total":  len(modelList),
		"offset": form.Offset,
		"limit":  form.Limit,
	})
	return
}

type CreateModelForm struct {
	Name         string `json:"name" binding:"required,gte=2,lte=20"`
	Type         uint   `json:"type" binding:"required,number"`
	URL          string `json:"url"`
	Desc         string `json:"desc"`
	PriceMonthly uint   `json:"price_monthly" binding:"number"`
	PriceYearly  uint   `json:"price_yearly" binding:"number"`
	PriceTotal   uint   `json:"price_total" binding:"number"`
	UserID       uint   `json:"user_id" binding:"required,number"`
}

func CreateModel(ctx *gin.Context) {
	var form CreateModelForm
	if err := ctx.ShouldBindJSON(&form); err != nil {
		ErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	model := models.Model{
		Name:         form.Name,
		Type:         form.Type,
		URL:          form.URL,
		Desc:         form.Desc,
		PriceMonthly: form.PriceMonthly,
		PriceYearly:  form.PriceYearly,
		PriceTotal:   form.PriceTotal,
		UserID:       form.UserID,
	}
	if err := model.Insert(); err != nil {
		ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, model)
}

func GetModelFile(ctx *gin.Context) {
	model, err := services.DownloadModel("name")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
	}
	ctx.JSON(http.StatusOK, gin.H{
		"model": model,
	})
}
