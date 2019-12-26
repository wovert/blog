package controllers

import (
	"database/sql"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
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

func (c *IndexController) ShowMySQL() {
	// 打开数据库
	conn, err := sql.Open("mysql", "test:test@tcp(47.95.192.23:3306)/test?charset=utf8")
	if err != nil {
		beego.Info("连接错误：", err)
		return
	}

	// 关闭数据库
	defer conn.Close()

	//创建数据库
	//_, err = conn.Exec("create table userInfo(id int,name varchar(20)) charset=utf8")
	//if err != nil {
	//	beego.Info("创建表错误：", err)
	//	return
	//}
	//c.Ctx.WriteString("创建表成功！")

	// 插入数据
	//_, err = conn.Exec("insert into userInfo(id,name) values(?,?)", 1, "张三")
	//if err != nil {
	//	beego.Info("插入数据错误：", err)
	//	return
	//}

	// 查询数据
	rows, err := conn.Query("select * from userInfo")
	var id int
	for rows.Next() {
		rows.Scan(&id)
		beego.Info(id)
	}
	//c.Ctx.WriteString("插入数据成功！" + string(id))
}