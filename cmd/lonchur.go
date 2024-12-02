package main

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/agneum14/lonchur/views/pages"
	"github.com/labstack/echo/v4"
)

func render(c echo.Context, cmp templ.Component) error {
	c.Response().Writer.WriteHeader(http.StatusOK)
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)
	return cmp.Render(c.Request().Context(), c.Response().Writer)
}

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return render(c, pages.Home())
	})
	e.Logger.Fatal(e.Start(":3000"))
}
