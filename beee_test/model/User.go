package model

import "github.com/astaxie/beego/orm"

// Model Struct
type User struct {
	Id   int
	Name string `orm:"size(50)"`
}

func init() {
	// set default database
	orm.RegisterDataBase("default", "mysql", "go:123456@tcp(120.79.32.170:3306)/go?charset=utf8mb4", 30)

	// register model
	orm.RegisterModel(new(User))

	// create table
	// orm.RunSyncdb("default", false, true)
}
