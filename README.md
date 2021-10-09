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





