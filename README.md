# go-pointy

Go-Pointy is a simple Instagram API Clone, made using GoLang. I had tried my best to not be lazy and finish the tasks, as a beginner to the GoLang programming language, it was pretty hard to conquer the language within a day or so but in the end I was able to produce some quality work. Below is the short summary of the entrance project based test.

## Summary

### Tasks
- [x] Endpoint to Create an User.
- [x] Endpoint to Display a User using id.
- [x] Endpoint to Create a Post.
- [x] Endpoint to Display a Post using id.
- [x] Endpoint to List all the Posts sent by a User.
- [x] Uses MongoDB as Database.
- [x] Quality of Code (Reusable, Naming, ...)
- [x] Password should be securely stored such that they can't be reverse engineered.
- [ ] Make the Server Thread safe.
- [ ] Add Pagination to the List Endpoint.
- [ ] Add Unit Tests

### Endpoints

Before running any of the following commands to test the endpoints, executing the `main.go` file is mandatory to start the server. To start the server, simply run:
```shell
 go run main.go
```

 You can test the following in a browser too but using `curl` command seems simpler and easier😅. Anyways, that following commands can be executed in a command prompt or a Linux Terminal, both works out.

- Create a User: 
  Enter this command in your terminal to send a JSON POST request to the given URL.
  ```shell
  curl -X POST http://localhost:8080/users -d "{\"user_id\":\"123456\", \"name\": \"Bob\", \"email\": \"bob@gmail.com\", \"password\": \"password1234\"}"
  ```

- Display a User by id:
  Enter this command in your terminal to send a JSON POST request to the given URL.
  ```shell
  curl -X GET http://localhost:8080/users/123456 
  ```

- Create a Post:
  Enter this command in your terminal to send a JSON POST request to the given URL.
  ```shell
   curl -X POST http://localhost:8080/posts -d "{\"post_id\":\"1759483975\", \"user_id\":\"123456\", \"caption\": \"Life of Bob\", \"image_url\":\"imgur.com/a/Ffk2322\", \"time_stamp\": \"2021/10/09 21:29:40 2021-10-09 21:29:40.8956683 +0530 IST m=+0.006981001\"}"
  ```

- Display a Post by id:
  Enter this command in your terminal to GET the details of the post by id.
  ```shell
   curl -X GET http://localhost:8080/posts/1759483975 
  ```

- Display all the Posts by a specific User:
  Enter this command in your terminal to GET the details of the posts by a user.
  ```shell
   curl -X GET http://localhost:8080/users/posts/648390
  ```

### Password Encryption

I have used md5 hash encryption to encrypt the password. Of course the hash phrase is present in the code and already I have created functions / methods to encrypt and decrypt passwords so it is also possible to decrypt the passwords, it's not implemented as an Endpoint as it was not mentioned in the tasks.

### Problem Statement

The task is to develop a basic version of the Instagram. You are only required to develop the API for the system. Below are the details.
You are required to Design and Develop an HTTP JSON API capable of the following operations,

#### Create an User
  *	Should be a POST request
  *	Use JSON request body
  *	URL should be ‘/users'
#### Get a user using id
  *	Should be a GET request
  *	Id should be in the url parameter
  *	URL should be ‘/users/<id here>’
#### Create a Post
  *	Should be a POST request
  *	Use JSON request body
  *	URL should be ‘/posts'
#### Get a post using id
  *	Should be a GET request
  *	Id should be in the url parameter
  *	URL should be ‘/posts/<id here>’
#### List all posts of a user
  *	Should be a GET request
  *	URL should be ‘/posts/users/<Id here>'

#### Additional Constraints/Requirements:
-	The API should be developed using Go.
-	MongoDB should be used for storage.
-	Only packages/libraries listed here and here can be used.

#### Users should have the following attributes
*	Id
*	Name
*	Email
*	Password

#### Posts should have the following Attributes. All fields are mandatory unless marked optional:
*	Id
*	Caption
*	Image URL
*	Posted Timestamp





