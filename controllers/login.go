package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"gin-restful-best-practice/middlewares"
	"gin-restful-best-practice/models"
	"net/http"
)

type LoginForm struct {
	Username string `json:"username" binding:"required,gte=2,lte=20"`
	Password string `json:"password" binding:"required"`
}

func Login(ctx *gin.Context) {
	var form LoginForm
	if err := ctx.ShouldBindJSON(&form); err != nil {
		ErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	user, err := models.GetUserByUsername(form.Username)
	// invalid username or password
	if err == gorm.ErrRecordNotFound || user.Password != form.Password {
		ErrorResponse(ctx, http.StatusForbidden, "invalid username or password")
		return
	}
	// sql error
	if err != nil {
		ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	token, err := middlewares.CreateJWT(user)
	if err != nil {
		ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"token": "Bearer " + token,
	})
}
