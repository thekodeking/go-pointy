package user

import (
	"bytes"
	"context"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"io"
	"log"
	"net/http"
	"strings"
	"time"
)

var Port string

const hashKey = "appointy-rocks"

const (
	DBName   = "go_pointy"
	AtlasURI = "mongodb+srv://kode-logger:2Dw5yK7VasbvZP1Y@gpcluster.htie8.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"
)

// User type, used to define a User's data.
type User struct {
	Id       string `json:"id" bson:"id"`
	Name     string `json:"name" bson:"name"`
	Email    string `json:"email" bson:"email"`
	Password string `json:"password" bson:"password"`
}

// GetUser , Gets the requested user data filtered by user-id and returns a JSON response.
func GetUser(responseWriter http.ResponseWriter, request *http.Request) {
	// if GET, then ...
	if request.Method == http.MethodGet{
		// Getting the user-id from the URL Path ...
		pathString := strings.Split(request.URL.Path, "/") // Split the Path by '/' character ...
		queryString := pathString[len(pathString)-2]       // Get the second last string in the path ...
		// if the url is valid, i.e: ".../users/123" then,
		if queryString == "users" && pathString[len(pathString)-1] != "" {
			// Setting id as the parameter obtained from the valid path ...
			id := pathString[len(pathString)-1]
			log.Println(id)
			// Logging the Endpoint Hit to the console
			log.Println("[LOG] -> Endpoint Hit: GetUser")

			// Setting the Header as application/json to let developers download the data as json from the browser
			responseWriter.Header().Set("Content-Type", "application/json")

			// Setting MongoDB Atlas URI to Client ...
			clientOptions := options.Client().ApplyURI(AtlasURI)

			// Connect to the MongoDB and return Client instance ...
			client, err := mongo.Connect(context.TODO(), clientOptions)
			if err != nil {
				log.Fatalf("mongo.Connect() ERROR: %v", err)
			}

			db := client.Database(DBName).Collection("user")

			// Defining an array to store document data ...
			var results []User

			// Passing the bson.D{{}} as the filter to get all documents in the collection
			cur, err := db.Find(context.TODO(), bson.D{{}})
			if err != nil {
				log.Fatal(err)
			}
			// Finding multiple documents returns a cursor
			// Iterate through the cursor allows us to decode documents one at a time

			for cur.Next(context.TODO()) {
				// Create a value into which the single document can be decoded
				var elem User
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
			log.Println("[ERROR] -> QueryIncorrect - Endpoint Hit: GetUser")
			fmt.Fprint(responseWriter, "Seems like the URL is invalid !")
		}
	}
}

// CreateUser , creates a new user from JSON request and adds it to the database ...
func CreateUser(responseWriter http.ResponseWriter, request *http.Request) {
	// if POST, then ...
	if request.Method == http.MethodPost{
		// Logging the Endpoint Hit to the console
		log.Println("[LOG] -> Endpoint Hit: CreateUser")

		// Decode JSON from request and store decode the data ...
		var user User
		decode := json.NewDecoder(request.Body)
		decode.DisallowUnknownFields()

		err := decode.Decode(&user)
		if err != nil {
			fmt.Fprint(responseWriter, "No Request Found or Incorrect Request!")
			return
		}

		user.Password = bytes.NewBuffer(encryptPass([]byte(user.Password), hashKey)).String()
		responseWriter.Header().Set("content-type", "application/json")

		clientOptions := options.Client().ApplyURI(AtlasURI)

		// Connect to the MongoDB and return Client instance ...
		client, err := mongo.Connect(context.TODO(), clientOptions)
		if err != nil {
			log.Fatalf("mongo.Connect() ERROR: %v", err)
		}

		ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
		collection := client.Database(DBName).Collection("user")
		result, _ := collection.InsertOne(ctx, user)
		fmt.Println(result)
		json.NewEncoder(responseWriter).Encode(user)
		fmt.Fprint(responseWriter, "Created New User Successfully !")
	}
}

func generateHash(key string) string {
	hasher := md5.New()
	hasher.Write([]byte(key))
	return hex.EncodeToString(hasher.Sum(nil))
}

func encryptPass(data []byte, passphrase string) []byte {
	block, _ := aes.NewCipher([]byte(generateHash(passphrase)))
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err.Error())
	}
	ciphertext := gcm.Seal(nonce, nonce, data, nil)
	return ciphertext
}

func decryptPass(data []byte, passphrase string) []byte {
	key := []byte(generateHash(passphrase))
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	nonceSize := gcm.NonceSize()
	nonce, ciphertext := data[:nonceSize], data[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		panic(err.Error())
	}
	return plaintext
}
