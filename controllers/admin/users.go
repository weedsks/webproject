package admin

import (
	"github.com/Echosong/beego_blog/util"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"time"
	"webproject/models"
)

type UserController struct {
	baseController
}
//@router /userIndex [*]
func (c *UserController) Index() {
	if c.Ctx.Request.Method == "POST" {
		where := map[string]string{
			"IsDel": "0",
		}
		name := c.GetString("name")
		if len(name) > 1 {
			where["name"] = name
		}
		start := c.GetString("start")
		end := c.GetString("end")

		if len(start) > 1 && len(end) > 1 {
			where["Created_at__gte"] = start
			where["Created_at__lte"] = end
		}

		order := "-id"
		var data map[int]models.ListUserPageResult
		var totalCount int
		var merr error

		totalCount, data, merr = models.NewUser().GetListPage(where, order, c.page, c.limit)

		if merr != nil && merr != orm.ErrNoRows {
			beego.Error(merr.Error())
		}
		c.JsonResult(0, "请求成功", totalCount, data)
	}

	c.TplName = c.getAdminFix() + "/member-list.html"
}
//@router /userDetail [get]
func (c *UserController) Detail() {
	id, err := c.GetInt("id")
	if err != nil {
		beego.Warning("参数错误->", err)
	}
	user, err := models.NewUser().GetInfoById(id)
	c.Data["info"] = user
	//角色
	role_list := models.NewRole().GetAll()
	c.Data["RoleList"] = role_list

	c.TplName = c.getAdminFix() + "/member-edit.html"
}
//@router /userUpdate [post]
func (c *UserController) Update() {
	id, err := c.GetInt("Id")
	if err != nil {
		c.JsonResult(-1, "参数错误ID！", 0)
	}

	oldUser, err := models.NewUser().GetInfoById(id)
	if err != nil {
		c.JsonResult(-1, err.Error(), 0)
	}

	name := c.GetString("Name", oldUser.Name)
	mobil := c.GetString("Mobile", oldUser.Mobile)
	email := c.GetString("Email", oldUser.Email)
	password := c.GetString("password")
	repassword := c.GetString("repassword")
	role_id, err := c.GetInt("RoleId", oldUser.RoleId)
	is_del, err := c.GetInt("IsDel", oldUser.IsDel)
	if err != nil {
		c.JsonResult(-1, "请选择角色！", 0)
	}

	if password != repassword {
		c.JsonResult(-1, "2次输入密码不一致！", 0)
	}

	user := models.UsersModel{}
	nowtime := time.Now().Format("2006-01-02 15:04:05")

	user.Id = id
	user.Name = name
	user.Mobile = mobil
	user.Email = email
	user.Updated_at = nowtime
	user.Created_at = oldUser.Created_at
	user.Salt = oldUser.Salt
	user.IsDel = is_del
	user.RoleId = role_id

	if len(password) > 1 {
		user.Password = util.Md5(password + beego.AppConfig.String("pwdsalt"))
	} else {
		user.Password = oldUser.Password
	}

	o := orm.NewOrm()
	o.Update(&user)

	c.JsonResult(0, "请求成功！", 0)

}
//@router /userCreate [*]
func (c *UserController) Create() {
	if c.Ctx.Request.Method=="POST"{
		name := c.GetString("Name")
		mobil := c.GetString("Mobile")
		email := c.GetString("Email")
		role_id, err := c.GetInt("RoleId")
		password := c.GetString("password")
		repassword := c.GetString("repassword")
		if len(name) < 1 {
			c.JsonResult(-1, "用户名不能为空！", 0)
		}
		if len(mobil) < 1 {
			c.JsonResult(-1, "手机号不能为空！", 0)
		}
		if len(password) < 1 {
			c.JsonResult(-1, "密码不能为空！", 0)
		}
		if err != nil {
			c.JsonResult(-1, "请选择角色！", 0)
		}
		if password != repassword {
			c.JsonResult(-1, "2次输入密码不一致！", 0)
		}

		user := models.UsersModel{}
		nowtime := time.Now().Format("2006-01-02 15:04:05")

		user.Name = name
		user.Mobile = mobil
		user.Email = email
		user.Updated_at = nowtime
		user.Created_at = nowtime
		user.Salt = ""
		user.IsDel = 0
		user.RoleId = role_id
		user.Password = util.Md5(password + beego.AppConfig.String("pwdsalt"))

		o := orm.NewOrm()
		_, err = o.Insert(&user)

		if err != nil {
			beego.Error(err)
			c.JsonResult(-1, "请求失败！", 0)
		}

		c.JsonResult(0, "请求成功！", 0)
	}
	//角色
	role_list := models.NewRole().GetAll()
	c.Data["RoleList"] = role_list
	c.TplName=c.getAdminFix() + "/member-add.html"

}
