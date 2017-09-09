package handler

import (
	"github.com/gin-gonic/gin"
	"mimi/djq/model"
	"mimi/djq/util"
	"net/http"
	"strconv"
)

func ShopClassificationList4Index(c *gin.Context) {
	shopClassificationList := make([]*model.ShopClassification, 0, 32)
	for i := 0; i < 32; i++ {
		shopClassification := &model.ShopClassification{}
		shopClassification.Id = "id" + strconv.Itoa(i)
		shopClassification.Name = "name" + strconv.Itoa(i)
		shopClassificationList = append(shopClassificationList, shopClassification)
	}
	c.JSON(http.StatusOK, util.BuildSuccessResult(shopClassificationList))
}
