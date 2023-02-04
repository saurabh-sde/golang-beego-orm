# golang-beego-orm
Beego ORM Sample project for database connection using MySQL driver and performing basic CRUD operations 

## Advantages and Usage of Beego ORM for Databases
Features:
- easy for usage, simple CRUD operation
- auto join with relation table
- Raw SQL query / mapper without orm model
- Secure and multiple driver connection (like MySQL, postgres, etc)

# Beego ORM implementation in Go lang

To Run project, follow these steps:

1. Create new database, and configure .env variables
2. Install Go lang if not installed then Get all project dependencies by running the following command in terminal:
> ``` go mod tidy ```

3. Check main.go file for basic details
4. Run the code by command in terminal:
> ```go run main.go```
5. Check logs in terminal.
6. Run the select query in user table to check results

For more implementation details check this out:
> https://pkg.go.dev/github.com/beego/beego/v2/client/orm#section-readme

<br>

# Detailed Documentations for Database usage
> https://github.com/beego/beedoc/tree/master/en-US/mvc/model

<br>

---

# Project Structure 
1. main.go - main program file
2. Model - database models and queries
3. db.sql - Database setup
4. .env - all environment variables
5. go.mod - dependencies

> NOTE - Project is not completely on MVC framework. Idea is to only integrate Beego ORM and perform CRUD operations. You can create your project with handle and controller, then use Postman for API interactions

