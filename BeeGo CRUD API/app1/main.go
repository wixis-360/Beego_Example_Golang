package main

import (
	"app1/models"
	_ "app1/routers"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/plugins/cors"
	_ "github.com/go-sql-driver/mysql"
	"html/template"
	"net/http"
)

func main() {
	// get database configuration from environment variables
	//dbUser := os.Getenv("db_user")
	//dbPwd := os.Getenv("db_pwd")
	//dbName := "go_user_mgmt"
	beego.ErrorHandler("404", func(rw http.ResponseWriter, r *http.Request) {
		t, _ := template.New("index.html").ParseFiles(beego.BConfig.WebConfig.ViewsPath + "/index.html")
		data := make(map[string]interface{})
		data["content"] = "page not found"
		t.Execute(rw, data)
	})
	orm.RegisterModel(new(models.Customer))
	dbString := "root" + ":" + "ijse" + "@/" + "beeGo" + "?charset=utf8"

	// Register Driver
	orm.RegisterDriver("mysql", orm.DRMySQL)

	// Register default database
	orm.RegisterDataBase("default", "mysql", dbString)

	// db alias
	name := "default"

	// drop table and re-create
	force := false

	// print log
	verbose := true

	// error
	err := orm.RunSyncdb(name, force, verbose)
	if err != nil {
		fmt.Println(err)
	}
	// CORS for https://foo.* origins, allowing:
	// - PUT and PATCH methods
	// - Origin header
	// // - Credentials share
	//beego.InsertFilter("*", beego.BeforeRouter,cors.Allow(&cors.Options{
	//	//AllowOrigins: []string{"https://*.foo.com"},
	//	AllowMethods: []string{"PUT", "GET","DELETE","POST"},
	//	AllowHeaders: []string{"Origin"},
	//	ExposeHeaders: []string{"Content-Length"},
	//	AllowCredentials: true,
	//}))
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		AllowCredentials: true,
	}))
	beego.Run()
}
