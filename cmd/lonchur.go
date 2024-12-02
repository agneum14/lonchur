package main

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/agneum14/lonchur/views/components"
	"github.com/agneum14/lonchur/views/pages"
	"github.com/labstack/echo/v4"
)

func render(c echo.Context, cmp templ.Component) error {
	c.Response().Writer.WriteHeader(http.StatusOK)
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)
	return cmp.Render(c.Request().Context(), c.Response().Writer)
}

func main() {
	colors := components.Colors{
		Title:           "#BD93F9",
		Background:      "#282A36",
		Text:            "#F8F8F2",
		BackgroundLight: "#44475A",
	}

	sectionData := []components.SectionData{
		{Icon: "fa-solid fa-gear", Title: "Administration"},
		{Icon: "fa-solid fa-cloud", Title: "Media"},
		{Icon: "fa-solid fa-user", Title: "Public"},
		{Icon: "fa-solid fa-paperclip", Title: "Miscellaneous"},
	}

	e := echo.New()
	e.Static("/static", "assets")
	e.GET("/", func(c echo.Context) error {
		return render(c, pages.Home(colors, sectionData))
	})
	e.Logger.Fatal(e.Start(":3000"))
}
