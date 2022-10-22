package router

import (
	"github.com/Rozelin-dc/SystemDesignTodoList/repository"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewRouter(repo *repository.Repository) *echo.Echo {
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

	return e
}
