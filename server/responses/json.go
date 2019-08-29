package responses

import (
	"github.com/labstack/echo"

	"git.jasonraimondi.com/jason/jasontest/app/lib/pagination"
)

type StatusCode int
type Data = map[string]interface{}

func SendAny(c echo.Context, s StatusCode, d interface{}) error {
	return c.JSON(int(s), d)
}

func SendData(c echo.Context, s StatusCode, d Data) error {
	return c.JSON(int(s), d)
}

func SendPaginator(c echo.Context, s StatusCode, p *pagination.Paginator) error {
	return c.JSON(int(s), p)
}

func SendMessage(c echo.Context, s StatusCode, message string) error {
	return c.JSON(int(s), Data{
		"message": message,
	})
}
