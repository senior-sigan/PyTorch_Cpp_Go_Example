package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetFormatter(&logrus.TextFormatter{
		ForceColors:      true,
		DisableTimestamp: true,
	})
	e := echo.New()
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.POST("/deepdream", func(c echo.Context) error {
		file, err := c.FormFile("file")
		if err != nil {
			logrus.Error(err)
			return err
		}
		src, err := file.Open()
		if err != nil {
			logrus.Error(err)
			return err
		}
		defer src.Close()

		name := fmt.Sprintf("%d.jpg", time.Now().Unix())
		fmt.Println(name)
		os.Mkdir("uploads", os.ModePerm)
		dst, err := os.Create(path.Join("uploads", name))
		if err != nil {
			logrus.Error(err)
			return err
		}
		defer dst.Close()

		if _, err = io.Copy(dst, src); err != nil {
			logrus.Error(err)
			return err
		}

		return c.String(http.StatusOK, "File saved")
	})
	e.Logger.Fatal(e.Start(":1323"))
}
