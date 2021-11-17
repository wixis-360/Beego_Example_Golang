package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

type Customer struct {
	Id   int
	Name 	 string    `orm:"null"`
	Address  string    `orm:"null"`
	Salary     float64    `orm:"null"`
	//RegTime   time.Time `orm:"auto_now_add;type(datetime)"`
	//RegDate       time.Time `orm:"auto_now_add;type(date)"`


}

func InsertOneCustomer(customer Customer) *Customer {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Customer))
	// get prepared statement
	i, _ := qs.PrepareInsert()
	var cust Customer
	// Insert
	id, err := i.Insert(&customer)
	if err == nil {
		// successfully inserted
		cust = Customer{Id: int(id)}
		err := o.Read(&cust)
		if err == orm.ErrNoRows {
			return nil
		}
	} else {
		return nil
	}
	return &cust
}
func GetAllCustomers() []*Customer {
	o := orm.NewOrm()
	var customers []*Customer
	o.QueryTable(new(Customer)).All(&customers)

	return customers
}
func UpdateCustomer(customer Customer)*Customer  {
	o := orm.NewOrm()
	u := Customer{Id: customer.Id}
	var updatedUser Customer
	// get existing customer
	if o.Read(&u) == nil {

		u = customer
		_, err := o.Update(&u)

		// read updated user
		if err == nil {
			// update successful
			updatedUser = Customer{Id: customer.Id}
			o.Read(&updatedUser)
		}
	}

	return &updatedUser

}
// DeleteUser deletes a Customer

func DeleteCustomer(id int) bool {
	fmt.Println(id)
	user:=FindCustomer(id)
	if user!=nil {
		o := orm.NewOrm()
		o.Delete(&Customer{Id: id})
		return true
	}
	return false
}


func FindCustomer(id int) *Customer {
	fmt.Println(id)
	o := orm.NewOrm()
	user := Customer{Id: id}
	err := o.Read(&user)
	if err != nil {
		return nil
	}
	return &user
}