package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type Student struct {
	ID    int
	Name  string
	Age   int
	Major string
}

var db *sql.DB

func main() {
	// Open a connection to the MySQL database
	var err error
	db, err = sql.Open("mysql", "root:Locations@12@tcp(localhost:3306)/college")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create a new router
	router := mux.NewRouter()

	// Define API endpoints
	router.HandleFunc("/students", GetStudents).Methods("GET")
	router.HandleFunc("/students/{id}", GetStudent).Methods("GET")
	router.HandleFunc("/students", CreateStudent).Methods("POST")
	router.HandleFunc("/students/{id}", UpdateStudent).Methods("PUT")
	router.HandleFunc("/students/{id}", DeleteStudent).Methods("DELETE")

	// Start the server
	log.Fatal(http.ListenAndServe(":8000", router))
}

func GetStudents(w http.ResponseWriter, r *http.Request) {
	// Fetch all students from the database
	rows, err := db.Query("SELECT * FROM students")
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	students := []Student{}
	for rows.Next() {
		var student Student
		err := rows.Scan(&student.ID, &student.Name, &student.Age, &student.Major)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		students = append(students, student)
	}

	// Convert the students slice to JSON
	jsonData, err := json.Marshal(students)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Set the content type and send the response
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

func GetStudent(w http.ResponseWriter, r *http.Request) {
	// Extract the student ID from the request parameters
	vars := mux.Vars(r)
	studentID, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Fetch the student with the given ID from the database
	var student Student
	err = db.QueryRow("SELECT * FROM students WHERE id = ?", studentID).Scan(&student.ID, &student.Name, &student.Age, &student.Major)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Convert the student struct to JSON
	jsonData, err := json.Marshal(student)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Set the content type and send the response
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

func CreateStudent(w http.ResponseWriter, r *http.Request) {
	// Decode the request body into a Student struct
	var student Student
	err := json.NewDecoder(r.Body).Decode(&student)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Insert the new student into the database
	_, err = db.Exec("INSERT INTO students (name, age, major) VALUES (?, ?, ?)", student.Name, student.Age, student.Major)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func UpdateStudent(w http.ResponseWriter, r *http.Request) {
	// Extract the student ID from the request parameters
	vars := mux.Vars(r)
	studentID, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Decode the request body into a Student struct
	var student Student
	err = json.NewDecoder(r.Body).Decode(&student)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Update the student in the database
	_, err = db.Exec("UPDATE students SET name = ?, age = ?, major = ? WHERE id = ?", student.Name, student.Age, student.Major, studentID)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func DeleteStudent(w http.ResponseWriter, r *http.Request) {
	// Extract the student ID from the request parameters
	vars := mux.Vars(r)
	studentID, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Delete the student from the database
	_, err = db.Exec("DELETE FROM students WHERE id = ?", studentID)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
