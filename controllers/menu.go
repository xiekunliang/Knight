package controllers

import (
	"Knight/utils"
)

// MenuController 菜单模块
type MenuControllertroller struct {
	BaseController
}

// Post 获取菜单
func (ctl *MenuControllertroller) Post() {
	response := make(map[string]interface{})
	response["code"] = utils.FailedCode
	response["msg"] = utils.NoMenuMsg
	if ctl.User.ID == "" {
		return
	}
	// var requestBody = make(map[string]string)
	// json.Unmarshal(ctl.Ctx.Input.RequestBody, &requestBody)
	// username := requestBody["groups"]
	// password := requestBody["isAdmin"]
	// if strings.TrimSpace(username) != "" {
	// 	if user, err := service.ServiceUserMenu(username, password); err == nil {
	// 		ctl.SetSession("User", user)
	// 		response["code"] = utils.SuccessCode
	// 		response["msg"] = utils.SuccessMsg
	// 		data := make(map[string]interface{})
	// 		data["user"] = &user
	// 		response["data"] = data
	// 		if groups, err := service.ServiceGetQuanxianGroupsByID(user.ID); err == nil {
	// 			leng := len(groups)
	// 			groupIDs := make([]string, leng, leng)
	// 			for index, group := range groups {
	// 				groupIDs[index] = group.ID
	// 			}
	// 			if len(groupIDs) == 0 {
	// 				response["msg"] = utils.NoPermissionMsg
	// 			} else {
	// 				data["groups"] = groupIDs
	// 			}
	// 		}
	// 	}
	// }
	ctl.Data["json"] = response
	ctl.ServeJSON()
}
