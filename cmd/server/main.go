package main

import (
	"net/http"

	"Musica-Backend/di"

	"Musica-Backend/internal/infrastructure/postgre"

	redis "Musica-Backend/internal/infrastructure/redis"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	postgreDb := postgre.NewGormPostgres()
	defer func() {
		d, _ := postgreDb.DB()
		d.Close()
	}()
	//eed.InsertInitialData(postgreDb)
	engine := echo.New()
	engine.Use(middleware.Logger())
	engine.Use(middleware.Recover())
	// CORS設定
	engine.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:5173"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))

	PostHander := di.InitializePostHandler()
	findAll := func(c echo.Context) error {
		posts, err := PostHander.FindAll(c)
		if err != nil {
			return c.JSON(500, err)
		}
		return c.JSON(200, posts)
	}
	UserHandler := di.InitializeUserHandler()
	findUserById := func(c echo.Context) error {

		user, err := UserHandler.FindUserById(c)
		if err != nil {
			return c.JSON(500, err)
		}
		return c.JSON(200, user)
	}
	redis.InitRedis()
	engine.GET("/users/:id", findUserById)
	engine.GET("/posts", findAll)
	engine.Start(":8080")

}
