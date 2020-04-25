package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["webproject/controllers/admin:ArticleController"] = append(beego.GlobalControllerRouter["webproject/controllers/admin:ArticleController"],
        beego.ControllerComments{
            Method: "Create",
            Router: `/articleCreate`,
            AllowHTTPMethods: []string{"*"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["webproject/controllers/admin:ArticleController"] = append(beego.GlobalControllerRouter["webproject/controllers/admin:ArticleController"],
        beego.ControllerComments{
            Method: "Detail",
            Router: `/articleDetail`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["webproject/controllers/admin:ArticleController"] = append(beego.GlobalControllerRouter["webproject/controllers/admin:ArticleController"],
        beego.ControllerComments{
            Method: "Index",
            Router: `/articleIndex`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["webproject/controllers/admin:ArticleController"] = append(beego.GlobalControllerRouter["webproject/controllers/admin:ArticleController"],
        beego.ControllerComments{
            Method: "Update",
            Router: `/articleUpdate`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["webproject/controllers/admin:ArticleController"] = append(beego.GlobalControllerRouter["webproject/controllers/admin:ArticleController"],
        beego.ControllerComments{
            Method: "GetQiNiuToken",
            Router: `/getQiNiuToken`,
            AllowHTTPMethods: []string{"*"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["webproject/controllers/admin:GoodsTypeCommon"] = append(beego.GlobalControllerRouter["webproject/controllers/admin:GoodsTypeCommon"],
        beego.ControllerComments{
            Method: "Create",
            Router: `/goodsTypeCreate`,
            AllowHTTPMethods: []string{"*"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["webproject/controllers/admin:GoodsTypeCommon"] = append(beego.GlobalControllerRouter["webproject/controllers/admin:GoodsTypeCommon"],
        beego.ControllerComments{
            Method: "Detail",
            Router: `/goodsTypeDetail`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["webproject/controllers/admin:GoodsTypeCommon"] = append(beego.GlobalControllerRouter["webproject/controllers/admin:GoodsTypeCommon"],
        beego.ControllerComments{
            Method: "Index",
            Router: `/goodsTypeIndex`,
            AllowHTTPMethods: []string{"*"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["webproject/controllers/admin:GoodsTypeCommon"] = append(beego.GlobalControllerRouter["webproject/controllers/admin:GoodsTypeCommon"],
        beego.ControllerComments{
            Method: "Update",
            Router: `/goodsTypeUpdate`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["webproject/controllers/admin:IndexController"] = append(beego.GlobalControllerRouter["webproject/controllers/admin:IndexController"],
        beego.ControllerComments{
            Method: "Index",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["webproject/controllers/admin:IndexController"] = append(beego.GlobalControllerRouter["webproject/controllers/admin:IndexController"],
        beego.ControllerComments{
            Method: "Login",
            Router: `/login`,
            AllowHTTPMethods: []string{"*"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["webproject/controllers/admin:IndexController"] = append(beego.GlobalControllerRouter["webproject/controllers/admin:IndexController"],
        beego.ControllerComments{
            Method: "Logout",
            Router: `/logout`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["webproject/controllers/admin:IndexController"] = append(beego.GlobalControllerRouter["webproject/controllers/admin:IndexController"],
        beego.ControllerComments{
            Method: "Welcome",
            Router: `/welcome`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["webproject/controllers/admin:MenusController"] = append(beego.GlobalControllerRouter["webproject/controllers/admin:MenusController"],
        beego.ControllerComments{
            Method: "Create",
            Router: `/menuCreate`,
            AllowHTTPMethods: []string{"*"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["webproject/controllers/admin:MenusController"] = append(beego.GlobalControllerRouter["webproject/controllers/admin:MenusController"],
        beego.ControllerComments{
            Method: "Detail",
            Router: `/menuDetail`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["webproject/controllers/admin:MenusController"] = append(beego.GlobalControllerRouter["webproject/controllers/admin:MenusController"],
        beego.ControllerComments{
            Method: "Index",
            Router: `/menuIndex`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["webproject/controllers/admin:MenusController"] = append(beego.GlobalControllerRouter["webproject/controllers/admin:MenusController"],
        beego.ControllerComments{
            Method: "Update",
            Router: `/menuUpdate`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["webproject/controllers/admin:UserController"] = append(beego.GlobalControllerRouter["webproject/controllers/admin:UserController"],
        beego.ControllerComments{
            Method: "Create",
            Router: `/userCreate`,
            AllowHTTPMethods: []string{"*"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["webproject/controllers/admin:UserController"] = append(beego.GlobalControllerRouter["webproject/controllers/admin:UserController"],
        beego.ControllerComments{
            Method: "Detail",
            Router: `/userDetail`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["webproject/controllers/admin:UserController"] = append(beego.GlobalControllerRouter["webproject/controllers/admin:UserController"],
        beego.ControllerComments{
            Method: "Index",
            Router: `/userIndex`,
            AllowHTTPMethods: []string{"*"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["webproject/controllers/admin:UserController"] = append(beego.GlobalControllerRouter["webproject/controllers/admin:UserController"],
        beego.ControllerComments{
            Method: "Update",
            Router: `/userUpdate`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
