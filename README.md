# Task management system

## Requirements
1. Project must contain .env file in root
    ```
        PORT=9090 <or any port to start app on>
        DB_HOST=
        DB_PORT=
        DB_USER=
        DB_PASSWORD=
        DB_NAME=
        HASH_SALT= <smiple salt to hash password>
        TOKEN_SIGNIN_KEY= <simple key to hash JWT token>
    ```
2. To start the app simply run from root dir `go run ./cmd/main.go`

```
go version  1.20
```