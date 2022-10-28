package main

import (
	"strings"

	"github.com/Rozelin-dc/SystemDesignTodoList/handler"
	mid "github.com/Rozelin-dc/SystemDesignTodoList/middleware"
	"github.com/Rozelin-dc/SystemDesignTodoList/util"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	db := sqlx.MustConnect("mysql", "root:password@tcp(mysql:3306)/todolist?parseTime=true")

	h := handler.NewHandler(db)

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Validator = util.GetValidator()

	e.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Skipper: func(c echo.Context) bool {
			return strings.HasPrefix(c.Request().URL.Path, "/api")
		},
		Root:  "web/dist",
		HTML5: true,
	}))

	e.POST("/api/login", h.PostLogin)

	api := e.Group("/api", mid.EnsureAuthorized(h))
	{
		api.POST("/logout", h.PostLogout)

		apiUser := api.Group("/user")
		{
			apiUser.POST("", h.PostUser)

			apiUser.GET("/me", h.GetUserMe)

			apiUserId := apiUser.Group("/:uid", mid.EnsureAccessRightToAccount(h))
			{
				apiUserId.PATCH("", h.PatchUser)
				apiUserId.DELETE("", h.DeleteUser)
			}
		}

		apiTask := api.Group("/task")
		{
			apiTask.POST("", h.PostTask)
			apiTask.GET("", h.GetTasks)

			apiTaskId := apiTask.Group("/:tid", mid.EnsureExistTaskAndHaveAccessRight(h))
			{
				apiTaskId.PATCH("", h.PatchTask)
				apiTaskId.DELETE("", h.DeleteTask)
			}
		}
	}

	e.Logger.Fatal(e.Start(":80"))
}
