package admin

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"time"
	"webproject/models"
)

type MenusController struct {
	baseController
}
//@router /menuIndex [get]
func (c *MenusController) Index() {
	data := models.NewMenus().GetAll()
	c.Data["list"] = data
	c.TplName = c.getAdminFix() + "/menus-list.html"
}
//@router /menuDetail [get]
func (c *MenusController) Detail() {
	id, err := c.GetInt("id")
	if err != nil {
		beego.Warning("参数错误->", err)
	}
	info, err := models.NewMenus().GetInfoById(id)
	c.Data["info"] = info

	parents := models.NewMenus().GetParents()
	c.Data["parents"] = parents
	c.TplName = c.getAdminFix() + "/menus-edit.html"
}
//@router /menuUpdate [post]
func (c *MenusController) Update() {
	id, err := c.GetInt("Id")
	if err != nil {
		c.JsonResult(-1, "参数错误ID！", 0)
	}

	oldInfo, err := models.NewMenus().GetInfoById(id)
	if err != nil {
		c.JsonResult(-1, err.Error(), 0)
	}

	ParentId, _ := c.GetInt("ParentId", oldInfo.ParentId)
	Title := c.GetString("Title", oldInfo.Title)
	Icon := c.GetString("Icon", oldInfo.Icon)
	Url := c.GetString("Url", oldInfo.Url)
	Order, err := c.GetInt("Order", oldInfo.Order)
	if err != nil {
		Order = 0
	}
	IsDel, err := c.GetInt("IsDel", oldInfo.IsDel)

	if len(Title)<1 {
		c.JsonResult(-1, "请填写栏目名！", 0)
	}

	menu := models.MenusModel{}
	nowtime := time.Now().Format("2006-01-02 15:04:05")

	menu.Id = id
	menu.ParentId = ParentId
	menu.Title = Title
	menu.Url = Url
	menu.Order = Order
	menu.IsDel = IsDel
	menu.UpdatedAt = nowtime
	menu.CreatedAt = oldInfo.CreatedAt
	menu.Icon = Icon

	o := orm.NewOrm()
	o.Update(&menu)

	c.JsonResult(0, "请求成功！", 0)

}
//@router /menuCreate [*]
func (c *MenusController) Create() {
	if c.Ctx.Request.Method=="POST" {
		Title := c.GetString("Title")
		Url := c.GetString("Url", "")
		Icon := c.GetString("Icon", "")
		Order, err := c.GetInt("Order", 0)
		if err != nil {
			Order = 0
		}
		ParentId, err := c.GetInt("ParentId", 0)
		if err != nil {
			ParentId = 0
		}
		if len(Title)<1 {
			c.JsonResult(-1, "请填写栏目名！", 0)
		}

		menu := models.MenusModel{}
		nowtime := time.Now().Format("2006-01-02 15:04:05")

		if ParentId==0 {
			Order = Order * 100
		}else{
			parentInfo,err := models.NewMenus().GetInfoById(ParentId)
			if err!=nil {
				c.JsonResult(-1, "父级查找出错！", 0)
			}
			Order = parentInfo.Order+Order
		}

		menu.ParentId = ParentId
		menu.Title = Title
		menu.Url = Url
		menu.Order = Order
		menu.IsDel = 0
		menu.UpdatedAt = nowtime
		menu.CreatedAt = nowtime
		menu.Icon = Icon

		o := orm.NewOrm()
		_, err = o.Insert(&menu)
		if err != nil {
			beego.Error(err)
			c.JsonResult(-1, "请求失败！", 0)
		}

		c.JsonResult(0, "请求成功！", 0)
	}
	ParentId, err := c.GetInt("ParentId", 0)
	if err != nil {
		ParentId = 0
	}
	c.Data["ParentId"] = ParentId

	parents := models.NewMenus().GetParents()
	c.Data["parents"] = parents
	c.TplName = c.getAdminFix()+"/menus-add.html";
}

