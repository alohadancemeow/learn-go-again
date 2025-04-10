package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"slices"
	"strconv"
	"strings"
)

// Course represents a training course with its details
type Course struct {
	CourseId    int     `json:"id"`
	CourseName  string  `json:"name"`
	CoursePrice float64 `json:"price"`
	Instructor  string  `json:"instructor"`
}

// CourseList stores all available courses
var CourseList []Course

// init initializes the CourseList with sample data
func init() {
	CourseJSON := `[
	{
		"id": 1,
		"name": "Java",
		"price": 1000,
		"instructor": "Suresh"
	},
	{
		"id": 2	,
		"name": "Python",
		"price": 2000,
		"instructor": "Ramesh"
	},
	{
		"id": 3,
		"name": "Go",
		"price": 3000,
		"instructor": "Rajesh"
	}
	]`

	err := json.Unmarshal([]byte(CourseJSON), &CourseList)
	if err != nil {
		log.Fatal(err)
	}

}

// getNextId determines the next available course ID
// by finding the highest existing ID and adding 1
func getNextId() int {
	highestId := -1
	for _, course := range CourseList {
		if highestId < course.CourseId {
			highestId = course.CourseId
		}
	}
	return highestId + 1
}

// courseHandler processes requests for individual courses by ID
// Supports:
// - GET: Retrieves a specific course by ID
// - PUT: Updates an existing course
func courseHandler(w http.ResponseWriter, r *http.Request) {
	// Extract course ID from URL path
	urlPathSegments := strings.Split(r.URL.Path, "course/")
	ID, err := strconv.Atoi(urlPathSegments[len(urlPathSegments)-1])
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// Find the course with the given ID
	course, listItemIndex := findID(ID)
	if course == nil {
		http.Error(w, "Course not found", http.StatusNotFound)
		return
	}

	switch r.Method {
	case http.MethodGet:
		// Handle GET request: Return the course details
		courseJSON, err := json.Marshal(course)
		if err != nil {
			http.Error(w, "Failed to marshal course", http.StatusInternalServerError)
			return
		}

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(courseJSON)

	case http.MethodPut:
		// Handle PUT request: Update the course details
		var updatedCourse Course
		BodyByte, err := io.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("400 - Bad Request"))
			return
		}
		err = json.Unmarshal(BodyByte, &updatedCourse)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("400 - Bad Request"))
			return
		}
		if updatedCourse.CourseId != ID {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("400 - Bad Request"))
			return
		}
		course = &updatedCourse
		CourseList[listItemIndex] = *course
		w.WriteHeader(http.StatusOK)
		return

	case http.MethodDelete:
		// Handle DELETE request: Remove the course
		CourseList = slices.Delete(CourseList, listItemIndex, listItemIndex+1)
		w.WriteHeader(http.StatusOK)
		return

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("405 - Method Not Allowed"))
		return
	}

}

// findID searches for a course by ID and returns both the course and its index
// Returns nil and 0 if the course is not found
func findID(ID int) (*Course, int) {
	for i, course := range CourseList {
		if course.CourseId == ID {
			return &course, i // Return the course and its index in the slice if found
		}
	}
	return nil, 0
}

// coursesHandler processes HTTP requests for the /courses endpoint
// Supports:
// - GET: Returns all available courses
// - POST: Creates a new course with auto-generated ID
func coursesHandler(w http.ResponseWriter, r *http.Request) {
	courseJSON, err := json.Marshal(CourseList)

	switch r.Method {
	case http.MethodGet:
		// Handle GET request: Return all courses
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("500 - Internal Server Error"))
			return
		}
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(courseJSON)

	case http.MethodPost:
		// Handle POST request: Add a new course
		var newCourse Course
		BodyByte, err := io.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("400 - Bad Request"))
			return
		}
		err = json.Unmarshal(BodyByte, &newCourse)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("400 - Bad Request"))
			return
		}

		newCourse.CourseId = getNextId() // Assign a unique ID to the new course
		CourseList = append(CourseList, newCourse)
		w.WriteHeader(http.StatusCreated)
		return
	}
}

func middlewareHandler(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Incoming request: %s %s", r.Method, r.URL.Path)
		handler.ServeHTTP(w, r)
		fmt.Println("Middleware executed")
	})
}

// main initializes the HTTP server and sets up routing
// - /course/{id}: Handles individual course operations
// - /courses: Handles operations on the course collection
func main() {
	// Register route handlers with the middleware
	http.Handle("/course/", middlewareHandler(http.HandlerFunc(courseHandler)))
	http.Handle("/courses", middlewareHandler(http.HandlerFunc(coursesHandler)))
	// http.HandleFunc("/course/", courseHandler)
	// http.HandleFunc("/courses", coursesHandler)

	// Start the server on port 8080
	http.ListenAndServe(":8080", nil)
}
