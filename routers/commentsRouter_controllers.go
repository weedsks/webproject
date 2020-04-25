package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["webproject/controllers:ArticleController"] = append(beego.GlobalControllerRouter["webproject/controllers:ArticleController"],
        beego.ControllerComments{
            Method: "Create",
            Router: `/articleCreate`,
            AllowHTTPMethods: []string{"*"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["webproject/controllers:ArticleController"] = append(beego.GlobalControllerRouter["webproject/controllers:ArticleController"],
        beego.ControllerComments{
            Method: "Detail",
            Router: `/articleDetail`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["webproject/controllers:ArticleController"] = append(beego.GlobalControllerRouter["webproject/controllers:ArticleController"],
        beego.ControllerComments{
            Method: "Index",
            Router: `/articleIndex`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["webproject/controllers:ArticleController"] = append(beego.GlobalControllerRouter["webproject/controllers:ArticleController"],
        beego.ControllerComments{
            Method: "Update",
            Router: `/articleUpdate`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["webproject/controllers:ArticleController"] = append(beego.GlobalControllerRouter["webproject/controllers:ArticleController"],
        beego.ControllerComments{
            Method: "GetQiNiuToken",
            Router: `/getQiNiuToken`,
            AllowHTTPMethods: []string{"*"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["webproject/controllers:GoodsTypeCommon"] = append(beego.GlobalControllerRouter["webproject/controllers:GoodsTypeCommon"],
        beego.ControllerComments{
            Method: "Create",
            Router: `/goodsTypeCreate`,
            AllowHTTPMethods: []string{"*"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["webproject/controllers:GoodsTypeCommon"] = append(beego.GlobalControllerRouter["webproject/controllers:GoodsTypeCommon"],
        beego.ControllerComments{
            Method: "Detail",
            Router: `/goodsTypeDetail`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["webproject/controllers:GoodsTypeCommon"] = append(beego.GlobalControllerRouter["webproject/controllers:GoodsTypeCommon"],
        beego.ControllerComments{
            Method: "Index",
            Router: `/goodsTypeIndex`,
            AllowHTTPMethods: []string{"*"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["webproject/controllers:GoodsTypeCommon"] = append(beego.GlobalControllerRouter["webproject/controllers:GoodsTypeCommon"],
        beego.ControllerComments{
            Method: "Update",
            Router: `/goodsTypeUpdate`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["webproject/controllers:IndexController"] = append(beego.GlobalControllerRouter["webproject/controllers:IndexController"],
        beego.ControllerComments{
            Method: "Index",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["webproject/controllers:IndexController"] = append(beego.GlobalControllerRouter["webproject/controllers:IndexController"],
        beego.ControllerComments{
            Method: "Login",
            Router: `/login`,
            AllowHTTPMethods: []string{"*"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["webproject/controllers:IndexController"] = append(beego.GlobalControllerRouter["webproject/controllers:IndexController"],
        beego.ControllerComments{
            Method: "Logout",
            Router: `/logout`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["webproject/controllers:IndexController"] = append(beego.GlobalControllerRouter["webproject/controllers:IndexController"],
        beego.ControllerComments{
            Method: "Welcome",
            Router: `/welcome`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["webproject/controllers:MenusController"] = append(beego.GlobalControllerRouter["webproject/controllers:MenusController"],
        beego.ControllerComments{
            Method: "Create",
            Router: `/menuCreate`,
            AllowHTTPMethods: []string{"*"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["webproject/controllers:MenusController"] = append(beego.GlobalControllerRouter["webproject/controllers:MenusController"],
        beego.ControllerComments{
            Method: "Detail",
            Router: `/menuDetail`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["webproject/controllers:MenusController"] = append(beego.GlobalControllerRouter["webproject/controllers:MenusController"],
        beego.ControllerComments{
            Method: "Index",
            Router: `/menuIndex`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["webproject/controllers:MenusController"] = append(beego.GlobalControllerRouter["webproject/controllers:MenusController"],
        beego.ControllerComments{
            Method: "Update",
            Router: `/menuUpdate`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["webproject/controllers:UserController"] = append(beego.GlobalControllerRouter["webproject/controllers:UserController"],
        beego.ControllerComments{
            Method: "Create",
            Router: `/userCreate`,
            AllowHTTPMethods: []string{"*"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["webproject/controllers:UserController"] = append(beego.GlobalControllerRouter["webproject/controllers:UserController"],
        beego.ControllerComments{
            Method: "Detail",
            Router: `/userDetail`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["webproject/controllers:UserController"] = append(beego.GlobalControllerRouter["webproject/controllers:UserController"],
        beego.ControllerComments{
            Method: "Index",
            Router: `/userIndex`,
            AllowHTTPMethods: []string{"*"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["webproject/controllers:UserController"] = append(beego.GlobalControllerRouter["webproject/controllers:UserController"],
        beego.ControllerComments{
            Method: "Update",
            Router: `/userUpdate`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
