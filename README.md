# go-mongo-tutorial

## How to run
#### Clone the repository
```shell
git clone https://github.com/go-tutorials/go-mongo-rest-api.git
cd go-mongo-rest-api
```

#### To run the application
```shell
go run main.go
```

## API Design
### Common HTTP methods
- GET: retrieve a representation of the resource
- POST: create a new resource
- PUT: update the resource
- PATCH: perform a partial update of a resource
- DELETE: delete a resource

## API design for health check
To check if the service is available.
#### *Request:* GET /health
#### *Response:*
```json
{
    "status": "UP",
    "details": {
        "mongo": {
            "status": "UP"
        }
    }
}
```

## API design for users
#### *Resource:* users

### Get all users
#### *Request:* GET /users
#### *Response:*
```json
[
    {
        "id": "spiderman",
        "username": "peter.parker",
        "email": "peter.parker@gmail.com",
        "phone": "0987654321",
        "dateOfBirth": "1962-08-25T16:59:59.999Z"
    },
    {
        "id": "wolverine",
        "username": "james.howlett",
        "email": "james.howlett@gmail.com",
        "phone": "0987654321",
        "dateOfBirth": "1974-11-16T16:59:59.999Z"
    }
]
```

### Get one user by id
#### *Request:* GET /users/:id
```shell
GET /users/wolverine
```
#### *Response:*
```json
{
    "id": "wolverine",
    "username": "james.howlett",
    "email": "james.howlett@gmail.com",
    "phone": "0987654321",
    "dateOfBirth": "1974-11-16T16:59:59.999Z"
}
```

### Create a new user
#### *Request:* POST /users 
```json
{
    "id": "wolverine",
    "username": "james.howlett",
    "email": "james.howlett@gmail.com",
    "phone": "0987654321",
    "dateOfBirth": "1974-11-16T16:59:59.999Z"
}
```
#### *Response:* 1: success, 0: duplicate key, -1: error
1: success, 0: duplicate key, -1: error
```json
1
```

### Update one user by id
#### *Request:* PUT /users/:id
```shell
PUT /users/wolverine
```
```json
{
    "username": "james.howlett",
    "email": "james.howlett@gmail.com",
    "phone": "0987654321",
    "dateOfBirth": "1974-11-16T16:59:59.999Z"
}
```
#### *Response:* Return *a number* 1: success, 0: duplicate key, -1: error
```json
1
```

### Delete a new user by id
#### *Request:* DELETE /users/:id
```shell
DELETE /users/wolverine
```
#### *Response:* Return *a number*
1: success, 0: duplicate key, -1: error
```json
1
```

## Common libraries
- [common-go/health](https://github.com/common-go/health): include HealthHandler, HealthChecker, SqlHealthChecker
- [common-go/config](https://github.com/common-go/config): to load the config file, and merge with other environments (SIT, UAT, ENV)
- [common-go/log](https://github.com/common-go/log)
- [common-go/middleware](https://github.com/common-go/middleware): to log all http requests, http responses. User can configure not to log the health check.

### common-go/health
To check if the service is available, refer to [common-go/health](https://github.com/common-go/health)
#### *Request:* GET /health
#### *Response:*
```json
{
    "status": "UP",
    "details": {
        "sql": {
            "status": "UP"
        }
    }
}
```
To create health checker, and health handler
```go
	db, err := sql.Open(conf.Driver, conf.DataSourceName)
	if err != nil {
		return nil, err
	}

	sqlChecker := health.NewSqlHealthChecker(db)
	checkers := []health.HealthChecker{sqlChecker}
	healthHandler := health.NewHealthHandler(checkers)
```

To handler routing
```go
	r := mux.NewRouter()
	r.HandleFunc("/health", healthHandler.Check).Methods("GET")
```

### common-go/config
To load the config from "config.yml", in "configs" folder
```go
package main

import (
	"github.com/common-go/config"
)

type Root struct {
	DB DatabaseConfig `mapstructure:"db"`
}

type DatabaseConfig struct {
	Driver         string `mapstructure:"driver"`
	DataSourceName string `mapstructure:"data_source_name"`
}

func main() {
	var conf Root
	err := config.Load(&conf, "configs/config")
	if err != nil {
		panic(err)
	}
}
```

### common-go/log *&* common-go/middleware
```go
import (
	"github.com/common-go/config"
	"github.com/common-go/log"
	m "github.com/common-go/middleware"
	"github.com/gorilla/mux"
)

func main() {
	var conf app.Root
	config.Load(&conf, "configs/config")

	r := mux.NewRouter()

	log.Initialize(conf.Log)
	r.Use(m.BuildContext)
	logger := m.NewStructuredLogger()
	r.Use(m.Logger(conf.MiddleWare, log.InfoFields, logger))
	r.Use(m.Recover(log.ErrorMsg))
}
```
To configure to ignore the health check, use "skips":
```yaml
middleware:
  skips: /health
```
