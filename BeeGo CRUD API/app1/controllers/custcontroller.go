package controllers

import (
	"app1/models"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"strconv"
)

type CustomerController struct {
	beego.Controller
}

//Add  Customer  function

func (customer *CustomerController) AddNewCustomer() {
	var u models.Customer
	f, _ := strconv.ParseFloat(customer.GetString("salary"), 64)
	u.Name = customer.GetString("name")
	u.Address = customer.GetString("address")
	u.Salary = f
	json.Unmarshal(customer.Ctx.Input.RequestBody, &u)
	cust := models.InsertOneCustomer(u)
	if cust != nil {
		//customer.Ctx.ResponseWriter.WriteHeader(201)
		customer.Data["json"] = cust
		customer.ServeJSON()

	} else {
		customer.Ctx.ResponseWriter.WriteHeader(500)
		customer.Data["json"] = "Customer Save Fail!!"
		customer.ServeJSON()
	}

}

//GetAll  Customer  function

func (customer *CustomerController) GetAllCustomers() {
	cust := models.GetAllCustomers()
	if len(cust) > 0 {
	//	customer.Ctx.ResponseWriter.WriteHeader(200)
		customer.Data["json"] = cust
		customer.ServeJSON()

	} else {
		customer.Ctx.ResponseWriter.WriteHeader(404)
		customer.Data["json"] = "Customer's Not Exists!"
		customer.ServeJSON()
	}

}

//Find  Customer  function

func (customer *CustomerController) SearchCustomer() {
	// get the id from query string
	id, _ := strconv.Atoi(customer.Ctx.Input.Param(":id"))
	//id, _ := strconv.ParseInt(customer.GetString("id"), 0, 64)

	// generate response

	cust := models.FindCustomer(id)
	if cust != nil {
		customer.Data["json"] = cust
		customer.ServeJSON()
	} else {
		customer.Ctx.ResponseWriter.WriteHeader(404)
		customer.Data["json"] = "Customer Not Found!"
		customer.ServeJSON()

	}

}

//Update Customer  function

func (customer *CustomerController) UpdateCustomer() {
	//var u models.Customer
	var u models.Customer
	e, _ := strconv.ParseInt(customer.GetString("id"), 0, 64)
	f, _ := strconv.ParseFloat(customer.GetString("salary"), 64)
	u.Name = customer.GetString("name")
	u.Address = customer.GetString("address")
	u.Salary = f
	u.Id = int(e)
	fmt.Println(u.Id, u.Name, u.Address, u.Salary)
	json.Unmarshal(customer.Ctx.Input.RequestBody, &u)

	user := models.UpdateCustomer(u)
	if user != nil {
		//customer.Ctx.ResponseWriter.WriteHeader(200)
		customer.Data["json"] = "Customer Update Successfully!"
		customer.ServeJSON()

	} else {
		customer.Ctx.ResponseWriter.WriteHeader(300)
		customer.Data["json"] = "Customer Update Fail!"
		customer.ServeJSON()

	}

}

//Delete function

func (customer *CustomerController) DeleteCustomer() {
	id, _ := strconv.Atoi(customer.Ctx.Input.Param(":id"))
	//id, _ := strconv.ParseInt(customer.GetString("id"), 0, 64)
	fmt.Println(id)
	// delete user
	deleted := models.DeleteCustomer(id)
	// generate response
	if deleted {
		//customer.Ctx.ResponseWriter.WriteHeader(200)
		customer.Data["json"] = "Customer Delete Successfully!"
		customer.ServeJSON()

	} else {
		customer.Ctx.ResponseWriter.WriteHeader(300)
		customer.Data["json"] = "Customer Delete Fail!!"
		customer.ServeJSON()

	}

}
