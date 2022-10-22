package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/sessions"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	db := sqlx.MustConnect("mysql", "root:password@tcp(mysql:3306)/todolist?parseTime=true")

	repo := repository.NewRepository(db)

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(session.Middleware(sessions.NewCookieStore([]byte("secret"))))

	e.POST("/login", repo.NotImpl)

	apiUser := e.Group("/user")
	{
		apiUser.POST("", repo.NotImpl)
		apiUser.DELETE("", repo.NotImpl)
		apiUser.PATCH("/:uid", repo.NotImpl)
		apiUser.GET("/me", repo.NotImpl)
	}

	apiTask := e.Group("/task")
	{
		apiTask.POST("", repo.NotImpl)
		apiTask.PATCH("/:tid", repo.NotImpl)
		apiTask.DELETE("/:tid", repo.NotImpl)
	}

	if err := e.Start(":80"); err != nil {
		panic(err)
	}
}
