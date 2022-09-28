# Clean Architecture in Go

This project is a starting point for a Flutter application.
- Using `Golang` language and `MongoDB` database.
- Implement CLEAN architecture.
- Simple authentication and jwt validation.

## Build

  - Run API: `go run -tags dev ./api/main.go`
  - Build API: `go build -tags dev -o ./bin/api ./api/main.go`
  - Build SEARCH: `go build -tags dev -o ./bin/search ./cmd/main.go`

NOTE: Replace `dev` to `staging` or `production` if you want to build/run other enviroment.

## API requests 

### Register
```
curl --location --request POST 'http://localhost:8080/api/auth/register' \
--header 'Content-Type: application/json' \
--data-raw '{
    "name" : "name",
    "email" : "email@gmail.com",
    "password" : "123456"
}'
```

### Login
```
curl --location --request POST 'http://localhost:8080/api/auth/login' \
--header 'Content-Type: application/json' \
--data-raw '{ 
    "email" : "email123@gmail.com",
    "password" : "123456"
}'
```

### Get Profile

```
curl --location --request GET 'http://localhost:8080/api/user/profile' \
--header 'Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiZmYwNGViMTYtMjE3OC00MDU5LWFjZWYtNDA4Mjc4ZWIyOWUxIiwiZXhwIjoxNjY0NDYyMDU0LCJpYXQiOjE2NjQzNzU2NTQsImlzcyI6InR1YW5od2luZyJ9.P8hoRDx-I2xK86Lpa7b8ud_-qQ1a_58Ne6pHTqHGq9E'
```

### Insert Book
```
curl --location --request POST 'http://localhost:8080/api/book/' \
--header 'Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiZmYwNGViMTYtMjE3OC00MDU5LWFjZWYtNDA4Mjc4ZWIyOWUxIiwiZXhwIjoxNjY0NDYyMDU0LCJpYXQiOjE2NjQzNzU2NTQsImlzcyI6InR1YW5od2luZyJ9.P8hoRDx-I2xK86Lpa7b8ud_-qQ1a_58Ne6pHTqHGq9E' \
--header 'Content-Type: application/json' \
--data-raw '{
    "title" : "title temporary",
    "description" : "lollsdasd"
}'
```

### Get Book
```
curl --location --request POST 'http://localhost:8080/api/book/' \
--header 'Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiZmYwNGViMTYtMjE3OC00MDU5LWFjZWYtNDA4Mjc4ZWIyOWUxIiwiZXhwIjoxNjY0NDYyMDU0LCJpYXQiOjE2NjQzNzU2NTQsImlzcyI6InR1YW5od2luZyJ9.P8hoRDx-I2xK86Lpa7b8ud_-qQ1a_58Ne6pHTqHGq9E' \
--header 'Content-Type: application/json' \
--data-raw '{
    "title" : "title temporary",
    "description" : "description"
}'