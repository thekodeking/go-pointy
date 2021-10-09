package post

import (
	"context"
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
	"strings"
	"time"
)

var Port string

const (
	DBName   = "go_pointy"
	AtlasURI = "mongodb+srv://kode-logger:2Dw5yK7VasbvZP1Y@gpcluster.htie8.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"
)

// Post type, used to define a Post.
type Post struct {
	Id        string    `json:"post_id"`
	UserId    string    `json:"user_id"`
	Caption   string    `json:"caption"`
	ImageURL  string    `json:"image_url"`
	TimeStamp string `json:"time_stamp"`
}

// CreatePost , creates a post by request JSON and adds it to te database ...
func CreatePost(responseWriter http.ResponseWriter, request *http.Request) {
	// if POST, then ...
	if request.Method == http.MethodPost{
		// Logging the Endpoint Hit to the console
		log.Println("[LOG] -> Endpoint Hit: CreatePost")

		// Decode JSON from request and store decode the data ...
		var post Post
		decode := json.NewDecoder(request.Body)
		decode.DisallowUnknownFields()

		err := decode.Decode(&post)
		if err != nil {
			fmt.Fprint(responseWriter, "No Request Found or Incorrect Request!")
			return
		}

		clientOptions := options.Client().ApplyURI(AtlasURI)

		// Connect to the MongoDB and return Client instance ...
		client, err := mongo.Connect(context.TODO(), clientOptions)
		if err != nil {
			log.Fatalf("mongo.Connect() ERROR: %v", err)
		}

		ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
		collection := client.Database(DBName).Collection("post")
		result, _ := collection.InsertOne(ctx, post)
		fmt.Println(result)
		json.NewEncoder(responseWriter).Encode(post)
		fmt.Fprint(responseWriter, "Created New Post Successfully !")
	}
}

// GetPost , returns a specific post ...
func GetPost(responseWriter http.ResponseWriter, request *http.Request) {
	// if GET, then ...
	if request.Method == http.MethodGet{
		// Getting the post-id from the URL Path ...
		pathString := strings.Split(request.URL.Path, "/") // Split the Path by '/' character ...
		queryString := pathString[len(pathString)-2]       // Get the second last string in the path ...
		// if the url is valid, i.e: ".../posts/123" then,
		if queryString == "posts" && pathString[len(pathString)-1] != "" {
			// Setting id as the parameter obtained from the valid path ...
			id := pathString[len(pathString)-1]

			// Logging the Endpoint Hit to the console
			log.Println("[LOG] -> Endpoint Hit: GetPost")

			// Getting data from the database ...
			responseWriter.Header().Set("Content-Type", "application/json")

			// Setting MongoDB Atlas URI to Client ...
			clientOptions := options.Client().ApplyURI(AtlasURI)

			// Connect to the MongoDB and return Client instance ...
			client, err := mongo.Connect(context.TODO(), clientOptions)
			if err != nil {
				log.Fatalf("mongo.Connect() ERROR: %v", err)
			}

			db := client.Database(DBName).Collection("post")

			// Defining an array to store document data ...
			var results []Post

			// Passing the bson.D{{}} as the filter to get all documents in the collection
			cur, err := db.Find(context.TODO(), bson.D{{}})
			if err != nil {
				log.Fatal(err)
			}
			// Finding multiple documents returns a cursor
			// Iterate through the cursor allows us to decode documents one at a time

			for cur.Next(context.TODO()) {
				// Create a value into which the single document can be decoded
				var elem Post
				err := cur.Decode(&elem)
				if err != nil {
					log.Fatal(err)
				}
				results = append(results, elem)
			}

			if err := cur.Err(); err != nil {
				log.Fatal(err)
			}

			//Close the cursor once finished
			cur.Close(context.TODO())

			// Variable to store the index of the Post Struct Object if found
			postIndex := -1
			// Traversing through the data array ...
			for index := 0; index < len(results); index++ {
				if results[index].Id == id {
					postIndex = index
					break
				}
			}

			// if the id is not found in the data ...
			if postIndex == -1 {
				fmt.Fprintf(responseWriter, "Oops, requested data not found !")
			} else {
				// Setting the Header as application/json to let developers download the data as json from the browser
				responseWriter.Header().Set("Content-Type", "application/json")

				// Setting the Header status as: 200
				responseWriter.WriteHeader(http.StatusOK)

				// Encode the obtained Post data as JSON
				encoder := json.NewEncoder(responseWriter)
				encoder.SetIndent("", "  ")        // Indents the JSON ...
				encoder.Encode(results[postIndex]) // Write the JSON data to the browser ...
			}
		} else {
			// Logs error as the input url was incorrect ...
			log.Println("[ERROR] -> QueryIncorrect - Endpoint Hit: GetPost")
			fmt.Fprint(responseWriter, "Seems like the URL is invalid !")
		}
	}
}

// GetAllUserPost , returns all the posts created by a specific user ...
func GetAllUserPost(responseWriter http.ResponseWriter, request *http.Request) {
	// if GET, then ...
	if request.Method == http.MethodGet{
		// Getting the user-id from the URL Path ...
		pathString := strings.Split(request.URL.Path, "/") // Split the Path by '/' character ...
		queryStringPosts := pathString[len(pathString)-3]  // Get the third last string in the path ...
		queryStringUsers := pathString[len(pathString)-2]  // Get the second last string in the path ...
		// if the url is valid, i.e: ".../users/123" then,
		if queryStringPosts == "posts" && queryStringUsers == "users" && pathString[len(pathString)-1] != "" {
			// Setting id as the parameter obtained from the valid path ...
			userId := pathString[len(pathString)-1]

			// Logging the Endpoint Hit to the console
			log.Println("[LOG] -> Endpoint Hit: GetAllUserPost")

			// Get data from the database ...
			responseWriter.Header().Set("Content-Type", "application/json")

			// Setting MongoDB Atlas URI to Client ...
			clientOptions := options.Client().ApplyURI(AtlasURI)

			// Connect to the MongoDB and return Client instance ...
			client, err := mongo.Connect(context.TODO(), clientOptions)
			if err != nil {
				log.Fatalf("mongo.Connect() ERROR: %v", err)
			}

			db := client.Database(DBName).Collection("post")

			// Defining an array to store document data ...
			var results []Post

			// Passing the bson.D{{}} as the filter to get all documents in the collection
			cur, err := db.Find(context.TODO(), bson.D{{}})
			if err != nil {
				log.Fatal(err)
			}
			// Finding multiple documents returns a cursor
			// Iterate through the cursor allows us to decode documents one at a time

			for cur.Next(context.TODO()) {
				// Create a value into which the single document can be decoded
				var elem Post
				err := cur.Decode(&elem)
				if err != nil {
					log.Fatal(err)
				}
				results = append(results, elem)
			}

			if err := cur.Err(); err != nil {
				log.Fatal(err)
			}

			//Close the cursor once finished
			cur.Close(context.TODO())

			// Array to store the Posts of a specific User ...
			var UserPosts []Post

			// Traversing through the data array ...
			for index := 0; index < len(results); index++ {
				if results[index].UserId == userId {
					UserPosts = append(UserPosts, results[index])
				}
			}

			// if the id is not found in the data ...
			if len(UserPosts) == 0 {
				fmt.Fprintf(responseWriter, "Oops, requested data not found !")
			} else {
				// Setting the Header as application/json to let developers download the data as json from the browser
				responseWriter.Header().Set("Content-Type", "application/json")

				// Setting the Header status as: 200
				responseWriter.WriteHeader(http.StatusOK)

				// Encode the obtained Post data as JSON
				encoder := json.NewEncoder(responseWriter)
				encoder.SetIndent("", "  ") // Indents the JSON ...
				encoder.Encode(UserPosts)   // Write the JSON data to the browser ...
			}
		} else {
			// Logs error as the input url was incorrect ...
			log.Println("[ERROR] -> QueryIncorrect - Endpoint Hit: GetPost")
			fmt.Fprint(responseWriter, "Seems like the URL is invalid !")
		}
	}
}
