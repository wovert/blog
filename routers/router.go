package routers

import (
	"blog/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/index", &controllers.IndexController{})
	beego.Router("/list", &controllers.IndexController{}, "get:GetList")
	beego.Router("/show", &controllers.IndexController{}, "get,post,Put,Delete:GetShow")

	// 精确匹配优先级高
	beego.Router("/level", &controllers.IndexController{}, "*:GetShow;get:GetLevel")

	// 正则路由
	beego.Router("/detail/?:id", &controllers.IndexController{}, "get:GetDetail")

	// 匹配/detail/ 失败
	beego.Router("/detail/:id", &controllers.IndexController{}, "get:GetDetail")

	// 自定义正则匹配  数字
	beego.Router("/detail/:id([0-9]+)", &controllers.IndexController{}, "get:GetDetail")

	// 匹配文件
	beego.Router("/download/*.*", &controllers.IndexController{}, "get:GetFile")

	// 全匹配
	beego.Router("/download/test/*", &controllers.IndexController{}, "get:GetFullMatch")

}
