package main

import (
	"fmt"
	"log"
	"os"
	"v1/model"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func init() {
	// *** Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// *** Connect to the database
	// driver
	orm.RegisterDriver(os.Getenv("DB_DRIVER"), orm.DRMySQL)
	// connection string
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_NAME"),
	)
	// register database
	orm.RegisterDataBase("default", "mysql", dsn)
	// debug sql logs
	orm.Debug = true
}

func main() {
	name := "user"
	user := model.User{Name: &name}
	// insert
	_, err := model.AddUser(&user)
	if err != nil {
		return
	}
	// update
	u, err := model.GetUserByName(*user.Name)
	if err != nil {
		return
	}
	name = "astaxie"
	u.Name = &name
	err = model.UpdateUser(u)
	if err != nil {
		return
	}
	// read all
	users, err := model.GetAllUsers()
	if err != nil {
		return
	}
	// print all users
	for id, user := range users {
		fmt.Printf("User %v : %+v \n", id+1, user)
	}

	// delete
	// uncomment the delete to remove user inserted into table
	// err = model.RemoveUser(&user)
	// if err != nil {
	// 	return
	// }
}
