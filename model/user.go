package model

import (
	"fmt"
	"time"

	"github.com/astaxie/beego/orm"
)

type User struct {
	Id        int64     `orm:"column(id)" json:"id"`
	Name      *string   `orm:"column(name)" json:"name"`
	Age       *int      `orm:"column(age)" json:"age"`
	Gender    *string   `orm:"column(gender)" json:"gender"`
	CreatedAt time.Time `orm:"column(created_at);auto_now_add"`
	UpdatedAt time.Time `orm:"column(updated_at);auto_now"`
}

func init() {
	orm.RegisterModel(new(User))
}

func AddUser(user *User) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(&user)
	if err != nil {
		fmt.Println("Err in insert: ", err)
	}
	return
}

func GetAllUsers() ([]*User, error) {
	o := orm.NewOrm()
	var users []*User
	_, err := o.QueryTable("user").All(&users)
	if err != nil {
		fmt.Println("Err in GetAllUsers: ", err)
		return nil, err
	}
	return users, err
}

func GetUserByName(name *string) (*User, error) {
	o := orm.NewOrm()
	var user *User
	err := o.QueryTable("user").Filter("name", name).One(&user)
	if err != nil {
		fmt.Println("Err in GetUserByName: ", err)
		return nil, err
	}
	return user, err
}

func UpdateUser(user *User) (err error) {
	o := orm.NewOrm()
	_, err = o.Update(&user)
	if err != nil {
		fmt.Println("Err in UpdateUser: ", err)
	}
	return
}

func RemoveUser(user *User) (err error) {
	o := orm.NewOrm()
	_, err = o.Delete(&user)
	if err != nil {
		fmt.Println("Err in RemoveUser")
	}
	return
}
