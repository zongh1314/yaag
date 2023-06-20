package middleware

import (
	"github.com/labstack/echo"
	"github.com/zongh1314/yaag/middleware"
	"github.com/zongh1314/yaag/yaag"
	"github.com/zongh1314/yaag/yaag/models"
)

func Yaag() echo.MiddlewareFunc {
	return echo.MiddlewareFunc(func(next echo.HandlerFunc) echo.HandlerFunc {
		return echo.HandlerFunc(func(c echo.Context) error {
			if !yaag.IsOn() {
				return next(c)
			}

			apiCall := models.ApiCall{}
			writer := middleware.NewResponseRecorder(c.Response().Writer)
			c.Response().Writer = writer
			middleware.Before(&apiCall, c.Request())
			err := next(c)
			middleware.After(&apiCall, writer, c.Request())
			return err
		})
	})
}
