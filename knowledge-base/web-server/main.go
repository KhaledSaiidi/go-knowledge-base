package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
)

type User struct {
	Name string `json:"name"`
}

var userCache = make(map[int]User)
var cacheMutex sync.RWMutex
var nextUserID = 1

func handleRoot(
	w http.ResponseWriter,
	r *http.Request,
) {
	fmt.Fprintf(w, "Hello world")
}

func createUser(
	w http.ResponseWriter,
	r *http.Request,
) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if user.Name == "" {
		http.Error(w, "Name is required", http.StatusBadRequest)
		return
	}
	cacheMutex.Lock()
	// Not Safe len(userCache) reflects only the current number of entries (not the highest ID ever used),
	// deleting items causes IDs to be reused, leading to overwriting existing users.
	// userCache[len(userCache)+1] = user
	id := nextUserID
	userCache[id] = user
	nextUserID++
	cacheMutex.Unlock()
	log.Printf("User created: %s", user.Name)
	w.WriteHeader(http.StatusNoContent)
}

func getUsers(
	w http.ResponseWriter,
	r *http.Request,
) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}
	cacheMutex.RLock()
	user, exists := userCache[id]
	cacheMutex.RUnlock()
	if !exists {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	j, err := json.Marshal(user)
	if err != nil {
		http.Error(w, "Error encoding user data", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

func updateUser(
	w http.ResponseWriter,
	r *http.Request,
) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}
	newName := r.URL.Query().Get("name")
	if newName == "" {
		http.Error(w, "Invalid new Name", http.StatusBadRequest)
		return
	}

	cacheMutex.Lock()
	defer cacheMutex.Unlock()
	user, exists := userCache[id]
	if !exists {
		http.Error(w, "User doesn't Exist", http.StatusNotFound)
		return
	}
	user.Name = newName
	userCache[id] = user
	log.Printf("User %d updated successfully", id)
	w.WriteHeader(http.StatusNoContent)
}

func getAllUsers(
	w http.ResponseWriter,
	r *http.Request,
) {
	cacheMutex.RLock()
	users := userCache
	if users == nil {
		cacheMutex.RUnlock()
		http.Error(w, "No users found", http.StatusNotFound)
		return
	}
	cacheMutex.RUnlock()

	w.Header().Set("Content-Type", "application/json")
	j, err := json.Marshal(users)
	if err != nil {
		http.Error(w, "Error encoding user data", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

func deleteUsers(
	w http.ResponseWriter,
	r *http.Request,
) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}
	cacheMutex.Lock()
	user, exists := userCache[id]
	if !exists {
		cacheMutex.Unlock()
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	delete(userCache, id)
	cacheMutex.Unlock()
	log.Printf("User deleted: %s", user.Name)
	w.WriteHeader(http.StatusNoContent)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handleRoot)

	mux.HandleFunc("POST /users", createUser)
	mux.HandleFunc("GET /users/{id}", getUsers)
	mux.HandleFunc("GET /allusers", getAllUsers)
	mux.HandleFunc("DELETE /users/{id}", deleteUsers)
	mux.HandleFunc("PUT /users/{id}", updateUser)
	fmt.Println("Server listenning to 8080")
	http.ListenAndServe(":8080", mux)
}
