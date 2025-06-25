package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)



type Course struct{
	CourseName string 	`json:"coursename"`
	CourseId string `json:"courseid"`
	Price int `json:"price"`
	Author *Author `json:"author"`
}

type Author struct{
	FullName string `json:"fullname"`
	Website string `json:"website"`
}

// fake db -> gonna be slice with Course struct
var courses []Course

// helper/middleware
func(c *Course) isEmpty() bool{
    // return c.CourseId == "" && c.CourseName == ""
    return c.CourseName == ""
}


// controller -> where all logic goes inside route handler just like serve home in MOD 21
func ServeHome(w http.ResponseWriter, r *http.Request){
   w.Write([]byte("<h1>Welcome to Kodnest</h1>"))
}


func getAllCourse(w http.ResponseWriter, r *http.Request){
   fmt.Println("Get all courses")
   w.Header().Set("Content-Type", "application/json")
   json.NewEncoder(w).Encode((courses))
}

// Just to remind r is reader i.e request reader so it will read params and say id and get the particular course then we write response.
// why pointer *http.Request? so that we manipulate request
func getOneCourse(w http.ResponseWriter, r *http.Request){
	fmt.Println("Get one course")
	var params = mux.Vars(r)
	fmt.Println("Params are: ", params)

	fmt.Printf("type of params %T:", params)

	for _, course := range courses{
		if course.CourseId == params["id"] {
            // as we want to send json data
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No course found with given id")
}



func createOneCourse(w http.ResponseWriter, r *http.Request){
   fmt.Println("Create one course")
   w.Header().Set("Content-Type", "application/json")

   if r.Body == nil{
	json.NewEncoder(w).Encode("Please send some data")
	return
   }

   var course Course
   json.NewDecoder(r.Body).Decode(&course)

   if course.isEmpty(){
	 json.NewEncoder(w).Encode("No data inside course")
	 return
   }

   // check if course is already present opr not
   for _, sub := range courses{
	 if sub.CourseName == course.CourseName{
		json.NewEncoder(w).Encode("Crourse is already created!")
	 }
   }

   // add unique courseId -> convert into string
   // append course into fake db
   id := rand.Intn(1000)
   fmt.Println("Id is: ", id)
   // string != int
  //  course.CourseId = id
   course.CourseId = strconv.Itoa(id)
   courses = append(courses, course)

   json.NewEncoder(w).Encode("Course created successfully")
   json.NewEncoder(w).Encode(course)
//    return
}

func updateOneCourse(w http.ResponseWriter, r *http.Request){
	fmt.Println("Update one course")
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, course := range courses{
		if course.CourseId == params["id"]{
           courses = append(courses[:index], courses[index+1:]...)
		   // as these are not real scenario we need to declare course that will be comming from request body in real body
		   var course Course
		   json.NewDecoder(r.Body).Decode(&course)
		   course.CourseId = params["id"]
		   courses = append(courses, course)
		   json.NewEncoder(w).Encode("Course updated successfully")
		   json.NewEncoder(w).Encode(course)
		   return
		}
	}

	// if id is not found
	json.NewEncoder(w).Encode("Course not found with given id")
	// return

}

func deleteOneCourse(w http.ResponseWriter, r *http.Request){
	fmt.Println("Delete one course")
    w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, course := range courses{
		if course.CourseId == params["id"]{
			courses = append(courses[:index], courses[index+1:]...)
			json.NewEncoder(w).Encode("course deleted successfully")
			return
		}
	}
	json.NewEncoder(w).Encode("Course not found with given id")
	// return
}

// mux will handle routing
func main(){
 fmt.Println("Welcome to Kodnest API")
 // create a new router
 r := mux.NewRouter()

 // seeding
 // why &Author, Author is pointer we need to pass a reference
 courses = append(courses, Course{CourseName: "ReactJS", CourseId: "1", Price: 299, Author: &Author{FullName: "John Doe", Website: "johndoe.com"}})
 courses = append(courses, Course{CourseName: "Angular", CourseId: "2", Price: 199, Author: &Author{FullName: "Rohit Doe", Website: "rohitdoe.com"}})
 
 // routes
 r.HandleFunc("/", ServeHome).Methods("GET")
 r.HandleFunc("/courses", getAllCourse).Methods("GET")
 r.HandleFunc("/course/{id}", getOneCourse).Methods("GET")
 r.HandleFunc("/course", createOneCourse).Methods("POST")
 r.HandleFunc("/course/{id}", updateOneCourse).Methods("PUT")
 r.HandleFunc("/course/{id}", deleteOneCourse).Methods("DELETE")
 

// listening
 log.Fatal(http.ListenAndServe(":4000", r))
}

