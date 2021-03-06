# go-mongodb-CRUD-REST-API
This is a simple CRUD operation on mongodb using golang. It's using http requests to operate.  
Endpoints are:
```
/ping
/users/find
/users/create
/users/delete
/users/update
```
## /ping  
Send a GET request.  
Returns a string:  
```
pong
```
This is only to check if the server is running.
## /users/find  
Send a GET request(only email as parameter). Ex)
```
localhost:8080/users/find?email=abhisheksharma_2@zapbuild.com
```
Returns a json:
```
{
  "_id": "<user id>",
  "name": "<user name>",
  "email": "<user email>",
  "password": "",
}
```
## /users/create
Send a POST request with a json body. Ex)
```
{
  "name": "AbhiShek Sharma",
  "email": "abhisheksharma_2@zapbuild.com",
  "password": "abcd1234"
}
```
Returns a json:  
```
{
  "_id": "<user id>",
  "name": "<user name>",
  "eamil": "<user email>",
  "password": ""
}
```
This endpoint creates a document in the *users* collection of the *users* database.  
## /users/update  
Send a GET request. Ex)
```
localhost:8080/users/update?email=abhisheksharma_2@zapbuild.com&field=name&value=Abhi
```
Returns a json:
```
{
  "_id": "<user id>",
  "name": "<user name>",
  "eamil": "<user email>",
  "password": ""
}
```
This endpoint updates the field *name* of the user with specified email to the value of *Abhi* and returns the updated user profile.  
## /users/delete
Send a GET request. Ex)
```
localhost:8080/users/delete?email=abhisheksharma_2@zapbuild.com
```
Returns a json:
```
{
  "isRemoved": true
}
```
This endpoint deletes the user based on the user eamil.  
## Errors
All the endpoints return an error in json format if something goes wrong. Ex)
```
{
   "Message": <message>,
   "Status": <status>,
   "Error": <error>
}
```
### Used:
lang: **go**  
mux: **github.com/gin-gonic/gin**  
mongodb driver: **go.mongodb.org/mongo-driver**  
### How to run app:
First, get libs and source code.
```
go get github.com/gin-gonic/gin
```
```
go get go.mongodb.org/mongo-driver
```
```
go get github.com/abshzap17797/go-mongo-CRUD-web/
```
Then change directory to *<$GOPATH>/src/github.com/abshzap17797
/go-mongo-CRUD-web/* and run:
```
go run main.go
```
It runs on port 8080 by default.
