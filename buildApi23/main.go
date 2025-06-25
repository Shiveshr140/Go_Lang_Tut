package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Model for course file
// As we read as json so better to use lowercase for json tags
// *Author is used to show that Author is a pointer to Author struct
//  Memory Efficiency (No Copying) --> below
// 1. Passing or embedding a struct by value (i.e. without *) copies all its fields.
// 2. Using a pointer avoids copying large structs and is more efficient, especially when nested.


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

// getting data from fake db
// why json.NewEncoder? 
// json.NewEncoder is used to encode Go data structures into JSON format and write them to an io.Writer, such as http.ResponseWriter.
// http.ResponseWriter is also an io.Writer Even though we think of http.ResponseWriter as: “This writes the web response”
// ...under the hood, it implements the io.Writer interface. So you can pass it to anything that needs an io.Writer — including:

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

	fmt.Printf("type of params %t:", params)

	for _, course := range courses{
		if course.CourseId == params["id"] {
            // as we want to send json data
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No course found with given id")
}

// here we getting json until now we are writing json
// why r.Body == {} showing error
// This causes a compile-time error because:
// {} is used to initialize struct literals, not compare objects
// r.Body is an io.ReadCloser interface, not a struct
// Go doesn't allow comparing interface values to struct literals like {}

// why & not *
// You use &course to pass a pointer to course, because:
// json.Decode() needs to modify the passed variable
// It must fill in the fields with data from the JSON
// If you pass course (a value), it would decode into a copy — and you'd lose the data

func createOneCourse(w http.ResponseWriter, r *http.Request){
   fmt.Println("Create one course")
   w.Header().Set("Content-Type", "application/json")

   if r.Body == nil{
	json.NewEncoder(w).Encode("Please send some data")
	return
   }

//    if r.Body == {}{
// 	json.NewEncoder(w).Encode("Please send some data")
// 	return
//   }
   var course Course
   json.NewDecoder(r.Body).Decode(&course)

   if course.isEmpty(){
	 json.NewEncoder(w).Encode("No data inside course")
	 return
   }

   // add unique courseId -> convert into string
   // append course into fake db
   id := rand.Intn(1000)
   fmt.Println("Id is: ", id)
   // string != int
  //  course.CourseId = id
   course.CourseId = strconv.Itoa(id)
   courses = append(courses, course)

   //you can send as json, json.Encoder is itself a return but you can use use return under it to be on safer side it will warn ypu as it is redundant
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
	return

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
	return
}


func main(){

}

