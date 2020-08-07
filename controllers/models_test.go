package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gin-restful-best-practice/conf"
	"gin-restful-best-practice/models"
	"gin-restful-best-practice/utils"
	"testing"
)

var (
	testUser = gin.H{
		"username":     "models_test_user",
		"password":     "none",
		"phone":        "11100001111",
		"avatar":       "https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1596792660550&di=93fdd606a70b2940d8315d603964b031&imgtype=0&src=http%3A%2F%2Fimg2018.cnblogs.com%2Fblog%2F294888%2F201903%2F294888-20190320024006427-1092074502.png",
		"mail":         "dev@github.com",
		"organization": "University of Beijing",
	}
	testModel = gin.H{
		"name": "models_test_model",
		"type": models.IMAGE_CLASSIFICATION,
	}
)

func TestMain(t *testing.M) {
	gin.SetMode(gin.TestMode)
	if conf.Conf().ENV != "dev" {
		panic("test: current ENV is not DEV, clean data failed")
	}
	if err := models.DB.DropTableIfExists("model", "user", "role").Error; err != nil {
		panic(err)
	}
	models.InitTableRole()
	models.InitTableUser()
	models.InitTableModel()
	t.Run()
}


func TestCreateModel(t *testing.T) {
	// create user
	res := testRequest(CreateUser, "POST", "/api/users", testUser)
	assert.Equal(t, 200, res.Code)
	// assign user_id to model
	testModel["user_id"] = res.Data["id"]
	// create model
	res = testRequest(CreateModel, "POST", "/api/models", testModel)
	assert.Equal(t, 200, res.Code)
	assert.Equal(t, testModel["name"], res.Data["name"])
	t.Log(utils.JSONToString(res.Data))
}

func TestFetchModelList(t *testing.T) {
	res := testRequest(FetchModelList, "GET", "/api/models", gin.H{
		"limit": 50,
	})
	assert.Equal(t, 200, res.Code)
	assert.Len(t, res.Data["data"], 1)
	t.Log(utils.JSONToString(res.Data))
}

func TestFetchModelListByUserID(t *testing.T) {
	res := testRequest(FetchModelList, "GET", "/api/models", gin.H{
		"limit": 50,
		"user_id": testModel["user_id"],
	})
	assert.Equal(t, 200, res.Code)
	assert.Len(t, res.Data["data"], 1)
	t.Log(utils.JSONToString(res.Data))
}
