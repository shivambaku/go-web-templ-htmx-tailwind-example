package handler

import (
	"github.com/labstack/echo"
	record "github.com/shivambaku/fuufu-app/views/records"
)

func HandlerRecords(c echo.Context) error {
	return view(c, record.RecordForm())
}
