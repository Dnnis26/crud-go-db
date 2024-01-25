package models

import (
	"database/sql"
	"fmt"
	"mysqldb/db"
)

type Students struct {
	Idstudents int64    `json:"id"`
	Username   string   `json:"username"`
	LastName   string   `json:"lastname"`
	Average    float64  `json:"average"`
	Project    []string `json:"project"`
}

type ListStudents []Students

const StudentSchema string = `CREATE TABLE students (
	idstudents INT NOT NULL AUTO_INCREMENT,
	username VARCHAR(45) NOT NULL,
	lastname VARCHAR(45) NOT NULL,
	average DECIMAL NULL,
	create_data TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	PRIMARY KEY (id))`

func NewStudents(username, lastname string, average float64, id int) *Students {
	student := &Students{Username: username, LastName: lastname, Average: Average(GetProjectNota(id)), Project: GetProjectByStudents(id)}
	return student
}
func CreateStudent(username, lastname string, average float64, id int) *Students {
	student := NewStudents(username, lastname, average, id)
	student.Save()
	return student
}
func Exec(query string, args ...interface{}) (sql.Result, error) {
	result, err := db.Exec(query, args...)
	if err != nil {
		fmt.Println(err)
	}
	return result, err
}
func Query(query string, args ...interface{}) (*sql.Rows, error) {
	rows, err := db.Query(query, args...)
	if err != nil {
		fmt.Println(err)
	}
	return rows, err
}
func (student *Students) Insert() {
	sql := "INSERT students SET username=?, lastname=?, average=?"
	result, _ := db.Exec(sql, student.Username, student.LastName, student.Average)
	student.Idstudents, _ = result.LastInsertId()
}
func ListStudent() ListStudents {
	sql := "SELECT idstudents, username, lastname FROM students"
	users := ListStudents{}
	rows, _ := db.Query(sql)

	for rows.Next() {
		user := Students{}
		rows.Scan(&user.Idstudents, &user.Username, &user.LastName)
		users = append(users, user)
	}
	return users
}

func GetStudent(id int) *Students {

	user := NewStudents("", "", Average(GetProjectNota(id)), id)

	sql := "SELECT idstudents, username, lastname, average FROM students WHERE idstudents=?"
	rows, _ := db.Query(sql, id)

	for rows.Next() {
		rows.Scan(&user.Idstudents, &user.Username, &user.LastName, &user.Average)
	}
	return user
}
func (student *Students) update() {
	sql := "UPDATE students SET username=?, lastname=?, average=? WHERE idstudents=?"
	db.Exec(sql, student.Username, student.LastName, student.Average, student.Idstudents)
}
func (student *Students) Save() {
	if student.Idstudents == 0 {
		student.Insert()
	} else {
		student.update()
	}
}
func (student *Students) Delete() {
	sql := "DELETE FROM students WHERE idstudents=?"
	db.Exec(sql, student.Idstudents)
}
