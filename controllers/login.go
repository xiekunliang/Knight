package controllers

import (
	service "Knight/services"
	"Knight/utils"
	"encoding/json"
	"strings"
)

// LoginController 登录模块
type LoginControllertroller struct {
	BaseController
}

// Post 登录请求
func (ctl *LoginControllertroller) Post() {
	response := make(map[string]interface{})
	response["code"] = utils.FailedCode
	response["msg"] = utils.UserPasswordMsg
	var requestBody = make(map[string]string)
	json.Unmarshal(ctl.Ctx.Input.RequestBody, &requestBody)
	username := requestBody["username"]
	password := requestBody["password"]
	if strings.TrimSpace(username) != "" {
		if user, err := service.ServiceUserLogin(username, password); err == nil {
			ctl.SetSession("User", user)
			response["code"] = utils.SuccessCode
			response["msg"] = utils.SuccessMsg
			data := make(map[string]interface{})
			data["user"] = &user
			response["data"] = data
			if groups, err := service.ServiceGetQuanxianGroupsByID(user.ID); err == nil {
				leng := len(groups)
				groupIDs := make([]string, leng, leng)
				for index, group := range groups {
					groupIDs[index] = group.ID
				}
				data["groups"] = groupIDs
			}
		}
	}
	ctl.Data["json"] = response
	ctl.ServeJSON()
}

// Get 注销登录请求
func (ctl *LoginControllertroller) Get() {
	response := make(map[string]interface{})
	response["code"] = utils.FailedCode
	response["msg"] = utils.FailedMsg
	if ctl.User.ID == "" {
		return
	}
	IDStr := ctl.Ctx.Input.Param(":id")
	if _, err := service.ServiceUserLogout(IDStr); err == nil {
		response["code"] = utils.SuccessCode
		response["msg"] = utils.SuccessMsg
		ctl.DelSession("User")
	}
	ctl.Data["json"] = response
	ctl.ServeJSON()
}
