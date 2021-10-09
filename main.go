/*
	Developer: N.Krishna Raj
 */

package main

import (
	"fmt"
	"go-pointy/appointy/post"
	"go-pointy/appointy/user"
	"log"
	"net/http"
	"strconv"
)

var Port = 8080

// homePage, endpoint for main-page
func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the MainPage!")
	log.Println("[LOG] -> Endpoint Hit: MainPage")
}

// handleRequests, handles all the requests made to the server
func handleRequests() {
	address := ":" + strconv.Itoa(Port)
	http.HandleFunc("/", homePage)
	http.HandleFunc("/users", user.CreateUser)
	http.HandleFunc("/users/", user.GetUser)
	http.HandleFunc("/posts", post.CreatePost)
	http.HandleFunc("/posts/", post.GetPost)
	http.HandleFunc("/posts/users/", post.GetAllUserPost)
	log.Fatal(http.ListenAndServe(address, nil))
}

func main() {
	// Setting the Server Ports to other Packages ...
	user.Port = strconv.Itoa(Port)
	post.Port = strconv.Itoa(Port)
	// Starting the Server ...
	handleRequests()
}
