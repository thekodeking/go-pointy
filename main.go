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

// Curl Command to test Post: // curl -X POST http://localhost:8080/posts -d "{\"post_id\":\"96453\", \"user_id\":\"648390\", \"caption\": \"Go Lang by Google\", \"image_url\":\"imgur.com/a/Ffkso2\", \"time_stamp\": \"2021/10/09 21:29:40 2021-10-09 21:29:40.8956683 +0530 IST m=+0.006981001\"}"
// Curl Command to test User: curl -X POST http://localhost:8080/users -d "{\"user_id\":\"964533\", \"name\": \"Alex\", \"email\": \"alex@gmail.com\", \"password\": \"google\"}"