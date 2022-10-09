# Clean Architecture in Go

This project is a starting point for a Flutter application.
- Using `Golang` language and `MongoDB` database.
- Implement CLEAN architecture.
- Simple authentication and jwt validation.
- Localization with Header.Accept-Language. (default = en)

## Build

  - Run API: `go run -tags dev ./main.go`
  - Build API: `go build -tags dev -o ./bin/api ./main.go`

NOTE: Replace `dev` to `staging` or `production` if you want to build/run other enviroment.

## API requests 

### Register
```
curl --location --request POST 'http://localhost:8080/api/auth/register' \
--header 'Accept-Language: vi' \
--header 'Content-Type: application/json' \
--data-raw '{
    "dial_code" : "+84",
    "phone" : "383703713"
}'
```

### Login
```
curl --location --request POST 'http://localhost:8080/api/auth/login' \
--header 'Accept-Language: vi' \
--header 'Content-Type: application/json' \
--data-raw '{ 
    "dial_code" : "+84",
    "phone" : "3837037313"
}'
```

### Code Verification
```
curl --location --request POST 'http://localhost:8080/api/auth/codeVerification' \
--header 'Accept-Language: vi' \
--header 'Content-Type: application/json' \
--data-raw '{
    "verification_id" : "eaa47776-ba1b-4b05-8be9-dfa1673f35ca",
    "code" : "123456"
}'
```

### Get Profile

```
curl --location --request GET 'http://localhost:8080/api/user/profile' \
--header 'Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiZDc2MGYwMzUtYjA4Ny00NDBlLWJiYzAtMTVjYTEyMWUzNzA5IiwiZXhwIjoxNjY1MDY1MDQ4LCJpYXQiOjE2NjQ5Nzg2NDgsImlzcyI6InR1YW5od2luZyJ9.sa-eaMqHNcwO-P_3zoF_NZ78eylZ9OAzxsInTqUlVn8' \
--header 'Accept-Language: vi'
```