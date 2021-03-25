package controller

import (
	"github.com/CaoJiayuan/go-workflow/utils"
	"github.com/CaoJiayuan/go-workflow/workflow-engine/service"
	"github.com/mumushuiding/util"
	"net/http"
)

func GetUser(writer http.ResponseWriter, request *http.Request) {
	if request.Method != "POST" {
		util.ResponseErr(writer, "只支持Post方法！！Only support Post ")
		return
	}

	var user service.UserInfo

	err := util.Body2Struct(request, &user)
	if err != nil {
		util.ResponseErr(writer, err)
		return
	}

	utils.ResponseJson(writer, user)
}
