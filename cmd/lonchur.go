package main

import (
	"log"
	"net/http"

	"github.com/a-h/templ"
	"github.com/agneum14/lonchur/views/pages"
	"github.com/agneum14/lonchur/yamlconfig"
	"github.com/labstack/echo/v4"
)

func render(c echo.Context, cmp templ.Component) error {
	c.Response().Writer.WriteHeader(http.StatusOK)
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)
	return cmp.Render(c.Request().Context(), c.Response().Writer)
}

func main() {
    config, colors, err := yamlconfig.Parse()
    if err != nil {
        log.Fatal(err)
    }

	e := echo.New()
	e.Static("/static", "assets")
	e.GET("/", func(c echo.Context) error {
		return render(c, pages.Home(config, colors))
	})
	e.Logger.Fatal(e.Start(":3000"))
}
