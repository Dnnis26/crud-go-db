package handlers

import (
	"encoding/json"
	"fmt"
	"mysqldb/db"
	"mysqldb/models"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetStudents(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")

	db.Connect()
	users := models.ListStudent()
	db.Close()

	output, _ := json.Marshal(users)
	fmt.Fprintln(rw, string(output))
}
func GetStudent(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	studentID, _ := strconv.Atoi(vars["id"])

	db.Connect()

	users := models.GetStudent(studentID)
	average := models.Average(models.GetProjectNota(studentID))
	users.Average = average
	users.Save()

	db.Close()

	output, _ := json.Marshal(users)
	fmt.Fprintln(rw, string(output))
}
func CreateStudent(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")

	model := models.Students{}

	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&model); err != nil {
		fmt.Println(rw, http.StatusUnprocessableEntity)
	} else {
		db.Connect()
		model.Save()
		db.Close()
	}

	output, _ := json.Marshal(model)
	fmt.Fprintln(rw, string(output))
}
func UpdateStudent(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")

	model := models.Students{}

	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&model); err != nil {
		fmt.Println(rw, http.StatusUnprocessableEntity)
	} else {
		db.Connect()
		model.Save()
		db.Close()
	}

	output, _ := json.Marshal(model)
	fmt.Fprintln(rw, string(output))
}
func DeleteStudent(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")

	//Obtener ID
	vars := mux.Vars(r)
	studentID, _ := strconv.Atoi(vars["id"])

	db.Connect()
	users := models.GetStudent(studentID)
	users.Delete()
	db.Close()

	output, _ := json.Marshal(users)
	fmt.Fprintln(rw, string(output))
}

// ------------------------- Projects --------------------------------
func GetProjects(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")

	db.Connect()
	users := models.ListProject()
	db.Close()

	output, _ := json.Marshal(users)
	fmt.Fprintln(rw, string(output))
}
func NewProject(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")

	project := models.Projects{}
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&project)
	if err != nil {
		fmt.Println(rw, http.StatusUnprocessableEntity)
	} else {
		db.Connect()
		project.Insertproject()
		db.Close()
	}

	output, _ := json.Marshal(project)
	fmt.Fprintln(rw, string(output))

}

func DeleteProject(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	projectID, _ := strconv.Atoi(vars["id"])

	db.Connect()
	users := models.DeleteProject(projectID)
	db.Close()

	output, _ := json.Marshal(users)
	fmt.Fprintln(rw, string(output))
}

func UpdateProject(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")

	model := models.Projects{}

	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&model); err != nil {
		fmt.Println(rw, http.StatusUnprocessableEntity)
	} else {
		db.Connect()
		model.Save()
		db.Close()
	}

	output, _ := json.Marshal(model)
	fmt.Fprintln(rw, string(output))
}
