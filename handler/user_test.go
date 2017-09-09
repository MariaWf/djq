package handler

import (
	"github.com/gin-gonic/gin/json"
	"mimi/djq/model"
	"mimi/djq/util"
	"testing"
)

func TestUserGet(t *testing.T) {
	user := &model.User{}
	user.Id = "id1"
	user.Name = "name1"
	result := util.BuildSuccessResult(user)
	//result := &ResultVO2{}
	result.Msg = "msg"
	result.Status = 1
	jsonBytes, err := json.Marshal(result)
	jsonBytes2, err := json.Marshal(user)
	if err != nil {
		t.Error(err)
	}
	t.Log(string(jsonBytes[:]))
	t.Log(string(jsonBytes2[:]))
}

func TestUserList(t *testing.T) {
	t.Log("wawa")
}
