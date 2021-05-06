package controller

import (
	"crypto/rand"
	"fmt"
	"github.com/labstack/echo/v4"
	"io"
	"net/http"
)

type Monitoring struct {
	transID string
}

func (this *Monitoring) Monitoring(c echo.Context) error {

	uuid := make([]byte, 16)
	n, err := io.ReadFull(rand.Reader, uuid)
	if n != len(uuid) || err != nil {
		return c.String(http.StatusInternalServerError, "InternalServerError")
	}

	uuid[8] = uuid[8]&^0xc0 | 0x80
	uuid[6] = uuid[6]&^0xf0 | 0x40

	this.transID = fmt.Sprintf("%x-%x-%x-%x-%x", uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:])
	return c.String(http.StatusOK, "SUCCESS")
}
