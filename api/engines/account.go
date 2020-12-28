package engine

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type AccountEngine struct {
}

func (ae *AccountEngine) Register(c *gin.Context) {
	name := c.Param("name")
	c.String(http.StatusOK, "Hello %s", name)
}
