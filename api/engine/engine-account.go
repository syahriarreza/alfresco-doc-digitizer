package engine

import (
	"fmt"

	"github.com/eaciit/toolkit"
	"github.com/gin-gonic/gin"
)

// AccountEngine AccountEngine
type AccountEngine struct {
}

// Register Register
func (eg *AccountEngine) Register(c *gin.Context) {
	SetResultOK(c, nil)
	return
}

// Login Login
func (eg *AccountEngine) Login(c *gin.Context) {
	req := struct {
		Name string `json:"name"`
	}{}
	if e := c.ShouldBindJSON(&req); e != nil {
		SetResultError(c, e)
		return
	}

	nameGet := c.Query("name")
	fmt.Println("### GET:", nameGet)

	fmt.Println("### POST:", req.Name)

	SetResultOK(c, toolkit.M{"fieldA": nameGet, "fieldB": 349})
	return
}

// Extract Extract
func (eg *AccountEngine) Extract(c *gin.Context) {
	res, e := ExtractPDF("test pdf.pdf")
	if e != nil {
		SetResultError(c, e)
		return
	}
	fmt.Println(res)

	SetResultOK(c, nil)
	return
}
