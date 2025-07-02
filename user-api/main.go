package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	"github.com/google/uuid"
)

// User struct with the three required attributes
type User struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// In-memory storage for users
var users []User

// Response structure for API responses
type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

// AddUser handles POST /users endpoint
func AddUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Validate required fields
	if user.Name == "" || user.Email == "" {
		response := Response{
			Success: false,
			Message: "Name and email are required",
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	// Generate unique ID
	user.ID = uuid.New().String()

	// Add user to in-memory storage
	users = append(users, user)

	response := Response{
		Success: true,
		Message: "User added successfully",
		Data:    user,
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

// ListUsers handles GET /users endpoint (with optional name filter)
func ListUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Check if name parameter is provided for filtering
	userName := r.URL.Query().Get("name")
	
	if userName != "" {
		// Filter by name (case-insensitive)
		var usersFound []User
		for _, user := range users {
			if strings.EqualFold(user.Name, userName) {
				usersFound = append(usersFound, user)
			}
		}

		if len(usersFound) == 0 {
			response := Response{
				Success: false,
				Message: "User not found",
			}
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(response)
			return
		}

		response := Response{
			Success: true,
			Message: "User found",
			Data:    usersFound,
		}
		json.NewEncoder(w).Encode(response)
		return
	}

	// No name filter, return all users
	response := Response{
		Success: true,
		Message: fmt.Sprintf("Found %d users", len(users)),
		Data:    users,
	}

	json.NewEncoder(w).Encode(response)
}

// GetUserByID handles GET /users/{id} endpoint
func GetUserByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	userID := vars["id"]

	// Find user by ID
	for _, user := range users {
		if user.ID == userID {
			response := Response{
				Success: true,
				Message: "User found",
				Data:    user,
			}
			json.NewEncoder(w).Encode(response)
			return
		}
	}

	// User not found
	response := Response{
		Success: false,
		Message: "User not found",
	}
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(response)
}



// HealthCheck endpoint for basic health monitoring
func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := Response{
		Success: true,
		Message: "API is running",
	}
	json.NewEncoder(w).Encode(response)
}

func main() {
	// Initialize router
	router := mux.NewRouter()

	// Define routes
	router.HandleFunc("/health", HealthCheck).Methods("GET")
	router.HandleFunc("/users", AddUser).Methods("POST")
	router.HandleFunc("/users", ListUsers).Methods("GET")
	router.HandleFunc("/users/{id}", GetUserByID).Methods("GET")

	// Start server
	port := 8080
	fmt.Printf("Server starting on port %d...\n", port)
	fmt.Println("Available endpoints:")
	fmt.Println("  POST   /users     - Add a new user")
	fmt.Println("  GET    /users     - List all users")
	fmt.Println("  GET    /users/{id} - Get user by ID")
	fmt.Println("  GET    /users?name=value - Get user by name")
	fmt.Println("  GET    /health    - Health check")
	fmt.Println()

	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(port), router))
}
