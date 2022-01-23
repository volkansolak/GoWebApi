package main

import (
	"fmt"
	config "goWebApi/models"
	"goWebApi/sql"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)


func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		return c.HTML(http.StatusOK, "Hello, World!")
	})

	e.GET("/ping", func(c echo.Context) error {
		return c.JSON(http.StatusOK, struct{ Status string }{Status: "OK"})
	})

	e.GET("/getalluser", func(c echo.Context) (err error){
		var db = sql.OpenConnection()
		defer db.Close()
		allUser,err := sql.GetAllUser(db)
		if err != nil{
			return c.JSON(http.StatusBadRequest, err)
		}
		return c.JSON(http.StatusOK, allUser)
	})

	e.POST("/adduser", func(c echo.Context) (err error) {
		u := new(config.User)
		if err = c.Bind(u); err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}

		var db = sql.OpenConnection()
		defer db.Close()

		err = sql.AddUser(db, *u)
		if err != nil{
			return c.JSON(http.StatusBadRequest, err)
		}
		return c.JSON(http.StatusOK,"")
	})

	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		httpPort = "8080"
	}

	e.Logger.Fatal(e.Start(":" + httpPort))
}

func Create(c echo.Context) error {
	fmt.Println("Test1")
	fmt.Println(c)
	fmt.Println("Test2")
	return nil
}

type User struct {
	Email string `json:"email" form:"email" query:"email"`
	Password  string `json:"password" form:"password" query:"password"`
}