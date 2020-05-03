package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["webproject/controllers/blog:CommentController"] = append(beego.GlobalControllerRouter["webproject/controllers/blog:CommentController"],
        beego.ControllerComments{
            Method: "Create",
            Router: `/commentCreate`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["webproject/controllers/blog:CommentController"] = append(beego.GlobalControllerRouter["webproject/controllers/blog:CommentController"],
        beego.ControllerComments{
            Method: "Index",
            Router: `/commentIndex`,
            AllowHTTPMethods: []string{"*"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["webproject/controllers/blog:IndexController"] = append(beego.GlobalControllerRouter["webproject/controllers/blog:IndexController"],
        beego.ControllerComments{
            Method: "Index",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["webproject/controllers/blog:IndexController"] = append(beego.GlobalControllerRouter["webproject/controllers/blog:IndexController"],
        beego.ControllerComments{
            Method: "ArticleDetail",
            Router: `/articleDetail`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
