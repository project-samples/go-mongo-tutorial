# go-mongo-tutorial

##How to run
####Clone the repository
```shell
git clone https://github.com/go-tutorials/go-mongo-rest-api.git
cd go-mongo-rest-api
```
####To run the application
```shell
go run main.go
```
##API Design
###Common HTTP methods
- GET: retrieve a representation of the resource
- POST: create a new resource
- PUT: update the resource
- PATCH: perform a partial update of a resource
- DELETE: delete a resource
###API design for users
- Resource: users
####Get all users
#####Request
- GET /users
#####Response
```json
[
    {
        "id": "spiderman",
        "username": "spiderman",
        "email": "peter.parker@gmail.com",
        "phone": "0987654321",
        "dateOfBirth": "1962-08-25T16:59:59.999Z"
    },
    {
        "id": "wolverine",
        "username": "wolverine",
        "email": "james.howlett@gmail.com",
        "phone": "0987654321",
        "dateOfBirth": "1974-11-16T16:59:59.999Z"
    }
]
```
####Get an user by id
#####Request
- GET /users/:id
- Sample
```url
GET /users/wolverine
```
#####Response
```json
{
    "id": "wolverine",
    "username": "wolverine",
    "email": "james.howlett@gmail.com",
    "phone": "0987654321",
    "dateOfBirth": "1974-11-16T16:59:59.999Z"
}
```
####Create a new user
####Request
- POST /users
- Request body
```json
{
    "id": "spiderman",
    "username": "spiderman",
    "email": "peter.parker@gmail.com",
    "phone": "0987654321",
    "dateOfBirth": "1962-08-25T16:59:59.999Z"
}
```
#####Response
1. *Return*: number
- 1: success
- 0: duplicate key
- -1: error
2. *Sample*
```json
1
```
####Update a new user by id
####Request
- PUT /users/:id
- Request URL
```url
PUT /users/wolverine
```
- Request body
```json
{
    "username": "wolverine",
    "email": "james.howlett@gmail.com",
    "phone": "0987654321",
    "dateOfBirth": "1974-11-16T16:59:59.999Z"
}
```
#####Response
1. *Return*: number
- 1: success
- 0: no data found
- -1: error
1. *Sample*
```json
1
```
####Delete a new user by id
####Request
- DELETE /users/:id
- Sample Request URL
```url
DELETE /users/wolverine
```
#####Response
1. *Return*: number
- 1: success
- 0: no data found
- -1: error
2. *Sample*
```json
1
```
