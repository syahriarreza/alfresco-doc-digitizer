package engine

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// SetResultOK SetResultOK
func SetResultOK(c *gin.Context, obj ...interface{}) {
	if len(obj) > 0 {
		c.JSON(http.StatusOK, obj[0])
	} else {
		c.JSON(http.StatusOK, gin.H{"status": "OK"})
	}
}

// SetResultError SetResultError
func SetResultError(c *gin.Context, err error) {
	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
}
