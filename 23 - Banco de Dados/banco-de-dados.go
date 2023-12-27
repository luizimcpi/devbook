package main

import (
	"database/sql"
	"fmt"
	"log"
	
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	stringConexao := "golang:golang@/devbook?charset=utf8&parseTime=True&loc=Local"
	db, erro := sql.Open("mysql", stringConexao)
	if erro != nil {
		log.Fatal(erro)
	}
	defer db.Close() //antes da função main terminar a conexão é fechada

	if erro = db.Ping(); erro != nil {
		log.Fatal(erro)
	}

	fmt.Println("Conexão está aberta!")

	linhas, erro := db.Query("select * from usuarios")
	if erro != nil {
		log.Fatal(erro)
	}
	defer linhas.Close()

	fmt.Println(linhas)

}