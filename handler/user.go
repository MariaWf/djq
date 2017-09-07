package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"mimi/djq/model"
	"strconv"
	"fmt"
	"log"
	"mimi/djq/util"
)

func UserList(c *gin.Context) {
	users := make([]*model.User, 0, 10)
	for i := 0; i < 10; i++ {
		user := &model.User{}
		user.Id = "id" + strconv.Itoa(i)
		user.Name = "name" + strconv.Itoa(i)
		users = append(users, user)
	}
	c.JSON(http.StatusOK, util.BuildSuccessPageResult(1, 10, 123, users))
}

func UserGet(c *gin.Context) {
	user := &model.User{}
	user.Id = c.Param("id")
	user.Name = "name1"
	result := util.BuildSuccessResult(user)
	str,err:=c.Cookie("test")
	if err!=nil{
		fmt.Println(err)
	}else{
		fmt.Println(str)
	}
	if str ==""{
		cookie := http.Cookie{Name: "test", Value: "mimi", Path:"/"}
		http.SetCookie(c.Writer, &cookie)
	}
	//fmt.Println(user)
	//fmt.Println(result)
	//c.JSON(http.StatusOK, BuildSuccessResult(user))
	//panic(errors.New("test_panic"))
	c.JSON(http.StatusOK, result)
}

func UserUploadAction(c *gin.Context) {
	file, _ := c.FormFile("file")
	log.Println(file.Filename)

	// Upload the file to specific dst.
	// c.SaveUploadedFile(file, dst)

	c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
}

func UserMultiUploadAction(c *gin.Context) {
	form, _ := c.MultipartForm()
	files := form.File["upload[]"]

	for _, file := range files {
		log.Println(file.Filename)

		// Upload the file to specific dst.
		// c.SaveUploadedFile(file, dst)
	}
	c.String(http.StatusOK, fmt.Sprintf("%d files uploaded!", len(files)))
}
