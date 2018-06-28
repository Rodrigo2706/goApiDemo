# GoApiDemo

A Golang API using mux(https://github.com/gorilla/mux),
HAL/JSON (https://github.com/nvellon/hal),
ozzo-validation (https://github.com/go-ozzo/ozzo-validation) and
MySQL (https://github.com/go-sql-driver/mysql)

## Installation
Steps to follow to use this API

```
Create a Trainings folder on d:
cd \Trainings
git clone https://github.com/Rodrigo2706/goApiDemo.git
cd GoApiDemo

set GOPATH=d:\Trainings
set GOBIN=d:\Trainings\bin

go get github.com/go-sql-driver/mysql
go get github.com/go-ozzo/ozzo-validation
go get github.com/go-ozzo/ozzo-validation/is
go get github.com/nvellon/hal
go get -u github.com/gorilla/mux

```

Under `src->utils->database->database.go` you can find the DB configuration:
https://www.screencast.com/t/n9SU8NgAO
Change accordingly and in any MySQL manager run the database_schema.sql file.

## Running the project
```
cd src/main
go run main.go
```

## Usage
From any Rest Client you can call
```
GET /users                      // To get a list of all users
GET /users?offset=#1&limit=#2   // To get a batch of user records
POST /users                     // To create a user in the database
{
  "Email": "aValid@ema.il",
  "Name": "aName",
  "Lastname": "aLastName"
}
PATCH /users/{email}            // To update a users info
{
  "Name": "aName",
  "Lastname": "aLastName"
}
GET /users/{email}              // To get a users info by email
```

