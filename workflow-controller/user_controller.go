package controller

import (
	"github.com/CaoJiayuan/go-workflow/utils"
	"github.com/CaoJiayuan/go-workflow/workflow-engine/service"
	"github.com/mumushuiding/util"
	"net/http"
)

func GetUser(writer http.ResponseWriter, request *http.Request) {
	if request.Method != "POST" {
		token := request.URL.Query().Get("token")
		if token == "" {
			util.ResponseErr(writer, "token 不能为空")
			return
		}
		user, err := service.GetUserinfoFromRedis(token)

		if err != nil {
			util.ResponseErr(writer, err)
			return
		}
		user.Token = token
		utils.ResponseJson(writer, user)
		return
	}

	var user service.UserInfo

	err := util.Body2Struct(request, &user)
	if err != nil {
		util.ResponseErr(writer, err)
		return
	}

	if user.ID == "" {
		util.ResponseErr(writer, "ID不能为空")
		return
	}
	if user.Company == "" {
		util.ResponseErr(writer, "Company不能为空")
		return
	}

	if user.Department == "" {
		util.ResponseErr(writer, "Department不能为空")
		return
	}

	if len(user.Roles) < 1 {
		util.ResponseErr(writer, "Roles不能为空")
		return
	}

	if len(user.Departments) < 1 {
		util.ResponseErr(writer, "Departments不能为空")
		return
	}

	user.Token, _ = service.SetUserInfoToRedis(user)

	utils.ResponseJson(writer, user)
}
