package controllers

import (
	"github.com/gin-gonic/gin"
	"gin-restful-best-practice/models"
	"net/http"
	"strconv"
)

type FetchUserListForm struct {
	Offset int `form:"offset" binding:"number"`
	Limit  int `form:"limit" binding:"required,number"`
}

func FetchUserList(ctx *gin.Context) {
	var form FetchUserListForm
	if err := ctx.ShouldBindQuery(&form); err != nil {
		ErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	users, err := models.GetUserList(form.Offset, form.Limit)
	if err != nil {
		ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	for i, _ := range users {
		users[i].Password = ""
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data":   users,
		"total":  len(users),
		"offset": form.Offset,
		"limit":  form.Limit,
	})
}

type CreateUserForm struct {
	Username     string `json:"username" binding:"required,gte=2,lte=20"`
	Password     string `json:"password" binding:"required"`
	Phone        string `json:"phone" binding:"required,lte=20"`
	Avatar       string `json:"avatar"`
	Mail         string `json:"mail"`
	Organization string `json:"organization"`
}

func CreateUser(ctx *gin.Context) {
	var form CreateUserForm
	if err := ctx.ShouldBindJSON(&form); err != nil {
		ErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	user := models.User{
		Username:     form.Username,
		Password:     form.Password,
		Phone:        form.Phone,
		Avatar:       form.Avatar,
		Mail:         form.Mail,
		Organization: form.Organization,
	}
	if err := user.Insert(); err != nil {
		ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	user.Password = ""
	ctx.JSON(http.StatusOK, user)
}

func FetchUserByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	user, err := models.GetUserByID(id)
	user.Password = ""

	if err != nil {
		ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data": user,
	})
}
