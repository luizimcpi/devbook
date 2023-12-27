# devbook - golang application


## MySql Instructions

```
docker pull mysql/mysql-server:latest

docker run -d -p 3306:3306 --name mysql-docker-container -e MYSQL_ROOT_PASSWORD=123456 -e MYSQL_DATABASE=devbook -e MYSQL_USER=golang -e MYSQL_PASSWORD=golang mysql/mysql-server:latest

sudo docker exec -it mysql-docker-container bash

mysql -u root -p
or
mysql -u golang -p
Password: golang

show databases;
use devbook;

CREATE TABLE usuarios(
    id int auto_increment primary key,
    nome varchar(50) not null,
    email varchar(50) not null
) ENGINE=INNODB;

show tables;
```

## Go db connection
```
go mod init banco-de-dados

go get github.com/go-sql-driver/mysql
```

## Run application
```
go run banco-de-dados.go
```