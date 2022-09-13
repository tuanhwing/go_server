# Clean Architecture in Go

[![Build Status](https://travis-ci.org/eminetto/clean-architecture-go-v2.svg?branch=master)](https://travis-ci.org/eminetto/clean-architecture-go-v2)

[Post: Clean Architecture, 2 years later](https://eltonminetto.dev/en/post/2020-07-06-clean-architecture-2years-later/)

## Build

  - Run API: `go run -tags dev ./api/main.go`
  - Build API: `go build -tags dev -o ./bin/api ./api/main.go`
  - Build SEARCH: `go build -tags dev -o ./bin/search ./cmd/main.go`

NOTE: Replace `dev` to `staging` or `production` if you want to build/run other enviroment.

## API requests 

### Add user

```
curl --location --request POST 'http://localhost:8080/v1/user' \
--header 'Content-Type: application/json' \
--data-raw '{
    "email" : "email@gmail.com",
    "password" : "123456",
    "first_name" : "name1",
    "last_name" : "name2"
}'
```
### Return a user

```
curl --location --request GET 'http://localhost:8080/v1/user/c83277c5-2568-497c-9757-37df102bc29c'
```

### Show user with name

```
curl --location --request GET 'http://localhost:8080/v1/user?name=name1'
```


## CMD 

### Search for users

```
./bin/search name1
```
