package helpers

import "github.com/labstack/echo/v4"

func SuccessResponse(c echo.Context, status int, message string, data interface{}) error {
	return c.JSON(status, echo.Map{
		"success": true,
		"messages": message,
		"data": data,
	})
}

func FailedResponse(c echo.Context, status int, message string, err error) error {
	return c.JSON(status, echo.Map{
		"success": false,
		"messages": message,
		"data": err.Error(),
	})
}