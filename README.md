# devbook - golang application crud

## Init app module
```
go mod init crud
```

## Go mux
```
go get github.com/gorilla/mux
```

## Go db connection
```
go get github.com/go-sql-driver/mysql
```

## Run application
```
go run main.go
```

## Curl 

### Create User 
```
curl --location --request POST 'localhost:5000/usuarios' \
--header 'Content-Type: application/json' \
--data-raw '{
    "nome": "Luiz",
    "email": "luiz.henrique@gmail.com"
}'
```

### Get users
```
curl --location --request GET 'localhost:5000/usuarios'
```