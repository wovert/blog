package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "wovert.com"
	c.Data["Email"] = "wovert@gmail.com"
	c.TplName = "index.tpl"
}

func (c *MainController) Post() {
	c.Data["Website"] = "wovert.com"
	c.Data["Email"] = "wovert@gmail.com"
	c.TplName = "index.tpl"
}
