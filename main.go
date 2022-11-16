package main

import (
	"Student_API/controllers"
	"Student_API/model"
	"github.com/labstack/echo"
)

// global variable across the package
var e = echo.New()

func main() {
	//Connect Database...
	model.CreateDatabase()

	//Routes
	e.POST("/students", controllers.CreateStudent)
	e.GET("/students", controllers.GetStudents)
	e.GET("/students/:id", controllers.GetStudent)
	e.PATCH("/students/:id", controllers.UpdateStudent)
	e.DELETE("students/:id", controllers.DeleteStudent)

	//start localhost
	err := e.Start("localhost:8080")
	if err != nil {
		return
	}
}
