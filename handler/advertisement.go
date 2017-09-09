package handler

import (
	"github.com/gin-gonic/gin"
	"mimi/djq/model"
	"mimi/djq/util"
	"net/http"
	"strconv"
)

func AdvertisementList4Index(c *gin.Context) {
	advertisementList := make([]*model.Advertisement, 0, 5)
	for i := 0; i < 5; i++ {
		advertisement := &model.Advertisement{}
		advertisement.Id = "id" + strconv.Itoa(i)
		advertisement.Name = "name" + strconv.Itoa(i)
		advertisement.Image = "Image" + strconv.Itoa(i)
		advertisement.Link = "link" + strconv.Itoa(i)
		advertisementList = append(advertisementList, advertisement)
	}
	c.JSON(http.StatusOK, util.BuildSuccessResult(advertisementList))
}
