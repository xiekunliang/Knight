package controllers

import (
	service "Knight/services"
	"Knight/utils"
)

// UserControllertroller 用户信息模块
type UserControllertroller struct {
	BaseController
}

func (ctl *UserControllertroller) Post() {
	response := make(map[string]interface{})
	response["code"] = utils.FailedCode
	response["msg"] = utils.FailedMsg
	if ctl.User.ID == "" {
		if id, err := service.ServiceCreateDengluYH(ctl.Ctx.Input.RequestBody, ctl.User.ID); err == nil {
			response["code"] = utils.SuccessCode
			response["msg"] = utils.SuccessMsg
			response["id"] = id
		}
	}
	ctl.Data["json"] = response
	ctl.ServeJSON()
}
