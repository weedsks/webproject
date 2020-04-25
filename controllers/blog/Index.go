package blog

import (
	"github.com/astaxie/beego"
	"webproject/models"
)

// CMS API
type IndexController struct {
	baseController
}
//@router / [get]
func (c *IndexController) Index()  {
	where := map[string]string{
		"is_verify" : "1",
	}
	name := c.GetString("name")
	if len(name) > 1 {
		where["name"] = name
	}
	order := "g.id desc"

	totalCount, data, merr := models.NewGoods().GetListPage(where, order, c.page, c.limit)
	if merr != nil {
		beego.Notice(merr)
	}
	c.SetPaginator(c.limit,totalCount)
	c.Data["list"] = data
	typeCommon_where := map[string]string{
		"isVerify": "1",
	}
	_, typeCommon, _ := models.NewTypeCommon().GetListPage(typeCommon_where,"-order", 1, 30)
	c.Data["typeCommon"] =  typeCommon
	//c.JsonResult(0,"请求成功", 1, data)
	c.TplName = c.getHomeFix() + "/index.html"


	//c.TplName = c.getHomeFix() + "/index.html"
}

//@router /articleDetail [get]
func (c *IndexController) ArticleDetail() {
	id, err := c.GetInt("id")
	if err != nil {
		beego.Warning("参数错误->", err)
	}
	info, err := models.NewArticle().GetInfoById(id)
	c.Data["info"] = info
	typeComon_arr, _ := models.NewTypeCommon().GetAll(1)
	//c.JsonResult(0,"weeds",0,info)
	c.Data["typeCommon"] = typeComon_arr
	c.TplName = c.getHomeFix() + "/detail.html"
}

