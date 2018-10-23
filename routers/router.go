package routers

import (
	"Knight/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.IndexController{})
	beego.Router("/api/login", &controllers.LoginControllertroller{}) //登录验证 Get为登出  Post登录
	beego.Router("/api/menu", &controllers.MenuControllertroller{})   //获取菜单
	beego.Router("/api/insert", &controllers.UserControllertroller{}) //用户信息
}
