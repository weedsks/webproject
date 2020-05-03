package routers

import (
	"github.com/astaxie/beego"
	"webproject/controllers/admin"
	"webproject/controllers/blog"
)

func init() {
	ns := beego.NewNamespace("/admin",
		beego.NSInclude(&admin.IndexController{}),
		//users
		beego.NSInclude(&admin.UserController{}),
		//商品类型
		beego.NSInclude(&admin.GoodsTypeCommon{}),
		//菜单
		beego.NSInclude(&admin.MenusController{}),
		//文章
		beego.NSInclude(&admin.ArticleController{}),
		//mac
		beego.NSInclude(&admin.MacController{}),
	)
	beego.Include(&blog.IndexController{})
	beego.Include(&blog.CommentController{})

	beego.AddNamespace(ns, ns)

}
