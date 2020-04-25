package admin

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"strconv"
	"time"
	"webproject/models"
)

type GoodsTypeCommon struct {
	baseController
}
//@router /goodsTypeIndex [*]
func (c *GoodsTypeCommon) Index() {
	if c.Ctx.Request.Method == "POST" {
		where := map[string]string{
			//"isVerify": "1",
		}
		typeId,err := c.GetInt("type")
		if err==nil && typeId != 0 {
			where["Type"] = strconv.Itoa(typeId)
		}
		isVerify,err := c.GetInt("isVerify")
		if err==nil {
			where["isVerify"] = strconv.Itoa(isVerify)
		}
		name := c.GetString("name")
		if len(name) > 1 {
			where["name"] = name
		}

		order := "-id"
		var data map[int]models.ListTypeCommonPageResult
		var totalCount int
		var merr error
		//c.JsonResult(0, "请求成功", 0, where)

		totalCount, data, merr = models.NewTypeCommon().GetListPage(where, order, c.page, c.limit)

		if merr != nil && merr != orm.ErrNoRows {
			beego.Error(merr.Error())
		}
		c.JsonResult(0, "请求成功", totalCount, data)
	}
	c.Data["GoodsTypeDesc"] = models.GoodsTypeDesc()

	c.TplName = c.getAdminFix() + "/goodsType-list.html"
}

//@router /goodsTypeDetail [get]
func (c *GoodsTypeCommon) Detail() {
	id, err := c.GetInt("id")
	if err != nil {
		beego.Warning("参数错误->", err)
	}
	info, err := models.NewTypeCommon().GetInfoById(id)
	c.Data["info"] = info
	//角色
	c.Data["GoodsTypeDesc"] = models.GoodsTypeDesc()
	c.TplName = c.getAdminFix() + "/goodsType-edit.html"
}
//@router /goodsTypeUpdate [post]
func (c *GoodsTypeCommon) Update() {
	id, err := c.GetInt("Id")
	if err != nil {
		c.JsonResult(-1, "参数错误ID！", 0)
	}

	oldInfo, err := models.NewTypeCommon().GetInfoById(id)
	if err != nil {
		c.JsonResult(-1, err.Error(), 0)
	}

	name := c.GetString("Name", oldInfo.Name)
	is_verify,_ := c.GetInt("IsVerify", oldInfo.IsVerify)
	order,_ := c.GetInt("Order", oldInfo.Order)
	type_id,_ := c.GetInt("Type", oldInfo.Type)
	icon := c.GetString("Email", oldInfo.Icon)
	if err != nil {
		c.JsonResult(-1, err.Error(), 0)
	}


	info := models.TypeCommonModel{}
	nowtime := time.Now().Format("2006-01-02 15:04:05")

	info.Id = id
	info.Name = name
	info.IsVerify = is_verify
	info.Order = order
	info.Type = type_id
	info.CreatedAt = oldInfo.CreatedAt
	info.Icon = icon
	info.UpdatedAt = nowtime

	o := orm.NewOrm()
	o.Update(&info)
	c.JsonResult(0, "请求成功！", 0)

}
//@router /goodsTypeCreate [*]
func (c *GoodsTypeCommon) Create() {
	if c.Ctx.Request.Method=="POST"{
		name := c.GetString("Name")
		is_verify,err := c.GetInt("IsVerify")
		order,err := c.GetInt("Order")
		type_id,err := c.GetInt("Type")
		icon := c.GetString("Email")
		if err != nil {
			c.JsonResult(-1, err.Error(), 0)
		}

		info := models.TypeCommonModel{}
		nowtime := time.Now().Format("2006-01-02 15:04:05")

		info.Name = name
		info.IsVerify = is_verify
		info.Order = order
		info.Type = type_id
		info.CreatedAt = nowtime
		info.Icon = icon

		o := orm.NewOrm()
		o.Insert(&info)
		c.JsonResult(0, "请求成功！", 0)
		}
	c.Data["GoodsTypeDesc"] = models.GoodsTypeDesc()
	c.TplName=c.getAdminFix() + "/goodsType-add.html"

}
