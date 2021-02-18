package routers

import (
	"app/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/admin", &controllers.MainController{}, "get:Admin")
	beego.Router("/page/:id", &controllers.PageController{}, "get:Get")
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/users",
			beego.NSRouter("/register", &controllers.UserController{}, "post:Register"),
			beego.NSRouter("/update/:id", &controllers.UserController{}, "put:Update"),
			beego.NSRouter("/login", &controllers.UserController{}, "post:Login"),
			beego.NSRouter("/logout", &controllers.UserController{}, "post:Logout"),
			beego.NSRouter("/info", &controllers.UserController{}, "get:GetInfo"),
			beego.NSRouter("/passwd", &controllers.UserController{}, "put:Passwd"),
			beego.NSRouter("/requestpasswd", &controllers.UserController{}, "put:RequestPasswd"),
			beego.NSRouter("/uploads", &controllers.UserController{}, "post:Uploads"),
			beego.NSRouter("/downloads", &controllers.UserController{}, "get:Downloads"),
		),
		beego.NSNamespace("/pages",
			beego.NSRouter("/list", &controllers.PageController{}, "get:List"),
			beego.NSRouter("/", &controllers.PageController{}, "post:New"),
			beego.NSRouter("/:id", &controllers.PageController{}, "get:GetPage"),
			beego.NSRouter("/update/:id", &controllers.PageController{}, "put:Update"),
		),
		beego.NSNamespace("/orders",
			beego.NSRouter("/list", &controllers.OrderController{}, "get:List"),
			//beego.NSRouter("/", &controllers.OrderController{}, "get:List"),
			beego.NSRouter("/:product_id", &controllers.OrderController{}, "post:CreateOrUpdate"),
			beego.NSRouter("/:order_id/", &controllers.OrderController{}, "post:ApplyPromoCode"),
		),
		beego.NSNamespace("/products",
			beego.NSRouter("/list", &controllers.ProductController{}, "get:List"),
			beego.NSRouter("/", &controllers.ProductController{}, "get:GetByQuery"),
			beego.NSRouter("/", &controllers.ProductController{}, "post:New"),
			beego.NSRouter("/update/:id", &controllers.ProductController{}, "put:Update"),
			beego.NSRouter("/:id", &controllers.ProductController{}, "get:Get"),
		),
		beego.NSNamespace("/menus",
			beego.NSRouter("/list", &controllers.MenuController{}, "get:List"),
			beego.NSRouter("/", &controllers.MenuController{}, "get:GetByQuery"),
			beego.NSRouter("/", &controllers.MenuController{}, "post:New"),
			beego.NSRouter("/update/:id", &controllers.MenuController{}, "put:Update"),
			beego.NSRouter("/:id", &controllers.MenuController{}, "get:Get"),
		),
		beego.NSNamespace("/promocodes",
			beego.NSRouter("/list", &controllers.PromoCodeController{}, "get:List"),
			beego.NSRouter("/", &controllers.PromoCodeController{}, "post:New"),
			beego.NSRouter("/update/:id", &controllers.PromoCodeController{}, "put:Update"),
		),
		beego.NSNamespace("/deliverycity",
			beego.NSRouter("/list", &controllers.DeliveryCityController{}, "get:List"),
			beego.NSRouter("/:id", &controllers.DeliveryCityController{}, "get:Get"),
			beego.NSRouter("/", &controllers.DeliveryCityController{}, "post:New"),
			beego.NSRouter("/update/:id", &controllers.DeliveryCityController{}, "put:Update"),
		),
		beego.NSNamespace("/roles",
			beego.NSRouter("/:id", &controllers.RoleController{}, "get:GetOne;put:Put;delete:Delete"),
			beego.NSRouter("/", &controllers.RoleController{}, "get:GetAll;post:Post"),
			beego.NSRouter("/auth", &controllers.RoleController{}, "post:Auth"),
		),
		beego.NSNamespace("/agency",
			beego.NSRouter("/update/", &controllers.AgencyController{}, "put:Update"),
			beego.NSRouter("/info/", &controllers.AgencyController{}, "get:Get"),
		),
	)
	beego.AddNamespace(ns)
}
