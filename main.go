package main

import (
	"github.com/Rozelin-dc/SystemDesignTodoList/repository"
	"github.com/Rozelin-dc/SystemDesignTodoList/router"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func main() {
	db := sqlx.MustConnect("mysql", "root:password@tcp(mysql:3306)/todolist?parseTime=true")

	repo := repository.NewRepository(db)

	echo := router.NewRouter(repo)
	if err := echo.Start(":80"); err != nil {
		panic(err)
	}
}
