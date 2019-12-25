package controllers

import (
	"github.com/astaxie/beego"
)

type IndexController struct {
	beego.Controller
}


func (c *IndexController) Post() {
	c.Data["name"] = "wovert.com"
	c.Data["email"] = "wovert@gmail.com"
	c.TplName = "post.tpl"
}

func (c *IndexController) GetList() {
	c.Data["name"] = "wovert.com"
	c.Data["email"] = "wovert@gmail.com"
	c.TplName = "list.tpl"
}

func (c *IndexController) GetShow() {
	c.Data["name"] = "wovert.com"
	c.Data["email"] = "wovert@gmail.com"
	c.TplName = "show.tpl"
}

func (c *IndexController) GetLevel() {
	c.Data["name"] = "wovert.com"
	c.Data["email"] = "wovert@gmail.com"
	c.TplName = "level.tpl"
}

func (c *IndexController) GetDetail() {
	c.Data["id"] = c.GetString(":id") // 获取 GET 参数
	c.TplName = "detail.tpl"
}

func (c *IndexController) GetFile() {
	path := c.GetString(":path")
	ext := c.GetString(":ext")
	beego.Info("path=", path, "ext=", ext)
	c.TplName = "file.tpl"
}

func (c *IndexController) GetFullMatch() {
	value := c.GetString(":splat")
	beego.Info("splat=", value)
	c.Data["id"] = value
	c.TplName = "fullmatch.tpl"
}