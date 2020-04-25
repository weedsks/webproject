package blog

import (
	"github.com/astaxie/beego/orm"
	"strconv"
	"webproject/models"
)

type CommentController struct {
	baseController
}

//@router /commentIndex [get]
func (c *CommentController) index() {
	goodsId, err := c.GetInt("id")
	if err != nil {
		c.JsonResult(-1, "缺少参数！"+err.Error(), 1)
	}
	where := map[string]string{}
	where["GoodsId"] = strconv.Itoa(goodsId)
	order := "order desc"
	totalCount, list, _ := models.NewComment().GetListPage(where, order, c.page, c.limit)

	c.JsonResult(0, "请求成功！", totalCount, list)
}

//@router /commentCreate [post]
func (c *CommentController) create() {
	goodsId, err := c.GetInt("id")
	if err != nil {
		c.JsonResult(-1, "缺少参数id！"+err.Error(), 1)
	}

	typeId, err := c.GetInt("type")
	if err != nil {
		c.JsonResult(-1, "缺少参数type！"+err.Error(), 1)
	}

	name := c.GetString("name")
	email := c.GetString("email")
	description := c.GetString("description")

	o := orm.NewOrm()
	info := models.CommentsModel{}
	info.GoodsId = goodsId
	info.Type = typeId
	info.Name = name
	info.Email = email
	info.Description = description
	rid, err := o.Insert(&info)
	if err != nil {
		c.JsonResult(-1, err.Error(), 1)
	}
	c.JsonResult(-1, "评论成功！", 1, rid)

}
