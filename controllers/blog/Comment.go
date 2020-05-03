package blog

import (
	"github.com/astaxie/beego/orm"
	"strconv"
	"webproject/models"
)

type CommentController struct {
	baseController
}

//@router /commentIndex [*]
func (c *CommentController) Index() {
	goodsId, err := c.GetInt("id")
	if err != nil {
		c.JsonResult(-1, "缺少参数！"+err.Error(), 1)
	}
	where := map[string]string{}
	where["GoodsId"] = strconv.Itoa(goodsId)
	where["IsVerify"] = "1"
	order := "-order"
	totalCount, list, _ := models.NewComment().GetListPage(where, order, c.page, 3)

	c.JsonResult(0, "请求成功！", totalCount, list)
}

//@router /commentCreate [post]
func (c *CommentController) Create() {
	goodsId, err := c.GetInt("id")
	if err != nil {
		c.JsonResult(-1, "缺少参数id！"+err.Error(), 1)
	}

	typeId, err := c.GetInt("type")
	if err != nil {
		c.JsonResult(-1, "缺少参数type！"+err.Error(), 1)
	}

	name := c.GetString("name")
	if len(name)<1 {
		c.JsonResult(-1, "请输入你的名字！", 1)
	}
	email := c.GetString("email")
	description := c.GetString("description")
	if len(description)<1 {
		c.JsonResult(-1, "请输入评论内容！", 1)
	}

	o := orm.NewOrm()
	info := models.CommentsModel{}
	info.GoodsId = goodsId
	info.Type = typeId
	info.Name = name
	info.Email = email
	info.Description = description
	info.IsVerify = 1
	rid, err := o.Insert(&info)
	if err != nil {
		c.JsonResult(-1, err.Error(), 1)
	}
	c.JsonResult(0, "评论成功！", 1, rid)

}
