package models

import "github.com/astaxie/beego/orm"

type User struct {
	Id int
	Name string
}

func init() {
	// 1. 连接数据库: 数据库别名(连接多个数据库时使用)
	orm.RegisterDataBase("default",  "mysql","test:test@tcp(47.95.192.23:3306)/test?charset=utf8")

	// 2. 注册表
	orm.RegisterModel(new(User))
	//orm.RegisterModel(new(User), new(Admin))

	// 3. 生成表: 别名，强制更新(true时清空数据表)，创建表过程
	orm.RunSyncdb("default", false, true)
}

