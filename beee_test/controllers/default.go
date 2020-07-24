package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"

	"beee_test/model"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

func (c *MainController) Hello() {
	c.Ctx.WriteString("hello")
}

// 查询数据
func (c *MainController) Sql() {
	id, _ := c.GetInt("id")
	o := orm.NewOrm()
	user := model.User{Id: id}

	err := o.Read(&user)

	if err == orm.ErrNoRows {
		c.Ctx.WriteString("查询不到")
	} else if err == orm.ErrMissPK {
		c.Ctx.WriteString("找不到主键")
	} else {
		c.Ctx.WriteString(user.Name)
	}
}
