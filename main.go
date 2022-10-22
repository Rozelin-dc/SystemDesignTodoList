package main

import (
	"github.com/Rozelin-dc/SystemDesignTodoList/handler"
	"github.com/gorilla/sessions"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	db := sqlx.MustConnect("mysql", "root:password@tcp(mysql:3306)/todolist?parseTime=true")

	h := handler.NewHandler(db)

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(session.Middleware(sessions.NewCookieStore([]byte("secret"))))

	e.POST("/api/login", h.PostLogin)
	e.POST("/api/logout", h.NotImpl)

	api := e.Group("/api", h.EnsureAuthorized())
	{
		apiUser := api.Group("/user")
		{
			apiUser.POST("", h.NotImpl)
			apiUser.DELETE("", h.NotImpl)
			apiUser.PATCH("/:uid", h.NotImpl)
			apiUser.GET("/me", h.NotImpl)
		}

		apiTask := api.Group("/task")
		{
			apiTask.POST("", h.NotImpl)
			apiTask.PATCH("/:tid", h.NotImpl)
			apiTask.DELETE("/:tid", h.NotImpl)
		}
	}
	if err := e.Start(":80"); err != nil {
		panic(err)
	}
}
