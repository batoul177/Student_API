package controllers

import (
	"Student_API/model"
	"github.com/labstack/echo"
	"gopkg.in/go-playground/validator.v9"
	"net/http"
	"strconv"
)

var v = validator.New()

type CreateStudentInput struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func CreateStudent(context echo.Context) error {
	//validate input
	var input CreateStudentInput
	err := context.Bind(&input)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	err1 := v.Struct(input)
	if err1 != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err1.Error())
	}
	//create student and register in database
	student := model.Student{FirstName: input.FirstName, LastName: input.LastName}
	model.Db.Create(&student)
	context.Response().WriteHeader(http.StatusCreated)
	return nil
}

func GetStudents(context echo.Context) error /*[]model.Student */ {
	var students []model.Student
	model.Db.Find(&students)

	err := context.JSON(http.StatusOK, func(students []model.Student) []model.Student {
		if len(students) == 0 {
			return []model.Student{}
		}
		var studentArray []model.Student
		for _, student := range students {
			studentArray = append(studentArray, student)
		}
		return studentArray
	})
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return nil
}

func GetSingleStudent(context echo.Context) (model.Student, error) {
	var student model.Student

	err := model.Db.Where("id = ?", context.Param("id")).First(&student).Error
	if err != nil {
		return model.Student{}, echo.NewHTTPError(http.StatusNotFound, "Status not found")
	}
	return model.Student{ID: student.ID, FirstName: student.FirstName, LastName: student.LastName}, nil
}

func GetStudent(context echo.Context) error {
	id, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	if student, err := GetSingleStudent(id); err != nil {
		return err
	} else {
		return context.JSON(http.StatusOK, student)
	}
}

type UpdateStudentInput struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func UpdateStudent(c echo.Context) error {
	var student model.Student
	if err := model.Db.Where("id = ?", c.Param("id")).First(&student).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound)
	}
	if err := customError(c.Bind(&student)); err != nil {
		return err
	}
	if err1 := customError(v.Struct(student)); err1 != nil {
		return err1
	}

	var input UpdateStudentInput
	updatedStudent := model.Student{FirstName: input.FirstName, LastName: input.LastName}
	model.Db.Model(&student).Updates(&updatedStudent)
	c.Response().WriteHeader(http.StatusCreated)
	return nil
}

func customError(err error) error {
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return nil
}

func DeleteStudent(c echo.Context) error {
	var student model.Student
	err := customError(model.Db.Where("id = ?", c.Param("id")).First(&student).Error)
	if err != nil {
		return err
	}
	model.Db.Delete(&student)
	c.Response().WriteHeader(http.StatusOK)
	return nil
}
