package routers

import (
	"app1/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
	//ns := beego.NewNamespace("/api/v1",
	//	beego.NSNamespace("/customer",
	//		beego.NSRouter("/", &controllers.CustomerController{}, "post:AddNewCustomer"),
	//		beego.NSRouter("/", &controllers.CustomerController{}, "get:GetAllCustomers"),
	//		beego.NSRouter("/", &controllers.CustomerController{}, "put:UpdateCustomer"),
	//		beego.NSRouter("/", &controllers.CustomerController{}, "delete:DeleteCustomer"),
	//		beego.NSRouter("/:id", &controllers.CustomerController{}, "get:SearchCustomer"),
	//
	//		),
	//	)
	//
	//beego.AddNamespace(ns)
	beego.Router("/api/v1/customer/add", &controllers.CustomerController{}, "post:AddNewCustomer")
	beego.Router("/api/v1/customer/all", &controllers.CustomerController{}, "get:GetAllCustomers")
	beego.Router("/api/v1/customer/search/:id", &controllers.CustomerController{}, "get:SearchCustomer")
	beego.Router("/api/v1/customer/update", &controllers.CustomerController{}, "put:UpdateCustomer")
	beego.Router("/api/v1/customer/delete/:id", &controllers.CustomerController{}, "delete:DeleteCustomer")

}
