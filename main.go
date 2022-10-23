package main

import (
	"github.com/Rozelin-dc/SystemDesignTodoList/handler"
	mid "github.com/Rozelin-dc/SystemDesignTodoList/middleware"
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
	e.POST("/api/user", h.PostUser)

	api := e.Group("/api", mid.EnsureAuthorized(h))
	{
		api.POST("/logout", h.PostLogout)

		apiUser := api.Group("/user")
		{
			apiUser.GET("/me", h.GetUserMe)

			apiUserId := apiUser.Group("/:uid", mid.EnsureAccessRightToAccount(h))
			{
				apiUserId.PATCH("", h.PatchUser)
				apiUserId.DELETE("", h.DeleteUser)
			}
		}

		apiTask := api.Group("/task")
		{
			apiTask.POST("", h.NotImpl)

			apiTaskId := apiTask.Group("/:tid", mid.EnsureExistTaskAndHaveAccessRight(h))
			{
				apiTaskId.PATCH("/:tid", h.NotImpl)
				apiTaskId.DELETE("/:tid", h.NotImpl)
			}
		}
	}
	if err := e.Start(":80"); err != nil {
		panic(err)
	}
}
