package main

import "Student_API/model"

type StudentList struct {
	StudentList []model.Student `json:"students"`
}
