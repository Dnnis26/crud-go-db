package main

import (
	"log"
	"mysqldb/db"
	"mysqldb/handlers"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {
	db.Connect()

	mux := mux.NewRouter()

	//Endpoinds
	mux.HandleFunc("/api/student/", handlers.GetStudents).Methods("GET")                 // Listar
	mux.HandleFunc("/api/student/{id:[0-9]+}", handlers.GetStudent).Methods("GET")       // Obtener
	mux.HandleFunc("/api/student/", handlers.CreateStudent).Methods("POST")              // Agregar
	mux.HandleFunc("/api/student/{id:[0-9]+}", handlers.UpdateStudent).Methods("PUT")    // Editar
	mux.HandleFunc("/api/student/{id:[0-9]+}", handlers.DeleteStudent).Methods("DELETE") // Eliminar

	mux.HandleFunc("/api/project/", handlers.GetProjects).Methods("GET")                 // Listar
	mux.HandleFunc("/api/project/", handlers.NewProject).Methods("POST")                 // Agregar
	mux.HandleFunc("/api/project/{id:[0-9]+}", handlers.DeleteProject).Methods("DELETE") // Eliminar
	mux.HandleFunc("/api/project/{id:[0-9]+}", handlers.UpdateProject).Methods("PUT")    // Editar

	log.Fatal(http.ListenAndServe(":3000", mux))

	db.Close()

}
