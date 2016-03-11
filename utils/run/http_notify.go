package main

import (
	"fmt"
	"github.com/labstack/echo"
	mw "github.com/labstack/echo/middleware"
	"io/ioutil"
	"net/http"
)

func main() {
	m := echo.New()

	// Middleware
	m.Use(mw.Recover())

	m.Any("/*", h)

	// Start server
	m.Run(":1323")
}

func h(c *echo.Context) error {
	request := c.Request()
	fmt.Println("\033[32;1mURL:", request.RequestURI)
	fmt.Print("\033[32;1mHeader:")
	fmt.Println("\033[0m", request.Header)
	fmt.Print("\033[32;1mFORM:")
	fmt.Println("\033[0m", request.PostForm)
	defer request.Body.Close()
	body, _ := ioutil.ReadAll(request.Body)
	fmt.Print("\033[32;1mBODY:")
	fmt.Println("\033[0m", string(body))
	fmt.Println()
	return c.JSON(http.StatusOK, map[string]int{
		"result": 0,
	})
}
