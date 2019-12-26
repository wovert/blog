package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"blog/models"
)

type OrmController struct {
	beego.Controller
}

type User struct {
	Id int
	Name string
}

func (c *OrmController) ShowOrm() {
	// 1. 连接数据库: 数据库别名(连接多个数据库时使用)
	orm.RegisterDataBase("default",  "mysql","test:test@tcp(47.95.192.23:3306)/test?charset=utf8")

	// 2. 注册表
	orm.RegisterModel(new(User))
	//orm.RegisterModel(new(User), new(Admin))

	// 3. 生成表: 别名，强制更新(true时清空数据表)，创建表过程
	orm.RunSyncdb("default", false, true)

	// 插入数据

	// 查询数据

	// 更新数据

	c.Ctx.WriteString("orm successful！")
}

func (c *OrmController) Insert() {
	// 1. 连接数据库: 数据库别名(连接多个数据库时使用)
	orm.RegisterDataBase("default",  "mysql","test:test@tcp(47.95.192.23:3306)/test?charset=utf8")

	// 2. 注册表
	orm.RegisterModel(new(User))
	//orm.RegisterModel(new(User), new(Admin))

	// 3. 生成表: 别名，强制更新(true时清空数据表)，创建表过程
	orm.RunSyncdb("default", false, true)

	// 4. 获取连接对象
	o := orm.NewOrm()

	// 5. 插入对象
	user := User{}
	user.Name = "zhangsan"

	// 6. 执行插入操作
	_, err := o.Insert(&user)
	if err != nil {
		beego.Info("插入失败：", err)
		return
	}



	c.Ctx.WriteString("Insert successful！")
}

func (c *OrmController) Select() {
	// 4. 获取连接对象
	o := orm.NewOrm()

	// 5. 获取查询列表
	user := models.User{Id:1}
	//user := User{Id:1}

	// 6. 执行查询操作
	err := o.Read(&user)

	if err != nil {
		beego.Info("查询失败：", err)
		return
	}
	beego.Info(user.Name)
	c.Ctx.WriteString("查询成功！")
}

func (c *OrmController) Update() {

	// 获取连接对象
	o := orm.NewOrm()

	// 获取查询列表
	user := models.User{Id:1}

	// 执行查询操作
	err := o.Read(&user)

	if err != nil {
		beego.Info("查询失败：", err)
		return
	}

	user.Name = "lisi"
	_, err = o.Update(&user)
	if err != nil {
		beego.Info("更新数据错误：", err)
		return
	}

	c.Ctx.WriteString("更新成功！")

}
