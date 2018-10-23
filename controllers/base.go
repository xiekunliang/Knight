package controllers

import (
	md "Knight/models"
	"time"

	"github.com/astaxie/beego"
)

// BaseController 基础controller
type BaseController struct {
	beego.Controller
	IsAdmin bool
	User    md.DengluYH
}

// Prepare implemented Prepare method for baseRouter.
func (ctl *BaseController) Prepare() {
	ctl.StartSession()
	ctl.Data["PageStartTime"] = time.Now()
	user := ctl.GetSession("User")
	if user != nil {
		ctl.User = user.(md.DengluYH)
		if ctl.User.IsAdmin {
			ctl.IsAdmin = true
		}
	} else {
		ctl.User = md.DengluYH{ID: ""}
		return
	}

}
