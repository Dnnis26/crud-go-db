package models

import (
	"mysqldb/db"

	_ "github.com/go-sql-driver/mysql"
)

type Projects struct {
	Idprojects  int64  `json:"id"`
	Projectname string `json:"projectname"`
	Nota        int    `json:"nota"`
	Studentsid  int    `json:"studentid"`
}
type ListProjects []Projects

const ProjectSchema string = `CREATE TABLE projects (
	idprojects INT NOT NULL AUTO_INCREMENT,
	projectname VARCHAR(45) NOT NULL,
	nota INT NOT NULL,
	id_students INT NOT NULL,
	create_data TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	PRIMARY KEY (id),
	KEY studentsid_fk_idx (id_students),
	ADD CONSTRAINT studentsid_fk
  		FOREIGN KEY (id_students)
  		REFERENCES students (id)
  		ON DELETE CASCADE
  		ON UPDATE CASCADE
	)`

func NewProject(projectname string, nota int, idstudents int) *Projects {
	project := &Projects{Projectname: projectname, Nota: nota, Studentsid: idstudents}
	return project
}
func CreateProject(projectname string, nota int, idStudent int) *Projects {
	project := NewProject(projectname, nota, idStudent)
	project.Insertproject()
	return project
}
func GetProjectByStudents(Studentsid int) []string {
	sql := "SELECT projectname FROM projects WHERE id_students=?"

	rows, _ := db.Query(sql, Studentsid)
	defer rows.Close()

	var project []string
	for rows.Next() {
		var projects string
		err := rows.Scan(&projects)
		if err != nil {
			return nil
		}

		project = append(project, projects)
	}
	return project

}
func ListProject() ListProjects {
	sql := "SELECT idprojects, projectname, nota FROM projects"
	users := ListProjects{}
	rows, _ := db.Query(sql)

	for rows.Next() {
		user := Projects{}
		rows.Scan(&user.Idprojects, &user.Projectname, &user.Nota)
		users = append(users, user)
	}
	return users
}
func (project *Projects) Insertproject() {
	sql := "INSERT projects SET projectname=?, nota=?, id_students=?"
	db.Exec(sql, project.Projectname, project.Nota, project.Studentsid)

}
func GetProjectNota(studentID int) []float64 {
	sql := "SELECT nota FROM projects WHERE id_students=?"

	rows, err := db.Query(sql, studentID)
	if err != nil {
		return nil
	}
	var grades []float64
	for rows.Next() {
		var grade float64
		err := rows.Scan(&grade)
		if err != nil {
			return nil
		}
		grades = append(grades, grade)
	}
	rows.Close()

	return grades
}
func Average(nota []float64) float64 {
	total := 0.0000
	nota = append(nota, nota...)
	for _, nota := range nota {
		total += nota
	}
	promedio := total / float64(len(nota))
	return promedio
}
func DeleteProject(projectID int) error {
	sql := "DELETE FROM projects WHERE idprojects=?"
	_, err := db.Query(sql, projectID)
	if err != nil {
		return err
	}
	db.Close()
	return nil
}
func (project *Projects) update() {
	sql := "UPDATE projects SET projectname=?, nota=? WHERE idprojects=?"
	db.Exec(sql, project.Projectname, project.Nota, project.Idprojects)
}
func (project *Projects) Save() {
	if project.Idprojects == 0 {
		project.Insertproject()
	} else {
		project.update()
	}
}
