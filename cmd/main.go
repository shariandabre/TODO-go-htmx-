package main

import (
	"fmt"
	"go/main/view/components"
	"go/main/view/layout"

	"github.com/labstack/echo/v4"
)

func main() {
	fmt.Println("server running..")
	app := echo.New()

	app.GET("/", func(c echo.Context) error {
		return layout.Base().Render(c.Request().Context(), c.Response())
	})

	app.POST("/add-item", func(c echo.Context) error {
		item := c.FormValue("add_input")
		return components.Todo_item(item).Render(c.Request().Context(), c.Response())
	})

	app.DELETE("/del-item", func(c echo.Context) error {
		return nil
	})

	app.PUT("/completed-item", func(c echo.Context) error {
		name := c.QueryParam("name")
		return components.Swap_items(name).Render(c.Request().Context(), c.Response())
	})

	app.Start(":3000")
}
